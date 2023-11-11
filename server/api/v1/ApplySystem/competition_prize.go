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

type CompetitionPrizeApi struct {
}

var CPService = service.ServiceGroupApp.ApplySystemServiceGroup.CompetitionPrizeService

// CreateCompetitionPrize 创建比赛获奖申报
// @Tags CompetitionPrize
// @Summary 创建比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.CompetitionPrize true "创建比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /CP/createCompetitionPrize [post]
func (CPApi *CompetitionPrizeApi) CreateCompetitionPrize(c *gin.Context) {
	var CP ApplySystem.CompetitionPrize
	err := c.ShouldBindJSON(&CP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":       {utils.NotEmpty()},
		"Student_name":     {utils.NotEmpty()},
		"Competition_name": {utils.NotEmpty()},
		"Award_type":       {utils.NotEmpty()},
		"Award_level":      {utils.NotEmpty()},
	}
	if err := utils.Verify(CP, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := CPService.CreateCompetitionPrize(&CP); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompetitionPrize 删除比赛获奖申报
// @Tags CompetitionPrize
// @Summary 删除比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.CompetitionPrize true "删除比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CP/deleteCompetitionPrize [delete]
func (CPApi *CompetitionPrizeApi) DeleteCompetitionPrize(c *gin.Context) {
	var CP ApplySystem.CompetitionPrize
	err := c.ShouldBindJSON(&CP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := CPService.DeleteCompetitionPrize(CP); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompetitionPrizeByIds 批量删除比赛获奖申报
// @Tags CompetitionPrize
// @Summary 批量删除比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /CP/deleteCompetitionPrizeByIds [delete]
func (CPApi *CompetitionPrizeApi) DeleteCompetitionPrizeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := CPService.DeleteCompetitionPrizeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompetitionPrize 更新比赛获奖申报
// @Tags CompetitionPrize
// @Summary 更新比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.CompetitionPrize true "更新比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CP/updateCompetitionPrize [put]
func (CPApi *CompetitionPrizeApi) UpdateCompetitionPrize(c *gin.Context) {
	var CP ApplySystem.CompetitionPrize
	err := c.ShouldBindJSON(&CP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Student_id":       {utils.NotEmpty()},
		"Student_name":     {utils.NotEmpty()},
		"Competition_name": {utils.NotEmpty()},
		"Award_type":       {utils.NotEmpty()},
		"Award_level":      {utils.NotEmpty()},
	}
	if err := utils.Verify(CP, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := CPService.UpdateCompetitionPrize(CP); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type QueryRequest struct {
	Id uint `json:"id"`
}

// FindCompetitionPrize 用id查询比赛获奖申报
// @Tags CompetitionPrize
// @Summary 用id查询比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystem.CompetitionPrize true "用id查询比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CP/findCompetitionPrize [get]
func (CPApi *CompetitionPrizeApi) FindCompetitionPrize(c *gin.Context) {
	var cr QueryRequest
	err := c.ShouldBindQuery(&cr)
	//var CP ApplySystem.CompetitionPrize
	//err := c.ShouldBindQuery(&CP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reCP, err := CPService.GetCompetitionPrize(cr.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCP": reCP}, c)
	}
}

// GetCompetitionPrizeList 分页获取比赛获奖申报列表
// @Tags CompetitionPrize
// @Summary 分页获取比赛获奖申报列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ApplySystemReq.CompetitionPrizeSearch true "分页获取比赛获奖申报列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CP/getCompetitionPrizeList [get]
func (CPApi *CompetitionPrizeApi) GetCompetitionPrizeList(c *gin.Context) {
	var pageInfo ApplySystemReq.CompetitionPrizeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := CPService.GetCompetitionPrizeInfoList(pageInfo); err != nil {
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
