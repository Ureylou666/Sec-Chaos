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
          <el-input placeholder="请输入想要查询的用户" v-model="queryInfo.query" clearable @clear="getUserList" @keyup.enter.native="getUserList">
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="5">
          <el-button type="primary" @click="addDialogVisible = true">添加用户</el-button>
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
        <!-- 用户操作区 -->
        <el-table-column label="操作" width="300px">
          <template slot-scope="scope">
            <!-- 用户编辑按钮 -->
            <el-button type="primary" icon="el-icon-edit" size="medium" @click="editvisible(scope.row.Username)"></el-button>
            <!-- 删除用户弹框 -->
            <el-popover placement="top" width="auto" :v-model="popoverShow" :v-show="popoverID == scope.$index">
              <p>是否确定删除用户：{{ UserInfo.username }}</p>
              <div style="text-align: right; margin: 0">
                <el-button type="text" size="mini" @click="popoverID = -1;">取 消</el-button>
                <el-button type="primary" size="mini" @click="popoverID = -1;userDelete(UserInfo.username);">确 定</el-button>
              </div>
              <el-button slot="reference" type="danger" icon="el-icon-delete" size="medium" @click="setUsername(scope.row.Username)"></el-button>
            </el-popover>
            <!-- 用户角色配置弹框 -->
            <el-tooltip class="item" effect="dark" content="角色分配" placement="top" :enterable="false">
              <el-button type="warning" icon="el-icon-setting" size="medium" @click="getRolelist();roleDialogVisible = true"></el-button>
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
    <!-- 添加用户对话框 -->
    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed" :close-on-click-modal="false">
      <!-- 内容主题区域 -->
      <el-form :model="UseraddForm" :rules="UseraddRules" ref="UseraddForm" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="UseraddForm.username" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="UseraddForm.password" prefix-icon="el-icon-lock"
                    type="password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkpass">
          <el-input v-model="UseraddForm.checkpass" prefix-icon="el-icon-lock"
                    type="password" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <!-- 内容底部区域 -->
      <span slot="footer" class="dialog-footer">
    <el-button @click="addDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="addDialogVisible = false;AddUser('UseraddForm')">确 定</el-button>
  </span>
    </el-dialog>
    <!-- 编辑用户弹框 -->
    <el-dialog title="编辑用户" :visible.sync="editDialogvisible" width="50%" @close="editDialogClosed" :close-on-click-modal="false">
      <el-form :model="UsereditForm" :rules="UsereditRules" ref="UsereditForm" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="UsereditForm.username" autocomplete="off" disabled></el-input>
        </el-form-item>
        <el-form-item label="原密码" prop="OldPassword">
          <el-input type="password" v-model="UsereditForm.OldPassword" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="新用户密码" prop="NewPassword">
          <el-input type="password" v-model="UsereditForm.NewPassword" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="再输入一遍" prop="CheckPassword">
          <el-input type="password" v-model="UsereditForm.CheckPassword" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <!-- 内容底部区域 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogClosed()">取 消</el-button>
        <el-button type="primary" @click="editUser('UsereditForm');editDialogClosed()">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 设置用户权限区域 -->
    <el-dialog title="用户角色" :visible.sync="roleDialogVisible" width="30%">
      <el-select v-model="rolelist" clearable placeholder="请选择">
        <el-option
          v-for="item in rolelist"
          :key="item.RoleUID"
          :label="item.RoleName"
          :value="item.RoleUID">
        </el-option>
      </el-select>
      <!-- 内容底部区域 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="roleDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="roleDialogVisible = false">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data () {
    // 验证两次密码是否相同
    var checkPassword = (rule, value, callback) => {
      if (value !== this.UseraddForm.password) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
    var checkPassword2 = (rule, value, callback) => {
      if (value !== this.UsereditForm.NewPassword) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
    return {
      // 获取用户列表参数
      queryInfo: {
        query: '',
        pagesize: 5,
        pagenum: 1
      },
      UseraddForm: {
        username: '',
        password: '',
        checkpassword: ''
      },
      UsereditForm: {
        username: '',
        OldPassword: '',
        NewPassword: '',
        CheckPassword: ''
      },
      UserInfo: {
        username: ''
      },
      // 添加用户表单验证规则
      UseraddRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名长度为4-12个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '密码长度需要为8-20位', trigger: 'blur' }
        ],
        checkpass: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '密码长度需要为8-20位', trigger: 'blur' },
          { validator: checkPassword, trigger: 'blur' }
        ]
      },
      UsereditRules: {
        NewPassword: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '密码长度需要为8-20位', trigger: 'blur' }
        ],
        CheckPassword: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '密码长度需要为8-20位', trigger: 'blur' },
          { validator: checkPassword2, trigger: 'blur' }
        ]
      },
      userlist: [],
      rolelist: [],
      total: 0,
      // 控制添加用户对话框
      addDialogVisible: false,
      // 删除用户提示框
      popoverID: -1,
      popoverShow: true,
      // 控制编辑用户对话框
      editDialogvisible: false,
      // 控制用户权限
      roleDialogVisible: false
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.post('users', this.queryInfo,
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
    },
    AddUser (formName) {
      this.$refs[formName].validate(async (valid) => {
        // 判断是否符合规范
        if (valid) {
          const { data: res } = await this.$http.post('users/add',
            {
              username: this.UseraddForm.username,
              password: this.UseraddForm.password
            },
            { headers: { 'Content-Type': 'application/json' } })
          if (res.Status !== 200) {
            return this.$message.error(res.Message)
          }
          if (res.Status === 200) {
            this.addDialogVisible = false
            console.log(res)
            return this.$message.success(res.Message)
          }
        } else {
          return this.$message.error('添加用户失败')
        }
      })
    },
    addDialogClosed () {
      this.$refs.UseraddForm.resetFields()
      this.getUserList()
    },
    editUser (formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          const { data: res } = await this.$http.post('users/editPassword',
            {
              Username: this.UsereditForm.username,
              OldPassword: this.UsereditForm.OldPassword,
              NewPassword: this.UsereditForm.NewPassword
            },
            { headers: { 'Content-Type': 'application/json' } })
          if (res.Status !== 200) {
            return this.$message.error(res.Message)
          }
          if (res.Status === 200) {
            this.$message.success('用户密码修改成功')
            this.getUserList()
          }
        } else {
          this.$message.error('用户密码修改失败')
        }
      })
    },
    editvisible (username) {
      this.UsereditForm.username = username
      this.editDialogvisible = true
    },
    editDialogClosed () {
      this.editDialogvisible = false
      this.$refs.UsereditForm.resetFields()
      this.getUserList()
    },
    // 传入删除用户名
    setUsername (Username) {
      this.UserInfo.username = Username
    },
    async userDelete (Username) {
      const { data: res } = await this.$http.post('users/delete',
        {
          username: Username
        },
        { headers: { 'Content-Type': 'application/json' } })
      if (res.Status !== 200) {
        return this.$message.error(res.Message)
      }
      if (res.Status === 200) {
        this.$message.success('用户已删除')
        this.getUserList()
      }
    },
    async getRolelist () {
      const { data: res } = await this.$http.get('roles', { headers: { 'Content-Type': 'application/json' } })
      this.rolelist = res.data
    }
  }
}
</script>

<style lang="less" scoped>
.el-button {
  margin-right: 10px;
}
</style>
