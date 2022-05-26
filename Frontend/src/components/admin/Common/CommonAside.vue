<template>
  <el-aside width="100%">
    <div class="toggle-button" @click=toggleCollapse>|||</div>
    <!-- 侧边栏菜单 -->
    <el-menu text-color="#FFF" background-color="#484848" active-text-color="#ffd04b" :unique-opened="true" :collapse="isCollapse" :collapse-transition="false">
      <!-- 一级菜单 -->
      <el-submenu v-for="item in MenuList" :index="item.MainMenuID + ''" :key="item.MainMenuID">
        <!-- 一级菜单模版区 -->
        <template slot="title">
          <!-- 图标 -->
          <i class="el-icon-menu"></i>
          <!-- 文本 -->
          <span>{{item.MainMenuName}}</span>
        </template>
        <!-- 二级菜单 -->
        <el-menu-item v-for="subitem in item.SubMenu" :index="subitem.SubMenuUID" :key="subitem.SubMenuUID">
          <template slot="title">
            <i class="el-icon-menu"></i>
            <span>{{subitem.SubMenuName}}</span>
          </template>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
import JwtDecode from 'jwt-decode'
export default {
  created () {
    this.getMenuList()
  },
  data () {
    return {
      // 左侧菜单栏
      MenuList: [],
      // 是否折叠
      isCollapse: false
    }
  },
  methods: {
    // 获取该用户授权菜单
    async getMenuList () {
      const RoleUID = JwtDecode(window.sessionStorage.getItem('token')).RoleUID
      const { data: res } = await this.$http.post('menu', { RoleUID: RoleUID }, { headers: { 'Content-Type': 'application/json' } })
      if (res.data.length === 0) return this.$message.error('获取菜单栏失败')
      this.MenuList = res.data
    },
    toggleCollapse () {
      this.isCollapse = !this.isCollapse
    }
  }
}
</script>

<style scoped>
.toggle-button {
  background-color: #454545;
  font-size: 10px;
  color: #FFF;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
