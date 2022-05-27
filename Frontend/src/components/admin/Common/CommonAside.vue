<template>
  <el-aside width="100%">
    <div class="toggle-button" @click=toggleCollapse>|||</div>
    <!-- 侧边栏菜单 -->
    <el-menu text-color="#FFF" background-color="#484848" active-text-color="#ffd04b" :unique-opened="true"
             :collapse="isCollapse" :collapse-transition="false" router :default-active="activePath">
      <!-- 一级菜单 -->
      <el-submenu v-for="item in MenuList" :index="item.MainMenuID"
                  :key="item.MainMenuID">
        <!-- 一级菜单模版区 -->
        <template slot="title">
          <!-- 图标 -->
          <i :class=iconsObj[item.MainMenuID]></i>
          <!-- 文本 -->
          <span>{{item.MainMenuName}}</span>
        </template>
        <!-- 二级菜单 -->
        <el-menu-item v-for="subitem in item.SubMenu"
                      :index="subitem.SubMenuPath" :key="subitem.SubMenuUID"
                      @click="saveNavStat(subitem.SubMenuPath)">
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
    this.activePath = window.sessionStorage.getItem('activePath')
  },
  data () {
    return {
      // 左侧菜单栏
      MenuList: [],
      // 定义一级菜单图标
      iconsObj: {
        // 文章管理
        '988d9bc7-74e8-4bad-945d-8223f4829371': 'el-icon-document',
        // 用户管理
        '906cfa43-bc3a-4c4e-9dd8-2771af5ed031': 'el-icon-user',
        // 权限管理
        '7958a19d-cac0-49a5-96e4-2a0e6ab492f9': 'el-icon-open'
      },
      // 是否折叠
      isCollapse: false,
      // 被激活链接 默认为空
      activePath: ''
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
    // 设置菜单栏折叠与展开
    toggleCollapse (activePath) {
      this.isCollapse = !this.isCollapse
    },
    // 保存链接菜单栏
    saveNavStat (activePath) {
      window.sessionStorage.setItem('activePath', activePath)
      this.activePath = activePath
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
