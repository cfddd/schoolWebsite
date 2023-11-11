package ApplySystem

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/res"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type CompetitionPrizeApi struct {
}

var CPService = service.ServiceGroupApp.ApplySystemServiceGroup.CompetitionPrizeService

type CompetitionPrizeRequest struct {
	Student_id       string     `json:"student_Id" binding:"required" msg:"学号必填"`
	Student_name     string     `json:"student_name" binding:"required" msg:"姓名必填"`
	Competition_name string     `json:"competition_name" binding:"required" msg:"比赛名称必填"`
	Award_time       *time.Time `json:"award_time" binding:"required" msg:"获奖时间必填"`
	Award_type       string     `json:"award_type" binding:"required" msg:"获奖类型必填"`
	Award_level      string     `json:"award_level" binding:"required" msg:"获奖等级必填"`
	Competition_type string     `json:"competition_type"`
	Description      string     `json:"description"`
}

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
	var cr CompetitionPrizeRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, cr, c)
		return
	}

	// 上传多个图片文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["uploads"]
	materialList := []ApplySystem.MaterialUploadModel{}
	for _, file := range fileList {
		var material ApplySystem.MaterialUploadModel
		material.Path = fmt.Sprintf("server/uploads/file" + file.Filename)
		// 存到本地去
		c.SaveUploadedFile(file, material.Path)

		materialList = append(materialList, material)
	}
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	err = CPService.CreateCompetitionPrize(&ApplySystem.CompetitionPrize{
		Student_id:           cr.Student_id,
		Student_name:         cr.Student_name,
		Competition_name:     cr.Competition_name,
		Award_time:           cr.Award_time,
		Award_type:           cr.Award_type,
		Award_level:          cr.Award_level,
		Competition_type:     cr.Competition_type,
		Description:          cr.Description,
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

//type QueryRequest struct {
//	Id uint `json:"id"`
//}

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
	//var cr QueryRequest
	//err := c.ShouldBindQuery(&cr)
	////var CP ApplySystem.CompetitionPrize
	////err := c.ShouldBindQuery(&CP)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	userID := utils.GetUserID(c)
	if reCP, err := CPService.GetCompetitionPrize(userID); err != nil {
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

// ResultDetailView 提交后查询
func ResultDetailView(c *gin.Context) {
	userID := utils.GetUserID(c)
	competitionPrize, err := CPService.GetCompetitionPrize(userID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithData(competitionPrize, c)
}
