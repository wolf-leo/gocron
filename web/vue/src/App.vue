<template>
  <el-container>
    <el-header>
      <app-nav-menu></app-nav-menu>
    </el-header>
    <el-main>
      <div id="main-container" v-cloak>
        <router-view/>
      </div>
    </el-main>
    <el-footer>
      <app-footer></app-footer>
    </el-footer>
  </el-container>
</template>

<script>
import installService from './api/install'
import appHeader from './components/common/header.vue'
import appNavMenu from './components/common/navMenu.vue'
import appFooter from './components/common/footer.vue'

export default {
  name: 'App',
  data () {
    return {}
  },
  created () {
    installService.status((data) => {
      if (!data) {
        this.$router.push('/install')
      }
    })
  },
  components: {
    appHeader,
    appNavMenu,
    appFooter
  }
}
</script>
<style>
[v-cloak] {
  display: none !important;
}

body {
  margin: 0;
}

.el-header {
  padding: 0;
  margin: 0;
}

.el-container {
  padding: 0;
  margin: 0;
  width: 100%;
  height: 100vh;
}

.el-main {
  padding: 0;
  margin: 0;
}

#main-container .el-main {
  height: 100vh;
  padding: 20px;
  position: relative;
  width: 100%;
  overflow: auto;
}

header {
  height: auto !important;
}

footer {
  height: auto !important;
  background: #545c64;
}

.el-aside {
  width: 120px !important;
  height: 100%;
  display: block;
  z-index: 99999;
}

.el-aside .el-menu {
  height: 100%;
  text-align: center;
}

.el-table__empty-block {
  min-height: calc(100vh / 3);
}

.el-menu.el-menu--horizontal {
  border: none;
}

@media screen and (max-width: 760px) {
  header .el-col {
    width: 20vw;
    text-align: center;
  }
}

@media screen and (min-width: 760px) {
  header .el-col {
    width: 10vw;
    text-align: center;
  }
}

.el-table td.el-table__cell div {
  padding: 0 7px;
}
.el-pagination{
  padding: 10px 0;
}
</style>
