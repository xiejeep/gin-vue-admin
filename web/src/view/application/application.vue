<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="应用名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="应用名称" prop="name" width="120" />
        <el-table-column align="left" label="API密钥" prop="apiKey" width="200">
          <template #default="scope">
            <span v-if="!scope.row.showApiKey">{{ maskApiKey(scope.row.apiKey) }}</span>
            <span v-else>{{ scope.row.apiKey }}</span>
            <el-button
              type="text"
              size="small"
              @click="toggleApiKeyVisibility(scope.row)"
            >
              {{ scope.row.showApiKey ? '隐藏' : '显示' }}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column align="left" label="基础URL" prop="baseUrl" width="200" />
        <el-table-column align="left" label="应用描述" prop="description" width="200" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">查看详情</el-button>
            <el-button type="primary" link class="table-button" @click="updateSysApplicationFunc(scope.row)">变更</el-button>
            <el-button type="primary" link class="table-button" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="应用名称:" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入应用名称" />
        </el-form-item>
        <el-form-item label="API密钥:" prop="apiKey">
          <el-input v-model="formData.apiKey" clearable placeholder="请输入API密钥" />
        </el-form-item>
        <el-form-item label="基础URL:" prop="baseUrl">
          <el-input v-model="formData.baseUrl" clearable placeholder="请输入基础URL" />
        </el-form-item>
        <el-form-item label="应用描述:" prop="description">
          <el-input v-model="formData.description" clearable placeholder="请输入应用描述" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="应用名称">
          {{ formData.name }}
        </el-descriptions-item>
        <el-descriptions-item label="API密钥">
          {{ formData.apiKey }}
        </el-descriptions-item>
        <el-descriptions-item label="基础URL">
          {{ formData.baseUrl }}
        </el-descriptions-item>
        <el-descriptions-item label="应用描述">
          {{ formData.description }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="formData.status === 1 ? 'success' : 'danger'">
            {{ formData.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysApplication,
  deleteSysApplication,
  deleteSysApplicationByIds,
  updateSysApplication,
  findSysApplication,
  getSysApplicationList
} from '@/api/application'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'SysApplication'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  apiKey: '',
  baseUrl: '',
  description: '',
  status: 1,
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  apiKey: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  baseUrl: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const searchInfo = ref({})

// 多选数据
const multipleSelection = ref([])
// 分页
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchText = ref('')
const elSearchFormRef = ref()
const elFormRef = ref()

// 弹窗控制标记
const dialogFormVisible = ref(false)
const detailShow = ref(false)

// 操作类型
const type = ref('')

// 更新时弹窗内容存储
const updateSysApplicationFunc = async (row) => {
  const res = await findSysApplication({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysApplication
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteSysApplicationFunc(row)
  })
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteSysApplicationFunc = async (row) => {
  const res = await deleteSysApplication({ ID: row.ID })
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

// 批量删除控制标记
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteSysApplicationByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const updateFunc = async () => {
  const res = await updateSysApplication(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })
    closeDialog()
    getTableData()
  }
}

// 删除行
const createFunc = async () => {
  const res = await createSysApplication(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建成功'
    })
    closeDialog()
    getTableData()
  }
}

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    switch (type.value) {
      case 'create':
        await createFunc()
        break
      case 'update':
        await updateFunc()
        break
      default:
        ElMessage({
          type: 'error',
          message: '未知操作'
        })
        break
    }
  })
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    apiKey: '',
    baseUrl: '',
    description: '',
    status: 1,
  }
}
// 弹窗确定
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findSysApplication({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.resysApplication
    detailShow.value = true
  }
}

const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    name: '',
    apiKey: '',
    baseUrl: '',
    description: '',
    status: 1,
  }
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 分页
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 获取数据
const getTableData = async () => {
  const table = await getSysApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// API密钥显示控制
const toggleApiKeyVisibility = (row) => {
  row.showApiKey = !row.showApiKey
}

// 遮盖API密钥
const maskApiKey = (apiKey) => {
  if (!apiKey) return ''
  const visibleLength = 8
  if (apiKey.length <= visibleLength) {
    return '*'.repeat(apiKey.length)
  }
  return apiKey.substring(0, 4) + '*'.repeat(apiKey.length - visibleLength) + apiKey.substring(apiKey.length - 4)
}

getTableData()
</script>

<style>
</style>