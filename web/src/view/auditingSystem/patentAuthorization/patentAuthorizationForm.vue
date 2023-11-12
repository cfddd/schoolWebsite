<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="学号:" prop="student_id">
          <el-input v-model="formData.student_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学号:" prop="student_name">
          <el-input v-model="formData.student_name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专利名称:" prop="patent_name">
          <el-input v-model="formData.patent_name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专利类别:" prop="patent_type">
          <el-input v-model="formData.patent_type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专利授权号:" prop="patent_grant_id">
          <el-input v-model="formData.patent_grant_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="获批时间:" prop="approval_time">
          <el-date-picker v-model="formData.approval_time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="是否第一发明人:" prop="is_first_invent">
          <el-switch v-model="formData.is_first_invent" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPatentAuthorization,
  updatePatentAuthorization,
  findPatentAuthorization
} from '@/api/patentAuthorization'

defineOptions({
    name: 'PatentAuthorizationForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            student_id: '',
            student_name: '',
            patent_name: '',
            patent_type: '',
            patent_grant_id: '',
            approval_time: new Date(),
            is_first_invent: false,
        })
// 验证规则
const rule = reactive({
               student_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               student_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               patent_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               patent_grant_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               is_first_invent : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPatentAuthorization({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rePA
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPatentAuthorization(formData.value)
               break
             case 'update':
               res = await updatePatentAuthorization(formData.value)
               break
             default:
               res = await createPatentAuthorization(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
