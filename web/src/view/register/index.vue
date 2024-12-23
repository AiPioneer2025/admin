<template>
  <div class="register-view">
    <el-form :model="registerForm" ref="registerForm" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="registerForm.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="registerForm.password" type="password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="confirmPassword">
        <el-input v-model="registerForm.confirmPassword" type="password" placeholder="请确认密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

export default {
  setup() {
    const router = useRouter();

    const registerForm = ref({
      username: '',
      password: '',
      confirmPassword: '',
    });

    const submitForm = () => {
      // 注册逻辑
      console.log('注册表单提交', registerForm.value);
    };

    return {
      registerForm,
      submitForm,
    };
  },
};
</script>

<style scoped>
.register-view {
  width: 300px;
  margin: 0 auto;
  padding-top: 100px;
}
</style>
