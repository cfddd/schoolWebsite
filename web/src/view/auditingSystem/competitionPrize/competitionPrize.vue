<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="学号" prop="student_id">
          <el-input v-model="searchInfo.student_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="学生姓名" prop="student_name">
          <el-input v-model="searchInfo.student_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="获奖时间" prop="award_time">

          <template #label>
            <span>
              获奖时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startAward_time" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endAward_time ? time.getTime() > searchInfo.endAward_time.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endAward_time" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startAward_time ? time.getTime() < searchInfo.startAward_time.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="审核状态" prop="student_name">
          <el-input v-model="searchInfo.status" placeholder="搜索审核状态" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openExportDialog">导出为excel</el-button>
        <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="审核状态" prop="status" width="120" />
        <el-table-column align="left" label="学号" prop="student_id" width="120" />
        <el-table-column align="left" label="学生姓名" prop="student_name" width="120" />
        <el-table-column align="left" label="比赛名称" prop="competition_name" width="120" />
        <el-table-column align="left" label="获奖时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.award_time) }}</template>
        </el-table-column>
        <el-table-column align="left" label="获奖类别" prop="award_type" width="120" />
        <el-table-column align="left" label="获奖等级" prop="award_level" width="120" />
        <el-table-column align="left" label="比赛类型" prop="competition_type" width="120" />
        <el-table-column align="left" label="说明" prop="description" width="120" />
        <el-table-column align="left" label="操作" fixed="right" width="350">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateCompetitionPrizeFunc(scope.row)">审核</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除（不可恢复）</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="exportFormVisible" :before-close="closeExportDialog" :title="'导出'" destroy-on-close>

    </el-dialog>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="'审核'" destroy-on-close>
      <el-scrollbar height="500px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="学号">
            {{ formData.student_id }}
          </el-descriptions-item>
          <el-descriptions-item label="学生姓名">
            {{ formData.student_name }}
          </el-descriptions-item>
          <el-descriptions-item label="比赛名称">
            {{ formData.competition_name }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖时间">
            {{ formatDate(formData.award_time) }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖类别">
            {{ formData.award_type }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖等级">
            {{ formData.award_level }}
          </el-descriptions-item>
          <el-descriptions-item label="比赛类型">
            {{ formData.competition_type }}
          </el-descriptions-item>
          <el-descriptions-item label="说明">
            {{ formData.description }}
          </el-descriptions-item>
          <el-descriptions-item label="审核状态">
            {{ formData.status }}
          </el-descriptions-item>
        </el-descriptions>
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
          <el-form-item label="是否通过:" prop="description">
            <el-select v-model="formData.status" placeholder="未审核">
              <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="审核意见:" prop="description">
            <el-input type="textarea" :autosize="{ minRows: 3, maxRows: 4 }" placeholder="请输入内容" v-model="auditOpinion">
            </el-input>
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="学号">
            {{ formData.student_id }}
          </el-descriptions-item>
          <el-descriptions-item label="学生姓名">
            {{ formData.student_name }}
          </el-descriptions-item>
          <el-descriptions-item label="比赛名称">
            {{ formData.competition_name }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖时间">
            {{ formatDate(formData.award_time) }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖类别">
            {{ formData.award_type }}
          </el-descriptions-item>
          <el-descriptions-item label="获奖等级">
            {{ formData.award_level }}
          </el-descriptions-item>
          <el-descriptions-item label="比赛类型">
            {{ formData.competition_type }}
          </el-descriptions-item>
          <el-descriptions-item label="说明">
            {{ formData.description }}
          </el-descriptions-item>
          <el-descriptions-item label="审核状态">
            {{ formData.status }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCompetitionPrize,
  deleteCompetitionPrize,
  deleteCompetitionPrizeByIds,
  updateCompetitionPrize,
  findCompetitionPrize,
  getCompetitionPrizeList
} from '@/api/competitionPrize'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CompetitionPrize'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  student_id: '',
  student_name: '',
  competition_name: '',
  award_time: new Date(),
  award_type: '',
  award_level: '',
  competition_type: '',
  description: '',
  status: ''
})
// 是否通过
const options = ref(
  [{
    value: '选项1',
    label: '通过'
  }, {
    value: '选项2',
    label: '不通过'
  }],
)
// 审核意见
const auditOpinion = ref('')

// 验证规则
const rule = reactive({
  student_id: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  student_name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  competition_name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  award_type: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  award_level: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
  award_time: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startAward_time && !searchInfo.value.endAward_time) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startAward_time && searchInfo.value.endAward_time) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startAward_time && searchInfo.value.endAward_time && (searchInfo.value.startAward_time.getTime() === searchInfo.value.endAward_time.getTime() || searchInfo.value.startAward_time.getTime() > searchInfo.value.endAward_time.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getCompetitionPrizeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCompetitionPrizeFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除 
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteCompetitionPrizeByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 多选导出
const onOutput = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要导出的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteCompetitionPrizeByIds({ ids })  //todo
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '导出成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    exportFormVisible.value = false
    getTableData()
  }
}

// 更新行
const updateCompetitionPrizeFunc = async (row) => {
  const res = await findCompetitionPrize({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reCP
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCompetitionPrizeFunc = async (row) => {
  const res = await deleteCompetitionPrize({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCompetitionPrize({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reCP
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    student_id: '',
    student_name: '',
    competition_name: '',
    award_time: new Date(),
    award_type: '',
    award_level: '',
    competition_type: '',
    description: '',
  }
}


// 打开弹窗
const openDialog = () => {
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    student_id: '',
    student_name: '',
    competition_name: '',
    award_time: new Date(),
    award_type: '',
    award_level: '',
    competition_type: '',
    description: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCompetitionPrize(formData.value)
        break
      case 'update':
        res = await updateCompetitionPrize(formData.value)
        break
      default:
        res = await createCompetitionPrize(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '审核成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 导出框控制标记
const exportFormVisible = ref(false)
// 打开导出框
const openExportDialog = () => {
  exportFormVisible.value = true
}

// 关闭导出框
const closeExportDialog = () => {
  exportFormVisible.value = false
}
</script>
<style></style>
