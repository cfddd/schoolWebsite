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
	"time"
)

type PatentAuthorizationApi struct {
}

var PAService = service.ServiceGroupApp.ApplySystemServiceGroup.PatentAuthorizationService

type PatentAuthorizationRequest struct {
	student_id      string     `json:"student_id" binding:"required" msg:"学号必填"`
	student_name    string     `json:"student_name" binding:"required" msg:"学生姓名必填"`
	patent_name     string     `json:"patent_name" binding:"required" msg:"专利名必填"`
	patent_type     string     `json:"patent_type"`
	patent_grant_id string     `json:"patent_grant_id" binding:"required" msg:"专利授权必填"`
	approval_time   *time.Time `json:"approval_time"`
	is_first_invent *bool      `json:"is_first_invent"`
}

// CreatePatentAuthorization 创建PatentAuthorization
// @Tags PatentAuthorization
// @Summary 创建PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.PatentAuthorization true "创建PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PA/createPatentAuthorization [post]
func (PAApi *PatentAuthorizationApi) CreatePatentAuthorization(c *gin.Context) {
	var PA ApplySystem.PatentAuthorization
	err := c.ShouldBindJSON(&PA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":      {utils.NotEmpty()},
		"Student_name":    {utils.NotEmpty()},
		"Patent_name":     {utils.NotEmpty()},
		"Patent_grant_id": {utils.NotEmpty()},
		"Is_first_invent": {utils.NotEmpty()},
	}
	if err := utils.Verify(PA, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PAService.CreatePatentAuthorization(&PA); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePatentAuthorization 删除PatentAuthorization
// @Tags PatentAuthorization
// @Summary 删除PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.PatentAuthorization true "删除PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PA/deletePatentAuthorization [delete]
func (PAApi *PatentAuthorizationApi) DeletePatentAuthorization(c *gin.Context) {
	var PA ApplySystem.PatentAuthorization
	err := c.ShouldBindJSON(&PA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PAService.DeletePatentAuthorization(PA); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePatentAuthorizationByIds 批量删除PatentAuthorization
// @Tags PatentAuthorization
// @Summary 批量删除PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PA/deletePatentAuthorizationByIds [delete]
func (PAApi *PatentAuthorizationApi) DeletePatentAuthorizationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PAService.DeletePatentAuthorizationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePatentAuthorization 更新PatentAuthorization
// @Tags PatentAuthorization
// @Summary 更新PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.PatentAuthorization true "更新PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PA/updatePatentAuthorization [put]
func (PAApi *PatentAuthorizationApi) UpdatePatentAuthorization(c *gin.Context) {
	var PA ApplySystem.PatentAuthorization
	err := c.ShouldBindJSON(&PA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":      {utils.NotEmpty()},
		"Student_name":    {utils.NotEmpty()},
		"Patent_name":     {utils.NotEmpty()},
		"Patent_grant_id": {utils.NotEmpty()},
		"Is_first_invent": {utils.NotEmpty()},
	}
	if err := utils.Verify(PA, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PAService.UpdatePatentAuthorization(PA); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPatentAuthorization 用id查询PatentAuthorization
// @Tags PatentAuthorization
// @Summary 用id查询PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystem.PatentAuthorization true "用id查询PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PA/findPatentAuthorization [get]
func (PAApi *PatentAuthorizationApi) FindPatentAuthorization(c *gin.Context) {
	var PA ApplySystem.PatentAuthorization
	err := c.ShouldBindQuery(&PA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rePA, err := PAService.GetPatentAuthorization(PA.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePA": rePA}, c)
	}
}

// GetPatentAuthorizationList 分页获取PatentAuthorization列表
// @Tags PatentAuthorization
// @Summary 分页获取PatentAuthorization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystemReq.PatentAuthorizationSearch true "分页获取PatentAuthorization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PA/getPatentAuthorizationList [get]
func (PAApi *PatentAuthorizationApi) GetPatentAuthorizationList(c *gin.Context) {
	var pageInfo ApplySystemReq.PatentAuthorizationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := PAService.GetPatentAuthorizationInfoList(pageInfo); err != nil {
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
