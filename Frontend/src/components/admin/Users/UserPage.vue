<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/welcome' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>用户管理</el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区 -->
    <el-card>
      <!-- 搜索与添加区域 -->
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
        <el-col :span="5">
          <el-button type="primary">添加用户</el-button>
        </el-col>
      </el-row>
      <!-- 用户列表区域 -->
      <el-table :data = "userlist" border stripe>
        <el-table-column type="index" label="#"></el-table-column>
        <el-table-column label="用户名" prop="Username"></el-table-column>
        <el-table-column label="角色" prop="RoleName"></el-table-column>
        <el-table-column label="状态">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.Status" @change="userStateChanged(scope.row.Username)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="300px">
          <template>
            <el-button type="primary" icon="el-icon-edit" size="medium"></el-button>
            <el-button type="danger" icon="el-icon-delete" size="medium"></el-button>
            <el-tooltip class="item" effect="dark" content="角色分配" placement="top" :enterable="false">
              <el-button type="warning" icon="el-icon-setting" size="medium"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[1, 2, 5, 10]"
        :page-size="queryInfo.pagesize"
        layout="total, sizes, prev, pager, next, jumper"
        :total=total>
      </el-pagination>
    </el-card>
  </div>
</template>

<script>
export default {
  data () {
    return {
      // 获取用户列表参数
      queryInfo: {
        pagesize: 5,
        pagenum: 1
      },
      userlist: [],
      total: 0
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList (Username) {
      const { data: res } = await this.$http.get('users',
        { params: this.queryInfo },
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Status !== 200) return this.$message.error('获取用户列表失败')
      this.userlist = res.Data
      this.total = res.Total
    },
    // 监听
    handleSizeChange (newSize) {
      this.queryInfo.pagesize = newSize
      this.getUserList()
    },
    handleCurrentChange (newPage) {
      this.queryInfo.pagenum = newPage
      this.getUserList()
    },
    async userStateChanged (StateChangedUser) {
      const { data: res } = await this.$http.post('users/editStatus',
        { username: StateChangedUser },
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Status !== 200) return this.$message.error('修改用户状态失败')
      if (res.Status === 200) return this.$message.success('用户状态更新成功')
    }
  }
}
</script>

<style lang="less" scoped>

</style>
