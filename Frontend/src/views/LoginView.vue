<template>
  <div class="login_container">
    <div class="login_box">
      <div class="avatar_box">
        <img src="../assets/avatar.png" alt="">
      </div>
      <!-- 用户登录框 -->
      <el-form :model="LoginForm" :rules="LoginRules" ref="LoginForm" label-width="80px" class="login_form">
        <el-form-item label="用户名" prop="Username">
          <el-input v-model="LoginForm.Username" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="Password">
          <el-input v-model="LoginForm.Password" prefix-icon="el-icon-lock" type="password"></el-input>
        </el-form-item>
        <!-- 按钮区 -->
        <el-form-item class="login_button">
          <el-button type="primary" @click="submitForm('LoginForm')">登录</el-button>
          <el-button type="info" @click="resetForm('LoginForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginView',
  data () {
    return {
      // 登录数据保存对象
      LoginForm: {
        Username: '',
        Password: ''
      },
      // 表单验证规则
      LoginRules: {
        Username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        Password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  background-color: #2b4b6b;
  height: 100%;
}
.login_box {
  width: 450px;
  height: 350px;
  background-color: floralwhite;
  border-radius: 5%;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  .avatar_box {
    height: 130px;
    width: 130px;
    border: 1px;
    border-radius: 50%;
    padding: 10px;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);
    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eeeeee;
    }
  }
}
.login_form {
  position: absolute;
  top: 30%;
  left: 0%;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}
.login_button {
  display: flex;
  justify-content: flex-end;
}
</style>
