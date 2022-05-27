<template>
  <div class="header">
    <div class="header-left">
      <span>Sec Chaos</span>
    </div>
    <div class="header-right">
      <el-button type="danger" @click='LogOut' class="button">退出</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CommonHeader',
  methods: {
    LogOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
.header {
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.header-left {
  width: 15%;
  height: 100%;
  display: flex;
  align-items: center;
  font-size: 20px;
}
span {
  margin-left: 15%;
}
.header-right{
  height: 100%;
  width: 75%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.button {
  margin-left: 15px;
  margin-right: 3%;
}
</style>
