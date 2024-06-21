<template>
  <div v-cloak>
    <el-menu
      :default-active="currentRoute"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router>
      <el-row>
        <el-col :span="2">
          <el-menu-item index="/task">📝任务管理</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item index="/host">📆任务节点</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item v-if="this.$store.getters.user.isAdmin" index="/user">⛄用户管理</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item v-if="this.$store.getters.user.isAdmin" index="/system">💻系统管理</el-menu-item>
        </el-col>
        <el-col :span="2" style="float:right;text-align: right;">
          <el-submenu v-if="this.$store.getters.user.token" index="userStatus">
            <template slot="title">{{ this.$store.getters.user.username }}</template>
            <el-menu-item index="/user/edit-my-password">修改密码</el-menu-item>
            <el-menu-item @click="logout" index="/user/logout">退出</el-menu-item>
          </el-submenu>
        </el-col>
      </el-row>
    </el-menu>
  </div>
</template>
<style>
aside .el-menu-item {
  padding: 0 !important;
}
</style>
<script>

export default {
  name: 'app-nav-menu',
  data () {
    return {}
  },
  computed: {
    currentRoute () {
      if (this.$route.path === '/') {
        return '/task'
      }
      const segments = this.$route.path.split('/')
      return `/${segments[1]}`
    }
  },
  methods: {
    logout () {
      this.$store.commit('logout')
      this.$router.push('/')
    }
  }
}
</script>
