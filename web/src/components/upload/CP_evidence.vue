<template>
  <div>     
      <el-upload
              action="' '" 
              list-type="picture-card"
              :limit="3"
              show-file-list
              :auto-upload="false"
              :on-change="change"
              multiple
            >
              <i class="el-icon-plus"></i>
            </el-upload>
            <el-dialog :visible.sync="dialogVisible">
              <img width="100%" :src="dialogImageUrl" alt="" />
        </el-dialog>
   <div class="submitFeedbackBtn" @click="submit">提交</div>
  </div>
</template>
<script>
import axios from "axios"; //数据请求    
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { isVideoMime, isImageMime } from '@/utils/image'
const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

export default {
  data() {
    return {
      dataList: "",
      dialogImageUrl: "",
    };
  },
 methods: {
    submit() {
      let config = {
        timeout: 30000,
        headers: {
          "Content-Type": "multipart/form-data", //设置headers
          "x-Token": userStore.token,
        },
      };
     const formData = new FormData();

      var that = this;
    // 首先判断是否上传了图片，如果上传了图片，将图片存入到formData中
      console.log(this.dataList);
      if (this.dataList) {
        that.dataList.forEach((item, index) => {
          // console.log(item)
          formData.append("uploads",item);
        });
      }
  // console.log(formData.get(0));
         axios.post(
          "`${path}/CP/createCompetitionPrize`", //请求后端的url
            formData,
            config
          )
          .then((res) => {
            console.log(res)
            if (res.data.status == 'ok') {
              //上传成功
              console.log("上传成功");
              this.$router.push({
                path:'./'
              })
            } else {
              alert('上传失败')
            }
          })
          .catch((error) => {
            console.log("请求失败");
          });
      //用户可以在上传完成之后将数组给清空，这里直接跳转到首页了，没有做清空的操作
    },
    change(file, fileList) {
   //将每次图片数组变化的时候，实时的进行监听，更改数组里面的图片数据
      var arr = [];
      fileList.forEach((item) => {
        arr.push(item.raw);
      });
      this.dataList = arr;
      console.log(arr);
    }
   }
}
</script>