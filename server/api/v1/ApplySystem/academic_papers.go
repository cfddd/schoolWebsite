package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AcademicPapersApi struct {
}

var APService = service.ServiceGroupApp.ApplySystemServiceGroup.AcademicPapersService

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
	var AP ApplySystem.AcademicPapers
	err := c.ShouldBindJSON(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":       {utils.NotEmpty()},
		"Student_name":     {utils.NotEmpty()},
		"Papers_name":      {utils.NotEmpty()},
		"Publish_at":       {utils.NotEmpty()},
		"Inclusion_status": {utils.NotEmpty()},
	}
	if err := utils.Verify(AP, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := APService.CreateAcademicPapers(&AP); err != nil {
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
	if err := APService.DeleteAcademicPapersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
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
	var AP ApplySystem.AcademicPapers
	err := c.ShouldBindJSON(&AP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":       {utils.NotEmpty()},
		"Student_name":     {utils.NotEmpty()},
		"Papers_name":      {utils.NotEmpty()},
		"Publish_at":       {utils.NotEmpty()},
		"Inclusion_status": {utils.NotEmpty()},
	}
	if err := utils.Verify(AP, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := APService.UpdateAcademicPapers(AP); err != nil {
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
