<template>
  <div id="userLayout" class="w-full h-full relative">
    <div
      class="rounded-lg flex items-center justify-evenly w-full h-full md:w-screen md:h-screen md:bg-[#194bfb] bg-white"
    >
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div
          class="oblique h-[130%] w-3/5 bg-white dark:bg-slate-900 transform -rotate-12 absolute -ml-52"
        />
        <!-- 分割斜块 -->
        <div
          class="z-[999] pt-12 pb-10 md:w-96 w-full rounded-lg flex flex-col justify-between box-border"
        >
          <div>
            <div class="flex items-center justify-center">
              <img class="w-24" :src="$GIN_VUE_ADMIN.appLogo" alt />
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">
                {{ $GIN_VUE_ADMIN.appName }}
              </p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">
                A management platform using Golang and Vue
              </p>
            </div>
            <el-form
              ref="registerForm"
              :model="registerFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >
              <el-form-item prop="userName" class="mb-6">
                <el-input
                  v-model="registerFormData.userName"
                  size="large"
                  placeholder="请输入用户名"
                  suffix-icon="userName"
                />
              </el-form-item>
              <el-form-item prop="email" class="mb-6">
                <el-input
                  v-model="registerFormData.email"
                  size="large"
                  placeholder="请输入邮箱"
                  suffix-icon="email"
                />
              </el-form-item>
              <el-form-item prop="passWord" class="mb-6">
                <el-input
                  v-model="registerFormData.passWord"
                  show-password
                  size="large"
                  type="passWord"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item prop="confirmPassword" class="mb-6">
                <el-input
                  v-model="registerFormData.confirmPassword"
                  show-password
                  size="large"
                  type="confirmPassword"
                  placeholder="请确认密码"
                />
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-active h-11 w-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                  >注 册</el-button
                >
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]">
        <img
          class="h-full"
          src="@/assets/login_right_banner.jpg"
          alt="banner"
        />
      </div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto w-full z-20">
      <div class="links items-center justify-center gap-2 hidden md:flex">
        <a href="https://www.gin-vue-admin.com/" target="_blank">
          <img src="@/assets/docs.png" class="w-8 h-8" alt="文档" />
        </a>
        <a href="https://support.qq.com/product/371961" target="_blank">
          <img src="@/assets/kefu.png" class="w-8 h-8" alt="客服" />
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img src="@/assets/github.png" class="w-8 h-8" alt="github" />
        </a>
        <a href="https://space.bilibili.com/322210472" target="_blank">
          <img src="@/assets/video.png" class="w-8 h-8" alt="视频站" />
        </a>
      </div>
    </BottomInfo>
  </div>
</template>

<script setup>
  import { checkDB } from '@/api/initdb'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { reactive, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'

  // 注册功能
  import {
    user_register
  } from '@/api/user'

  defineOptions({
    name: 'Register'
  })

  const router = useRouter()

  const registerForm = ref(null)
  const registerFormData = reactive({
    userName: '',
    email: '',
    passWord: '',
    confirmPassword: '',
  })

  const checkConfirmPassword = (rule, value, callback) => {
    if (value !== registerFormData.passWord) {
      callback(new Error('确认密码与密码不一致'))
    } else {
      callback()
    }
  }

  const rules = reactive({
    userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
    ],
    email: [
      { required: true, message: '请输入邮箱', trigger: 'blur' },
      {
        pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,
        message: '请输入正确的邮箱',
        trigger: 'blur'
      }
    ],
    passWord: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度应为 6 到 20 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请输入确认密码', trigger: 'blur' },
    { validator: checkConfirmPassword, trigger: 'blur' }
  ]
  })

  const submitForm = () => {
    // 注册逻辑
    registerForm.value.validate(async (v) => {
      if (!v) {
        ElMessage({
          type: 'error',
          message: '请正确填写注册信息',
          showClose: true
        })
        return false
      }

      const req = {
        userName: registerFormData.userName,
        email: registerFormData.email,
        passWord: registerFormData.passWord,
      }
      const res = await user_register(req)

      if (res.code === 0) {
        ElMessage({ type: 'success', message: '创建成功' })

      }
    })


  }

</script>
