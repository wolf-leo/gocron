<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="width: 500px;">
        <el-form-item label="原密码" prop="old_password">
          <el-input v-model="form.old_password" type="password"></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="new_password">
          <el-input v-model="form.new_password" type="password"></el-input>
        </el-form-item>
        <el-form-item label="确认新密码" prop="confirm_new_password">
          <el-input v-model="form.confirm_new_password" type="password"></el-input>
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
  name: 'user-edit-my-password',
  data: function () {
    return {
      form: {
        old_password: '',
        new_password: '',
        confirm_new_password: ''
      },
      formRules: {
        old_password: [
          {required: true, message: '请输入原密码', trigger: 'blur'},
          {min: 1, max: 30, message: '密码长度错误', trigger: 'blur'}
        ],
        new_password: [
          {required: true, message: '请输入新密码', trigger: 'blur'},
          {min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur'}
        ],
        confirm_new_password: [
          {required: true, message: '请再次输入新密码', trigger: 'blur'},
          {min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        if (this.form.confirm_new_password !== this.form.new_password) {
          this.$message.error('两次密码不一致')
          return false
        }
        if (this.form.confirm_new_password === this.form.old_password) {
          this.$message.error('您的新旧密码不能相同')
          return false
        }
        this.save()
      })
    },
    save () {
      userService.editMyPassword(this.form, (e, code, msg) => {
        if (code !== 0) {
          this.$message.error(msg)
          return
        }
        this.$message.success({
          duration: 1000,
          message: '操作成功',
          onClose: () => {
            this.$router.back()
          }
        })
      })
    },
    cancel () {
      this.$router.back()
    }
  }
}
</script>
