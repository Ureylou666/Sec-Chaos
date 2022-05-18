<template>
  <div class="login_container">
    <div class="background">
      <img :src="imgSrc" width="100%" height="100%" alt="" />
    </div>
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
          <el-input v-model="LoginForm.Password" prefix-icon="el-icon-lock" type="password" @keyup.enter.native="submitForm('LoginForm')"></el-input>
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
import JwtDecode from 'jwt-decode'

export default {
  name: 'LoginView',
  data () {
    return {
      // 定义背景图片
      imgSrc: require('../assets/Background.jpg'),
      // 登录数据保存对象
      LoginForm: {
        Username: '',
        Password: ''
      },
      // 表单验证规则
      LoginRules: {
        Username: [
          { required: true, min: 4, max: 12, message: '请输入正确的用户名', trigger: 'blur' }
        ],
        Password: [
          { required: true, min: 8, max: 20, message: '请输入正确密码', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 表单登录
    submitForm (formName) {
      this.$refs[formName].validate(async (valid) => {
        // 判断密码是否符合规范
        if (valid) {
          const { data: res } = await this.$http.post('login', this.LoginForm)
          // 判断密码是否正确
          if (res.code !== 200) return this.$message.error(res.message)
          window.sessionStorage.setItem('token', res.token)
          const RoleId = JwtDecode(res.token).role_id
          // 判断用户是否为管理员
          if (RoleId !== 1) {
            return this.$message.error('无权限，装老师傅？')
          } else {
            await this.$router.push('admin')
          }
        } else {
          this.$message.error('装老师傅？')
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
  width: 100%;
  height: 100%;
}
.background {
  width: 100%;
  height: 100%;
  z-index: -1;
  position: absolute;
}
.login_box {
  width: 450px;
  height: 350px;
  background-color: floralwhite;
  border-radius: 5%;
  position: absolute;
  left: 75%;
  top: 50%;
  transform: translate(-50%, -50%);
}
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
