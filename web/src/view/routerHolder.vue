<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive :include="routerStore.keepAliveRouters">
          <div>
            <component :is="Component" />
          </div>
        </keep-alive>
      </transition>
    </router-view>

    <!-- 问答机器人部分 -->
    <div class="qa-robot">
      <h2>问答机器人</h2>
      <input
        type="text"
        v-model="question"
        placeholder="请输入你的问题"
        @keyup.enter="askQuestion"
      />
      <button @click="askQuestion">提问</button>
      <div v-if="loading">正在获取答案...</div>
      <div v-else-if="answer">{{ answer }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouterStore } from '@/pinia/modules/router'
import axios from 'axios' // 确保axios已经安装

defineOptions({
  name: 'RouterHolder'
})

const routerStore = useRouterStore()

// 问答机器人数据
const question = ref('')
const answer = ref('')
const loading = ref(false)

// 问答机器人方法
const askQuestion = async () => {
  if (!question.value) return

  loading.value = true
  answer.value = '' // 清空上一次的答案

  try {
    const response = await axios.post('https://your-api-endpoint.com/answer', { question: question.value })
    answer.value = response.data.answer // 假设API返回的数据结构中有answer字段
  } catch (error) {
    answer.value = '发生错误，请稍后再试。'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.qa-robot {
  /* 样式根据需求自定义 */
  margin-top: 20px;
  padding: 10px;
  border: 1px solid #ccc;
}
</style>
