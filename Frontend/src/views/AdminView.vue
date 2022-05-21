,<template>
  <el-container>
    <el-header>
      <common-header></common-header>
    </el-header>
    <el-container>
      <el-aside width="10%">
        <common-aside></common-aside>
      </el-aside>
      <el-main width="90%">
        Main
      </el-main>
    </el-container>
    <el-footer>
      <common-Footer></common-Footer>
    </el-footer>
  </el-container>
</template>

<script>
import CommonHeader from '@/components/admin/Common/CommonHeader'
import CommonAside from '@/components/admin/Common/CommonAside'
import CommonFooter from '@/components/admin/Common/CommonFooter'
export default {
  name: 'HomeView',
  components: { CommonHeader, CommonAside, CommonFooter },
  data () {
    return {}
  }
}
</script>
