<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="max-width: 1200px;">
        <el-form-item>
          <el-input v-model="form.id" type="hidden"></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="name">
          <el-input v-model="form.name" :disabled="form.id>0"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <template v-if="!form.id">
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" type="password"></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="confirm_password">
            <el-input v-model="form.confirm_password" type="password"></el-input>
          </el-form-item>
        </template>
        <el-form-item label="角色" prop="is_admin">
          <el-radio-group v-model="form.is_admin">
            <el-radio :label="0">普通用户</el-radio>
            <el-radio :label="1">管理员</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="授权任务" v-if="tasks[0]">
          <el-checkbox :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">全选 (管理员无视授权)</el-checkbox>
          <el-checkbox-group v-model="checkedtasks" @change="handleCheckedtasksChange" size="small">
            <el-checkbox border v-for="taskIds in tasks" :label="taskIds.id" :key="taskIds.id" :checked="taskIds.checked">{{ taskIds.name }} ({{ taskIds.id }})</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submit()">保存</el-button>
          <el-button @click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import userService from '../../api/user'

export default {
  name: 'user-edit',
  data: function () {
    return {
      checkAll: false,
      checkedtasks: [],
      tasks: [],
      isIndeterminate: true,
      form: {
        id: '',
        name: '',
        email: '',
        is_admin: 0,
        password: '',
        confirm_password: '',
        status: 1
      },
      formRules: {
        name: [
          {required: true, message: '请输入用户名', trigger: 'blur'}
        ],
        email: [
          {type: 'email', required: true, message: '请输入有效邮箱地址', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'}
        ],
        confirm_password: [
          {required: true, message: '请再次输入密码', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    const id = this.$route.params.id
    if (!id) {
      return
    }
    userService.detail(id, (data) => {
      if (!data) {
        this.$message.error('数据不存在')
        return
      }
      let taskIdsData = data.task_ids === '' ? [] : JSON.parse(data.task_ids)
      let taskIdsLen = Object.keys(taskIdsData).length
      taskIdsData = JSON.parse(JSON.stringify(taskIdsData))
      let checkedLen = 0
      for (const taskIdsLenKey in taskIdsData) {
        if (taskIdsData[taskIdsLenKey]['checked']) checkedLen++
      }
      this.checkAll = checkedLen === taskIdsLen;
      this.isIndeterminate = checkedLen > 0 && checkedLen < taskIdsLen;
      this.form.id = data.id
      this.form.name = data.name
      this.form.email = data.email
      this.form.is_admin = data.is_admin
      this.form.status = data.status
      this.tasks = taskIdsData
    })
  },
  methods: {
    handleCheckAllChange(val) {
      let ids = [];
      let tasks = JSON.parse(JSON.stringify(this.tasks))
      for (let i = 0; i < Object.keys(tasks).length; i++) {
        ids.push(tasks[i]['id']);
      }
      this.checkedtasks = val ? ids : [];
      this.isIndeterminate = false;
    },
    handleCheckedtasksChange(value) {
      let tasks = JSON.parse(JSON.stringify(this.tasks))
      let checkedCount = value.length;
      let taskCount = Object.keys(tasks).length;
      this.checkAll = checkedCount === taskCount;
      this.isIndeterminate = checkedCount > 0 && checkedCount < taskCount;
    },
    submit() {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        this.save()
      })
    },
    save() {
      let checkedTaskIds = JSON.parse(JSON.stringify(this.checkedtasks))
      this.form.task_ids = JSON.stringify(checkedTaskIds)
      userService.update(this.form, (e, code, msg) => {
        if (code !== 0) {
          this.$message.error(msg)
          return
        }
        this.$message.success({
          duration: 1000,
          message: '操作成功',
          onClose: () => {
            this.$router.push('/user')
          }
        })
      })
    },
    cancel() {
      this.$router.push('/user')
    }
  }
}
</script>
