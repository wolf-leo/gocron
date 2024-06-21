<template>
  <el-aside>
    <el-menu
      :default-active="currentRoute"
      mode="vertical"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router>
      <el-menu-item index="/task">🕘定时任务</el-menu-item>
      <el-menu-item index="/task/log">📖任务日志</el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'task-sidebar',
  data () {
    return {}
  },
  computed: {
    currentRoute () {
      if (this.$route.path === '/task/log') {
        return '/task/log'
      }
      return '/task'
    }
  }
}
</script>
