<template>
  <el-container>
    <el-aside width="auto">
      <common-aside></common-aside>
    </el-aside>
    <el-container>
      <el-header></el-header>
      <el-main></el-main>
    </el-container>
  </el-container>
</template>

<script>
import CommonAside from '@/components/CommonAside'
export default {
  name: 'HomeView',
  components: {
    CommonAside
  },
  data () {
    return {}
  }
}
</script>

<style>
.el-header {
  color: #333;
  background-color: #B3C0D1;
  text-align: center;
  line-height: 100%;
  max-height: 20%;
}

.el-aside {
  color: #333;
  text-align: center;
  line-height: 100%;
  max-height: 100%;
}

.el-main {
  color: #333;
  text-align: center;
  line-height: 80%;
}
</style>
