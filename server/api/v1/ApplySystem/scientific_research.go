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

type ScientificResearchApi struct {
}

var SRService = service.ServiceGroupApp.ApplySystemServiceGroup.ScientificResearchService

// CreateScientificResearch 创建科研项目填报
// @Tags ScientificResearch
// @Summary 创建科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.ScientificResearch true "创建科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SR/createScientificResearch [post]
func (SRApi *ScientificResearchApi) CreateScientificResearch(c *gin.Context) {
	var SR ApplySystem.ScientificResearch
	err := c.ShouldBindJSON(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":              {utils.NotEmpty()},
		"Student_name":            {utils.NotEmpty()},
		"Scientific_project_name": {utils.NotEmpty()},
	}
	if err := utils.Verify(SR, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SRService.CreateScientificResearch(&SR); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteScientificResearch 删除科研项目填报
// @Tags ScientificResearch
// @Summary 删除科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.ScientificResearch true "删除科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SR/deleteScientificResearch [delete]
func (SRApi *ScientificResearchApi) DeleteScientificResearch(c *gin.Context) {
	var SR ApplySystem.ScientificResearch
	err := c.ShouldBindJSON(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SRService.DeleteScientificResearch(SR); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteScientificResearchByIds 批量删除科研项目填报
// @Tags ScientificResearch
// @Summary 批量删除科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /SR/deleteScientificResearchByIds [delete]
func (SRApi *ScientificResearchApi) DeleteScientificResearchByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SRService.DeleteScientificResearchByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateScientificResearch 更新科研项目填报
// @Tags ScientificResearch
// @Summary 更新科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.ScientificResearch true "更新科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SR/updateScientificResearch [put]
func (SRApi *ScientificResearchApi) UpdateScientificResearch(c *gin.Context) {
	var SR ApplySystem.ScientificResearch
	err := c.ShouldBindJSON(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":              {utils.NotEmpty()},
		"Student_name":            {utils.NotEmpty()},
		"Scientific_project_name": {utils.NotEmpty()},
	}
	if err := utils.Verify(SR, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SRService.UpdateScientificResearch(SR); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindScientificResearch 用id查询科研项目填报
// @Tags ScientificResearch
// @Summary 用id查询科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystem.ScientificResearch true "用id查询科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SR/findScientificResearch [get]
func (SRApi *ScientificResearchApi) FindScientificResearch(c *gin.Context) {
	var SR ApplySystem.ScientificResearch
	err := c.ShouldBindQuery(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reSR, err := SRService.GetScientificResearch(SR.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSR": reSR}, c)
	}
}

// GetScientificResearchList 分页获取科研项目填报列表
// @Tags ScientificResearch
// @Summary 分页获取科研项目填报列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystemReq.ScientificResearchSearch true "分页获取科研项目填报列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SR/getScientificResearchList [get]
func (SRApi *ScientificResearchApi) GetScientificResearchList(c *gin.Context) {
	var pageInfo ApplySystemReq.ScientificResearchSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := SRService.GetScientificResearchInfoList(pageInfo); err != nil {
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
