package ApplySystem

import (
	"github.com/fatih/structs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type AcademicPapersApi struct {
}

var APService = service.ServiceGroupApp.ApplySystemServiceGroup.AcademicPapersService

type AcademicPapersRequest struct {
	Student_id       string     `json:"student_id" binding:"required" msg:"学号必填"`
	Student_name     string     `json:"student_name" binding:"required" msg:"姓名必填"`
	Papers_name      string     `json:"papers_name" binding:"required" msg:"论文名称必填"`
	Publish_at       string     `json:"publish_at" binding:"required" msg:"发表期刊必填"`
	Publish_time     *time.Time `json:"publish_time"`
	Inclusion_status string     `json:"inclusion_status" binding:"required" msg:"收录情况必填"`
}

// CreateAcademicPapers 创建学术论文
// @Tags AcademicPapers
// @Summary 创建学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.AcademicPapers true "创建学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AP/createAcademicPapers [post]
func (APApi *AcademicPapersApi) CreateAcademicPapers(c *gin.Context) {
	var AP AcademicPapersRequest
	err := c.ShouldBindJSON(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 证明材料
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("资料上传失败", c)
		global.GVA_LOG.Error("资料上传失败", zap.Error(err))
		return
	}

	fileList, ok := form.File["uploads"]
	if !ok {
		response.FailWithMessage("不存在的文件", c)
		global.GVA_LOG.Error("不存在的文件")
		return
	}

	materialList := []ApplySystem.MaterialUploadModel{}
	for _, file := range fileList {
		var material ApplySystem.MaterialUploadModel
		MaterialUpload(&material, file, c)
		materialList = append(materialList, material)
	}

	err = APService.CreateAcademicPapers(&ApplySystem.AcademicPapers{
		Student_id:           AP.Student_id,
		Student_name:         AP.Student_name,
		Papers_name:          AP.Papers_name,
		Publish_at:           AP.Publish_at,
		Publish_time:         AP.Publish_time,
		Inclusion_status:     AP.Inclusion_status,
		Audit_status:         0,
		Hint_message:         "",
		MaterialUploadModels: materialList,
	})

	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAcademicPapers 删除学术论文
// @Tags AcademicPapers
// @Summary 删除学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.AcademicPapers true "删除学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AP/deleteAcademicPapers [delete]
func (APApi *AcademicPapersApi) DeleteAcademicPapers(c *gin.Context) {
	var AP ApplySystem.AcademicPapers
	err := c.ShouldBindJSON(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 删除对应的材料
	for _, material := range AP.MaterialUploadModels {
		err = MaterialDelete(&material)
		if err != nil {
			global.GVA_LOG.Error("资料删除失败!", zap.Error(err))
			response.FailWithMessage("资料删除失败", c)
			return
		}
	}

	if err := APService.DeleteAcademicPapers(AP); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAcademicPapersByIds 批量删除学术论文
// @Tags AcademicPapers
// @Summary 批量删除学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /AP/deleteAcademicPapersByIds [delete]
func (APApi *AcademicPapersApi) DeleteAcademicPapersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 跟举IDS查找对应要删除的申请
	var APS []ApplySystem.AcademicPapers
	for _, ID := range IDS.Ids {
		AP, err := APService.GetAcademicPapers(ID)
		if err != nil {
			global.GVA_LOG.Error("没有这条数据", zap.Error(err))
			response.FailWithMessage("没有这条数据", c)
			return
		}
		APS = append(APS, AP)
	}
	// 删除材料
	for _, AP := range APS {
		for _, material := range AP.MaterialUploadModels {
			err = MaterialDelete(&material)
			if err != nil {
				global.GVA_LOG.Error("资料删除失败!", zap.Error(err))
				response.FailWithMessage("资料删除失败", c)
				return
			}
		}
	}

	if err := APService.DeleteAcademicPapersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

type AcademicPapersUpdateRequest struct {
	Student_id       string     `json:"student_id"`
	Student_name     string     `json:"student_name"`
	Papers_name      string     `json:"papers_name"`
	Publish_at       string     `json:"publish_at"`
	Publish_time     *time.Time `json:"publish_time"`
	Inclusion_status string     `json:"inclusion_status"`
	ID               uint       `json:"id"`
}

// UpdateAcademicPapers 更新学术论文
// @Tags AcademicPapers
// @Summary 更新学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.AcademicPapers true "更新学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AP/updateAcademicPapers [put]
func (APApi *AcademicPapersApi) UpdateAcademicPapers(c *gin.Context) {
	var AP AcademicPapersUpdateRequest
	err := c.ShouldBindJSON(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 资料更新
	form, err := c.MultipartForm()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	fileList, ok := form.File["uploads"]
	if !ok {
		global.GVA_LOG.Error("不存在文件", zap.Error(err))
		response.FailWithMessage("不存在文件", c)
		return
	}

	// 查找旧的申请
	AcademicPapersOld, err := APService.GetAcademicPapers(AP.ID)
	if err != nil {
		global.GVA_LOG.Error("申请不存在", zap.Error(err))
		response.FailWithMessage("申请不存在", c)
		return
	}

	// 查看当前申请的状态是否可以修改
	if AcademicPapersOld.Audit_status == 0 || AcademicPapersOld.Audit_status == 1 {
		global.GVA_LOG.Error("申请不可修改", zap.Error(err))
		response.FailWithMessage("申请不可修改", c)
		return
	}

	// 判断有没有传新的证明材料，如果有就要把旧的删掉
	materialList := []ApplySystem.MaterialUploadModel{} // 材料
	if len(fileList) > 0 {
		// 材料表对应部分删掉
		for _, material := range AcademicPapersOld.MaterialUploadModels {
			err = MUService.DeleteMaterialUpload(material)
		}
		// 将新的材料部分加入材料表数据库
		for _, file := range fileList {
			var material ApplySystem.MaterialUploadModel
			err = MaterialUpload(&material, file, c)
			materialList = append(materialList, material)
		}
	}

	// 将要更新的参数赋值成数据库存储的结构体的类型
	competitionPrizeNew := ApplySystem.AcademicPapers{
		UpdatedAt:            time.Now().Format("2006-01-02 15:04:05"),
		Student_id:           AP.Student_id,
		Student_name:         AP.Student_name,
		Papers_name:          AP.Papers_name,
		Publish_at:           AP.Publish_at,
		Publish_time:         AP.Publish_time,
		Inclusion_status:     AP.Inclusion_status,
		MaterialUploadModels: materialList, // 资料也要被替换
	}

	// 转化成map类型
	maps := structs.Map(&competitionPrizeNew)
	var dataMap = make(map[string]interface{})
	for k, v := range maps {
		switch t := v.(type) {
		case string:
			if t == "" {
				continue
			}
		case *time.Time:
			if t == nil {
				continue
			}
		case []ApplySystem.MaterialUploadModel:
			if len(t) == 0 {
				continue
			}
		}
		dataMap[k] = v
	}

	if err := APService.UpdateAcademicPapers(AP.ID, dataMap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAcademicPapers 用id查询学术论文
// @Tags AcademicPapers
// @Summary 用id查询学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystem.AcademicPapers true "用id查询学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AP/findAcademicPapers [get]
func (APApi *AcademicPapersApi) FindAcademicPapers(c *gin.Context) {
	var AP ApplySystem.AcademicPapers
	err := c.ShouldBindQuery(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reAP, err := APService.GetAcademicPapers(AP.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reAP": reAP}, c)
	}
}

// GetAcademicPapersList 分页获取学术论文列表
// @Tags AcademicPapers
// @Summary 分页获取学术论文列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystemReq.AcademicPapersSearch true "分页获取学术论文列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AP/getAcademicPapersList [get]
func (APApi *AcademicPapersApi) GetAcademicPapersList(c *gin.Context) {
	var pageInfo ApplySystemReq.AcademicPapersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := APService.GetAcademicPapersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
