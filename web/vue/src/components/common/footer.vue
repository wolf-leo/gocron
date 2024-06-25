<template>
  <div class="footer">
    <span class="version_text">当前版本：v2.0.6</span>
  </div>
</template>
<style>
.el-footer {
  padding: 0;
}

.footer {
  color: #333333;
}

.footer .version_text {
  text-align: right;
  color: #545c64;
  padding: 2px 0;
  font-size: .90rem;
  position: fixed;
  width: 100%;
  background: #FFFFFF;
  bottom: 0;
  right: 35px;
  z-index: 9999;
}
</style>
<script>
export default {
  name: 'app-footer',
  data () {
    return {}
  }
}
</script>
