<template>
    <div class="chat-container">
      <el-container>
        <!-- <el-header>
          <h1>请输入梦境</h1>
        </el-header> -->
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
            v-model="inputMessage">
          </el-input>
          <div style="text-align: center; margin-top: 20px;"> <!-- 添加一个div用于居中按钮 -->
            <el-button type="primary" :icon="Edit" round @click="sendMessage">发送</el-button>
          </div>
        </el-footer>
      </el-container>
    </div>
  </template>

<script setup>
import { ref, nextTick } from 'vue';

const inputMessage = ref('');
const messages = ref([]);

const sendMessage = async () => {
  if (inputMessage.value.trim() !== '') {
    messages.value.push({ sender: 'user', content: inputMessage.value });
    try {
      // 发送 POST 请求到后端 /api/oneiromancy/query 路由
      const response = await fetch('/api/oneiromancy/query', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ query: inputMessage.value })
      });

      // 解析 JSON 响应
      const data = await response.json();

      // 假设后端返回的数据结构是 { result: "回复内容" }
      if (response.ok) {
        // 将后端返回的回复添加到消息列表中
        messages.value.push({ sender: 'AI', content: data.data });
        scrollToBottom();
      } else {
        // 如果响应不是 200 OK，处理错误情况
        console.error('Error:', data.error);
        // 可以选择将错误信息显示给用户
        messages.value.push({ sender: 'AI', content: '发生错误，请重试。'});
        scrollToBottom();
      }
    } catch (error) {
      // 处理网络错误或其他异常
      console.error('Fetch error:', error);
      messages.value.push({ sender: 'AI', content: '网络错误，请检查您的连接。'});
      scrollToBottom();
    }
    inputMessage.value = '';
    scrollToBottom();
  }
};

const scrollToBottom = () => {
  nextTick(() => {
    const container = document.querySelector('.scrollbar'); // 假设你的滚动容器有一个类名 'scrollbar'
    if (container) {
      container.scrollTop = container.scrollHeight;
    }
  });
};
</script>

<style scoped>
.chat-container {
  height: 100vh;
}
.chat-scrollbar {
  height: calc(100vh - 120px);
  overflow-y: auto;
}
.message {
  margin: 10px;
}
.item {
  text-align: left;
}
.el-main {
  height: calc(100vh - 300px); /* 视口高度减去200px */
}
</style>
        