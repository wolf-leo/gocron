<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="width: 500px;">
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
  name: 'user-edit-password',
  data: function () {
    return {
      dialogVisible: false,
      form: {
        id: '',
        new_password: '',
        confirm_new_password: ''
      },
      formRules: {
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
  created () {
    const id = this.$route.params.id
    if (!id) {
      return
    }
    this.form.id = id
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
        this.save()
      })
    },
    save () {
      userService.editPassword(this.form, (e, code, msg) => {
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
    cancel () {
      this.$router.push('/user')
    }
  }
}
</script>
