<template>
  <el-aside width="100%">
    <el-radio-group v-model="isCollapse" class="radio-group">
      <el-radio-button :label="false">展开</el-radio-button>
      <el-radio-button :label="true">收起</el-radio-button>
    </el-radio-group>
    <!-- 侧边栏菜单 -->
    <el-menu :default-active="$route.path" router unique-opened class="el-menu-vertical" @open="handleOpen" @close="handleClose" :collapse="isCollapse" text-color="#FFF" background-color="#484848" active-text-color="#409EFF">
      <!-- 模版保留 -->
      <el-submenu index="0">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span slot="title" style="color: #eeeeee">导航一</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="0-1">选项1</el-menu-item>
          <el-menu-item index="0-2">选项2</el-menu-item>
          <el-menu-item index="0-3">选项3</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="0-4">
          <span slot="title">选项4</span>
          <el-menu-item index="0-4-1">选项1</el-menu-item>
        </el-submenu>
      </el-submenu>
      <!-- 漏洞文章管理 -->
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-document"></i>
          <span slot="title" style="color: #eeeeee">内容管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="1-1">分类管理</el-menu-item>
          <el-menu-item index="1-2">文章管理</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <!-- 工具管理 -->
      <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-aim"></i>
          <span slot="title" style="color: #eeeeee">工具管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="2-1">免杀工具</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <!-- 系统管理 -->
      <el-submenu index="3">
        <template slot="title">
          <i class="el-icon-setting"></i>
          <span slot="title" style="color: #eeeeee">系统管理</span>
        </template>
        <el-submenu index="3-1">
          <span slot="title">用户管理</span>
          <el-menu-item :route="{ path: '/admin/user' }" index="3-1-1">用户列表</el-menu-item>
          <el-menu-item index="3-1-2">新增用户</el-menu-item>
        </el-submenu>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  data () {
    return {
      isCollapse: true
    }
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style scoped>
.radio-group {
  display: flex;
  justify-content: center;
  align-items: center;
  padding-bottom: 10%;
}
</style>
