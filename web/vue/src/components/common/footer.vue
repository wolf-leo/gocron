<template>
  <div class="footer">
    <span class="version_text">当前版本：v2.0.2</span>
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
  margin-right: 25px;
  padding: 0;
  font-size: .85rem;
  position: fixed;
  width: 200px;
  background: #FFFFFF;
  bottom: 0;
  right: 0;
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
