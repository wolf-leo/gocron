<template>
<el-container>
  <task-sidebar></task-sidebar>
  <el-main>
    <el-form :inline="true" ref="searchParams" :model="searchParams">
      <el-row>
        <el-form-item label="任务ID" prop="id" size="small">
          <el-input v-model.trim="searchParams.id"></el-input>
        </el-form-item>
        <el-form-item label="任务名称" prop="name" size="small">
          <el-input v-model.trim="searchParams.name"></el-input>
        </el-form-item>
        <el-form-item label="标签" prop="tag" size="small">
          <el-input v-model.trim="searchParams.tag"></el-input>
        </el-form-item>
        <el-form-item label="执行方式" prop="protocol" size="small">
          <el-select v-model.trim="searchParams.protocol">
            <el-option label="全部" value=""></el-option>
            <el-option
              v-for="item in protocolList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
      </el-row>
      <el-row>
        <el-form-item label="任务节点" prop="host_id" size="small">
          <el-select v-model.trim="searchParams.host_id">
            <el-option label="全部" value=""></el-option>
            <el-option
              v-for="item in hosts"
              :key="item.id"
              :label="item.alias + ' - ' + item.name + ':' + item.port "
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status" size="small">
          <el-select v-model.trim="searchParams.status">
            <el-option label="全部" value=""></el-option>
            <el-option
              v-for="item in statusList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <span><el-button type="primary" size="small" @click="search()">搜索</el-button></span>
        <span><el-button type="info" size="small" @click="resetForm">重置</el-button></span>
      </el-row>
    </el-form>
    <el-row type="flex" justify="right">
      <el-col>
        <el-button type="success" @click="changeStyle" size="small">切换样式</el-button>
        <el-button type="primary" v-if="this.$store.getters.user.isAdmin" @click="toEdit(null)" size="small">新增</el-button>
        <el-button type="info" @click="refresh" size="small">刷新</el-button>
      </el-col>
    </el-row>
    <el-pagination
      background
      layout="prev, pager, next, sizes, total"
      :total="taskTotal"
      :page-size="20"
      :current-page="searchParams.page"
      @size-change="changePageSize"
      @current-change="changePage"
      @prev-click="changePage"
      @next-click="changePage">
    </el-pagination>
    <template v-if="listStyle==='2'">
      <el-row :gutter="10">
        <el-col :span="6" v-for="(o, index) in tasks.length" :key="o">
          <el-card class="task_card" :body-style="{ padding: '10px' }">
            <div class="task_descriptions">
              <el-descriptions :title="tasks[index].name" direction="vertical" :column="3" border >
                <el-descriptions-item label="任务ID" :span="1" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">{{ tasks[index].id }}</el-descriptions-item>
                <el-descriptions-item label="cron表达式" :span="2" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">{{ tasks[index].spec }}</el-descriptions-item>
                <el-descriptions-item label="执行方式" :span="1" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">{{ tasks[index] | formatProtocol }}</el-descriptions-item>
                <el-descriptions-item label="状态" :span="1" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">
                  <template>
                    <el-switch
                      v-if="tasks[index].level === 1"
                      v-model="tasks[index].status"
                      :active-value="1"
                      :inactive-vlaue="0"
                      active-color="#13ce66"
                      @change="changeStatus(tasks[index])"
                      inactive-color="#ff4949">
                    </el-switch>
                  </template>
                </el-descriptions-item>
                <el-descriptions-item label="下次执行" :span="3" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">
                  <div class="descriptions-tooltip">
                    <el-tooltip class="item" effect="dark" :content="tasks[index].next_run_time | formatTime" placement="top-start">
                      <el-button type="text"> {{ tasks[index].next_run_time | formatTime }}</el-button>
                    </el-tooltip>
                  </div>
                </el-descriptions-item>
                <el-descriptions-item label="标签" :span="1" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">
                  <div class="descriptions-tooltip">
                    <el-tooltip class="item" effect="dark" :content="tasks[index].tag || '-'" placement="top-start">
                      <el-button type="text"> {{ tasks[index].tag || '-' }}</el-button>
                    </el-tooltip>
                  </div>
                </el-descriptions-item>
                <el-descriptions-item label="备注" :span="2" content-class-name="el-descriptions-item-content" label-class-name="el-descriptions-item-label">
                  <div class="descriptions-tooltip">
                    <el-tooltip class="item" effect="dark" :content="tasks[index].remark || '-'" placement="top-start">
                      <el-button type="text"> {{ tasks[index].remark || '-' }}</el-button>
                    </el-tooltip>
                  </div>
                </el-descriptions-item>
              </el-descriptions>
            </div>
            <el-row type="flex" class="row-bg" justify="center">
              <el-button-group>
                <el-button type="primary" @click="toEdit(tasks[index])" size="small">编辑</el-button>
                <el-button type="success" @click="runTask(tasks[index])" size="small">手动执行</el-button>
                <el-button type="info" @click="jumpToLog(tasks[index])" size="small">查看日志</el-button>
                <el-button type="danger" @click="remove(tasks[index])" size="small">删除</el-button>
              </el-button-group>
            </el-row>
          </el-card>
        </el-col>
      </el-row>
    </template>
    <template v-else>
      <el-table
        :data="tasks"
        tooltip-effect="dark"
        border
        style="width: 100%">
        <el-table-column type="expand">
          <template slot-scope="scope">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="任务创建时间:">
                {{ scope.row.created | formatTime }} <br>
              </el-form-item>
              <el-form-item label="任务类型:">
                {{ scope.row.level | formatLevel }} <br>
              </el-form-item>
              <el-form-item label="单实例运行:">
                {{ scope.row.multi | formatMulti }} <br>
              </el-form-item>
              <el-form-item label="超时时间:">
                {{ scope.row.timeout | formatTimeout }} <br>
              </el-form-item>
              <el-form-item label="重试次数:">
                {{ scope.row.retry_times }} <br>
              </el-form-item>
              <el-form-item label="重试间隔:">
                {{ scope.row.retry_interval | formatRetryTimesInterval }}
              </el-form-item>
              <br>
              <el-form-item label="任务节点">
                <div v-for="item in scope.row.hosts" :key="item.host_id">
                  {{ item.alias }} - {{ item.name }}:{{ item.port }} <br>
                </div>
              </el-form-item>
              <br>
              <el-form-item label="命令:" style="width: 100%">
                {{ scope.row.command }}
              </el-form-item>
              <br>
              <el-form-item label="备注" style="width: 100%">
                {{ scope.row.remark }}
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column
          prop="id"
          min-width="10"
          align="center"
          label="任务ID">
        </el-table-column>
        <el-table-column
          prop="name"
          label="任务名称"
          min-width="30"
          align="center">
        </el-table-column>
        <el-table-column
          prop="tag"
          label="标签"
          min-width="10" align="center">
        </el-table-column>
        <el-table-column
          prop="spec"
          label="cron表达式"
          min-width="15"
          align="center">
        </el-table-column>
        <el-table-column label="下次执行时间" min-width="15" align="center">
          <template slot-scope="scope">
            {{ scope.row.next_run_time | formatTime }}
          </template>
        </el-table-column>
        <el-table-column
          prop="protocol"
          :formatter="formatProtocol"
          min-width="10"
          label="执行方式" align="center">
        </el-table-column>
        <el-table-column
          min-width="10" align="center"
          label="状态" v-if="this.isAdmin">
          <template slot-scope="scope">
            <el-switch
              v-if="scope.row.level === 1"
              v-model="scope.row.status"
              :active-value="1"
              :inactive-vlaue="0"
              active-color="#13ce66"
              @change="changeStatus(scope.row)"
              inactive-color="#ff4949">
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="120" align="center" v-else>
          <template slot-scope="scope">
            <el-switch
              v-if="scope.row.level === 1"
              v-model="scope.row.status"
              :active-value="1"
              :inactive-vlaue="0"
              active-color="#13ce66"
              :disabled="true"
              inactive-color="#ff4949">
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="25" align="center" v-if="this.isAdmin">
          <template slot-scope="scope">
            <el-row>
              <el-button-group>
                <el-button type="primary" @click="toEdit(scope.row)" size="small">编辑</el-button>
                <el-button type="success" @click="runTask(scope.row)" size="small">手动执行</el-button>
                <el-button type="info" @click="jumpToLog(scope.row)" size="small">查看日志</el-button>
                <el-button type="danger" @click="remove(scope.row)" size="small">删除</el-button>
              </el-button-group>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <el-pagination
      background
      layout="prev, pager, next, sizes, total"
      :total="taskTotal"
      :page-size="20"
      :current-page="searchParams.page"
      @size-change="changePageSize"
      @current-change="changePage"
      @prev-click="changePage"
      @next-click="changePage">
    </el-pagination>
    <el-dialog :title="dialogData.name" :visible.sync="dialogVisible">
      <el-table :data="[dialogData]">
        <el-table-column property="tag" label="标签"></el-table-column>
        <el-table-column property="remark" label="备注"></el-table-column>
      </el-table>
    </el-dialog>
  </el-main>
</el-container>
</template>

<script>
import taskSidebar from './sidebar'
import taskService from '../../api/task'

export default {
  name: 'task-list',
  data () {
    return {
      tasks: [],
      hosts: [],
      dialogData: [],
      listStyle: localStorage.getItem('taskLogStyle') || '1',
      taskTotal: 0,
      dialogVisible: false,
      searchParams: {
        page_size: 20,
        page: 1,
        id: '',
        protocol: '',
        name: '',
        tag: '',
        host_id: '',
        status: ''
      },
      isAdmin: this.$store.getters.user.isAdmin,
      protocolList: [
        {
          value: '1',
          label: 'http'
        },
        {
          value: '2',
          label: 'shell'
        }
      ],
      statusList: [
        {
          value: '2',
          label: '激活'
        },
        {
          value: '1',
          label: '停止'
        }
      ]
    }
  },
  components: {taskSidebar},
  created () {
    const hostId = this.$route.query.host_id
    if (hostId) {
      this.searchParams.host_id = hostId
    }

    this.search()
  },
  filters: {
    formatLevel (value) {
      if (value === 1) {
        return '主任务'
      }
      return '子任务'
    },
    formatTimeout (value) {
      if (value > 0) {
        return value + '秒'
      }
      return '不限制'
    },
    formatRetryTimesInterval (value) {
      if (value > 0) {
        return value + '秒'
      }
      return '系统默认'
    },
    formatProtocol (row) {
      if (row.protocol === 2) {
        return 'shell'
      }
      if (row.http_method === 1) {
        return 'http-get'
      }
      return 'http-post'
    },
    formatMulti (value) {
      if (value > 0) {
        return '否'
      }
      return '是'
    }
  },
  methods: {
    changeStatus (item) {
      if (item.status) {
        taskService.enable(item.id)
      } else {
        taskService.disable(item.id)
      }
    },
    formatProtocol (row, col) {
      if (row[col.property] === 2) {
        return 'shell'
      }
      if (row.http_method === 1) {
        return 'http-get'
      }
      return 'http-post'
    },
    changePage (page) {
      this.searchParams.page = page
      this.search()
    },
    changePageSize (pageSize) {
      this.searchParams.page_size = pageSize
      this.search()
    },
    search (callback = null) {
      taskService.list(this.searchParams, (tasks, hosts) => {
        this.tasks = tasks.data
        this.taskTotal = tasks.total
        this.hosts = hosts
        if (callback) {
          callback()
        }
      })
    },
    resetForm () {
      this.$refs['searchParams'].resetFields()
      this.refresh()
    },
    runTask (item) {
      this.$appConfirm(() => {
        taskService.run(item.id, () => {
          this.$message.success('任务已开始执行')
        })
      }, true)
    },
    remove (item) {
      this.$appConfirm(() => {
        taskService.remove(item.id, () => {
          this.refresh()
        })
      })
    },
    jumpToLog (item) {
      this.$router.push(`/task/log?task_id=${item.id}`)
    },
    refresh () {
      this.search(() => {
        this.$message.success('刷新成功')
      })
    },
    changeStyle () {
      let nowStyle = this.listStyle === '1' ? '2' : '1'
      localStorage.setItem('taskLogStyle', nowStyle.toString())
      this.listStyle = nowStyle
    },
    showMore (item) {
      this.dialogData = item
      this.dialogVisible = true
    },
    toEdit (item) {
      let path = ''
      if (item === null) {
        path = '/task/create'
      } else {
        path = `/task/edit/${item.id}`
      }
      this.$router.push(path)
    }
  }
}
</script>

<style scoped>

.demo-table-expand {
  font-size: 0;
}

.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}

.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}

.task_card {
  background-color: #f9fafc;
  margin-bottom: 10px;
}

.row-bg {
  padding-top: 10px;
}

.task_spec {
  height: 25px;
  background: #242424;
  color: #FFFFFF;
  padding: 10px;
  border-radius: 5px;
  word-break: break-all;
  font-size: 0.90rem;
}

.descriptions-tooltip {
  height: 30px;
}

.el-button--text{
  color: #333333;
}

/deep/ .task_descriptions .el-descriptions__title {
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

/deep/ .el-descriptions .is-bordered {
  table-layout: fixed;
}

/deep/ .el-descriptions-item-content {
  height: 50px !important;
  line-height: 1 !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

</style>
