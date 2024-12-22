<template>
  <div class="chat-container">
    <el-container>
      <el-main>
        <el-scrollbar ref="scrollbar" class="chat-scrollbar">
          <div v-for="(message, index) in messages" :key="index" class="message">
            <el-badge value="user" class="item" v-if="message.sender === 'user'">
              <el-card shadow="hover">{{ message.content }}</el-card>
            </el-badge>
            <el-badge value="AI" class="item" v-else>
              <el-card shadow="hover">{{ message.content }}</el-card>
            </el-badge>
          </div>
        </el-scrollbar>
      </el-main>
      <el-footer>
        <el-input
          type="textarea"
          placeholder="请输入梦境内容"
          @keyup.enter="handleKeyup"
          v-model="inputMessage">
        </el-input>
        <div style="text-align: center; margin-top: 20px;">
          <el-button type="primary" :icon="Edit" round @click="handleSend" :disabled="isSending">发送</el-button>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup>
import { ref, nextTick, onUnmounted, onMounted } from 'vue';

const isSending = ref(false);
const inputMessage = ref('');
const messages = ref([]);
const eventSource = ref(null);
const scrollbar = ref(null);

const scrollToBottom = () => {
  nextTick(() => {
    const wrap = scrollbar.value?.wrapRef; // 使用 wrapRef 替代 wrap
    if (wrap) {
      wrap.scrollTop = wrap.scrollHeight;
    } else {
      console.error('scrollbar.wrapRef is undefined');
    }
  });
};

const handleSend = async () => {
  isSending.value=true;
  const inputValue = inputMessage.value;
  // 防止发送空消息
  if (inputMessage.value.trim() == '') {
    return;
  }
  inputMessage.value = '';
  
  // 添加用户消息
  messages.value.push({ sender: 'user', content: inputValue.trim() });
  scrollToBottom();

  try {
    const query = new URLSearchParams(inputValue.trim()).toString();
    const eventSource = new EventSource(`/api/oneiromancy/query_stream?query=${query}`);
    // 监听消息事件
    // 创建一个AI消息的条目，如果列表中还没有的话
    let aiMessageEntry = '';
    if (!aiMessageEntry) {
      aiMessageEntry = { sender: 'AI', content: '' };
      messages.value.push(aiMessageEntry);
    }
    eventSource.onmessage = function(event) {
      aiMessageEntry.content += event.data + '\n\n'; // 添加换行符以分隔消息
      console.log(aiMessageEntry.content)
      messages.value.pop();
      messages.value.push({ sender: 'AI', content: aiMessageEntry.content });
    };
    // 监听错误事件
    eventSource.onerror = function(error) {
      console.log(eventSource.CLOSED)
      if (eventSource.CLOSED === EventSource.CLOSED) {
        eventSource.close();
        isSending.value=false;
        return
      } else {
        eventSource.close();
        console.error('An error occurred:', error);
        messages.value.push({ sender: 'AI', content: '发生错误，请重试。' });
        isSending.value=false;
        return
      }
    };} catch (error) {
      // 处理网络错误或其他异常
      eventSource.close();
      isSending.value=false;
      console.error('Fetch error:', error);
      messages.value.push({ sender: 'AI', content: '网络错误，请检查您的连接。' });
      scrollToBottom();
      return
    }
  scrollToBottom();
};

// 处理回车键按下事件
const handleKeyup = (event) => {
  console.log(isSending.value)
  if (event.key === 'Enter' && !isSending.value) {
    console.log('Enter key pressed');
    handleSend();
  }
};

onMounted(() => {
});

onUnmounted(() => {
});
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.chat-scrollbar {
  flex: 1;
  overflow-y: auto;
}

.message {
  margin: 10px;
}

.item {
  text-align: left;
}

.el-main {
  flex: 1;
}

.el-footer {
  padding: 10px;
  height: calc(100vh - 500px);
}
</style>