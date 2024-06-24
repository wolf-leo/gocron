<template>
  <el-container>
    <el-main>
      <el-form :inline="true" ref="searchParams" :model="searchParams">
        <el-row>
          <el-form-item label="节点ID" size="small" prop="id">
            <el-input v-model.trim="searchParams.id"></el-input>
          </el-form-item>
          <el-form-item label="主机名" size="small" prop="name">
            <el-input v-model.trim="searchParams.name"></el-input>
          </el-form-item>
          <span><el-button type="primary" size="small" @click="search()">搜索</el-button></span>
          <span><el-button type="info" size="small" @click="resetForm">重置</el-button></span>
        </el-row>
      </el-form>
      <el-row type="flex" justify="right">
        <el-col>
          <el-button type="primary" v-if="this.$store.getters.user.isAdmin" @click="toEdit(null)" size="small">新增</el-button>
          <el-button type="info" @click="refresh" size="small">刷新</el-button>
        </el-col>
      </el-row>
      <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="hostTotal"
        :page-size="20"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
      </el-pagination>
      <el-table
        :data="hosts"
        tooltip-effect="dark"
        border
        style="width: 100%">
        <el-table-column
          prop="id"
          label="节点ID" min-width="15" align="center">
        </el-table-column>
        <el-table-column
          prop="alias"
          label="节点名称" min-width="30" align="center">
        </el-table-column>
        <el-table-column
          prop="name" min-width="25"
          label="主机名" align="center">
        </el-table-column>
        <el-table-column
          prop="port"
          label="端口" min-width="15" align="center">
        </el-table-column>
        <el-table-column label="查看任务" min-width="15" align="center">
          <template slot-scope="scope">
            <el-button-group>
              <el-button type="success" @click="toTasks(scope.row)" size="small">查看任务</el-button>
            </el-button-group>
          </template>
        </el-table-column>
        <el-table-column
          prop="remark" min-width="25"
          label="备注" align="center">
        </el-table-column>
        <el-table-column label="操作" align="center" min-width="30" v-if="this.isAdmin">
          <template slot-scope="scope">
            <el-row>
              <el-button-group>
              <el-button type="primary" @click="toEdit(scope.row)" size="small">编辑</el-button>
              <el-button type="info" @click="ping(scope.row)" size="small">测试连接</el-button>
              <el-button type="danger" @click="remove(scope.row)" size="small">删除</el-button>
              </el-button-group>
            </el-row>
            <br>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="hostTotal"
        :page-size="20"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
      </el-pagination>
    </el-main>
  </el-container>
</template>

<script>
import hostService from '../../api/host'
export default {
  name: 'host-list',
  data () {
    return {
      hosts: [],
      hostTotal: 0,
      searchParams: {
        page_size: 20,
        page: 1,
        id: '',
        name: '',
        alias: ''
      },
      isAdmin: this.$store.getters.user.isAdmin
    }
  },
  created () {
    this.search()
  },
  methods: {
    changePage (page) {
      this.searchParams.page = page
      this.search()
    },
    changePageSize (pageSize) {
      this.searchParams.page_size = pageSize
      this.search()
    },
    search (callback = null) {
      hostService.list(this.searchParams, (data) => {
        this.hosts = data.data
        this.hostTotal = data.total
        if (callback) {
          callback()
        }
      })
    },
    resetForm () {
      this.$refs['searchParams'].resetFields()
      this.refresh()
    },
    remove (item) {
      this.$appConfirm(() => {
        hostService.remove(item.id, () => this.refresh())
      })
    },
    ping (item) {
      hostService.ping(item.id, () => {
        this.$message.success('连接成功')
      })
    },
    toEdit (item) {
      let path = ''
      if (item === null) {
        path = '/host/create'
      } else {
        path = `/host/edit/${item.id}`
      }
      this.$router.push(path)
    },
    refresh () {
      this.search(() => {
        this.$message.success('刷新成功')
      })
    },
    toTasks (item) {
      this.$router.push(
        {
          path: '/task',
          query: {
            host_id: item.id
          }
        })
    }
  }
}
</script>
