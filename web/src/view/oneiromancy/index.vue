<template>
  <div class="chat-container">
    <el-container class="chat-wrapper">
      <el-main>
        <el-scrollbar ref="scrollbar" class="chat-scrollbar">
          <div v-for="(message, index) in messages" :key="index" class="message">
            <div 
              :class="['message-wrapper', message.sender === 'user' ? 'user-message' : 'ai-message']">
              <el-card 
                class="message-card" 
                shadow="always" 
                v-html="message.content">
              </el-card>
            </div>
          </div>
        </el-scrollbar>
      </el-main>
      <el-footer class="footer">
        <el-input
          type="textarea"
          placeholder="请输入梦境内容"
          @keyup.enter="handleKeyup"
          v-model="inputMessage"
          rows="3"
          clearable>
        </el-input>
        <div class="footer-button">
          <el-button 
            type="primary" 
            icon="Edit" 
            round 
            @click="handleSend" 
            :disabled="isSending">发送</el-button>
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
    const query = new URLSearchParams("query").toString();
    const eventSource = new EventSource(`/api/oneiromancy/query_stream?${query}${inputValue.trim()}`);
    // 监听消息事件
    // 创建一个AI消息的条目，如果列表中还没有的话
    let aiMessageEntry = '';
    if (!aiMessageEntry) {
      aiMessageEntry = { sender: 'AI', content: '' };
      messages.value.push(aiMessageEntry);
    }
    eventSource.onmessage = function(event) {
      const dataWithBreaks = event.data.replace(/\n\n/g, '<br>');
      aiMessageEntry.content +=dataWithBreaks; // 添加换行符以分隔消息
      messages.value.pop();
      messages.value.push({ sender: 'AI', content: aiMessageEntry.content });
    };
    // 监听错误事件
    eventSource.onerror = function(error) {
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
  height: 75vh;
  background-color: #f4f4f5;
}

.chat-wrapper {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-scrollbar {
  flex: 1;
  padding: 15px;
  background-color: #fff;
  border-radius: 8px;
  margin-bottom: 10px;
  overflow-y: auto;
  max-height: calc(100vh - 250px); /* 限制高度 */
}

.message {
  display: flex;
  margin-bottom: 15px;
}

.message-wrapper {
  max-width: 60%;
}

.user-message {
  justify-content: flex-end;
  align-items: flex-end;
}

.ai-message {
  justify-content: flex-start;
  align-items: flex-start;
}

.message-card {
  padding: 10px;
  border-radius: 10px;
  background-color: #e0f7fa;
  color: #00796b;
}

.user-message .message-card {
  background-color: #e3f2fd;
  color: #1565c0;
}

.ai-message .message-card {
  background-color: #f1f8e9;
  color: #33691e;
}

.footer-button {
  text-align: center;
  margin-top: 15px;
}

.el-input {
  width: 100%;
  border-radius: 6px;
}
</style>