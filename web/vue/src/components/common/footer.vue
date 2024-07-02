<template>
  <el-main>
    <el-dialog :visible.sync="dialogVisible" title="gocron - 统一任务调度" class="dialog-contain" width="42%">
      <div class="dialog-box">
        <p>版本号：{{ version }}</p>
        <p> Github：<a href="https://github.com/wolf-leo/gocron" target="_blank">https://github.com/wolf-leo/gocron</a></p>
        <p> Gitee：<a href="https://gitee.com/wolf18/gocron" target="_blank">https://gitee.com/wolf18/gocron</a></p>
      </div>
    </el-dialog>
    <div class="footer">
      <a class="version_text" @click="showDialog">当前版本：{{ version }} </a>
    </div>
  </el-main>
</template>

<script>
export default {
  name: 'app-footer',
  data () {
    return {
      version: 'v2.0.7',
      dialogVisible: false
    }
  },
  methods: {
    showDialog () {
      this.dialogVisible = true
    }
  }
}
</script>

<style>

.el-footer {
  padding: 0;
}

a {
  cursor: pointer;
  text-decoration: none;
}

.footer {
  color: #333333;
  background: #FFFFFF;
  bottom: 0;
  width: 100%;
  height: 25px;
  z-index: 111;
  position: absolute;
}

.footer .version_text {
  text-align: right;
  color: #545c64;
  padding: 2px 0;
  font-size: .90rem;
  right: 15px;
  z-index: 9999;
  position: absolute;
  bottom: 0;
}

.dialog-contain {
  display: block;
}

.dialog-box {
  border: 1px solid #e2e2e2;
  padding: 10px;
  border-radius: 2px;
}

</style>
