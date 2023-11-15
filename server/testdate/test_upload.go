package main

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/res"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"reflect"
	"time"
)

var MUService = service.ServiceGroupApp.ApplySystemServiceGroup.MaterialUploadService
var CPService = service.ServiceGroupApp.ApplySystemServiceGroup.CompetitionPrizeService

type CompetitionPrizeRequest struct {
	Student_id       string     `json:"student_id" binding:"required" msg:"学号必填"`
	Student_name     string     `json:"student_name" binding:"required" msg:"姓名必填"`
	Competition_name string     `json:"competition_name" binding:"required" msg:"比赛名称必填"`
	Award_time       *time.Time `json:"award_time" binding:"required" msg:"获奖时间必填"`
	Award_type       string     `json:"award_type" binding:"required" msg:"获奖类型必填"`
	Award_level      string     `json:"award_level" binding:"required" msg:"获奖等级必填"`
	Competition_type string     `json:"competition_type"`
	Description      string     `json:"description"`
}

func testUpload(c *gin.Context) {
	var ccr CompetitionPrizeRequest
	cover := c.PostForm("cover")
	json.Unmarshal([]byte(cover), &ccr)
	fmt.Println("Cover: ", ccr, "Type", reflect.TypeOf(ccr))
	// 检查当前学号是否正确，是否存在

	// 上传多个图片文件
	form, err := c.MultipartForm()
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["uploads"]
	if !ok {
		global.GVA_LOG.Error("不存在的文件!")
		res.FailWithMessage("不存在的文件", c)
		return
	}
	for _, file := range fileList {
		var material ApplySystem.MaterialUploadModel
		fmt.Println(file.Filename)
		material.Path = fmt.Sprintf("uploads/file/" + file.Filename)
		err = c.SaveUploadedFile(file, material.Path)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}
}

func main() {
	r := gin.Default()

	r.POST("/test_upload", testUpload)

	r.Run("127.0.0.1:8888")
}
