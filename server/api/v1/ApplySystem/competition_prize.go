package ApplySystem

import (
	"github.com/fatih/structs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/res"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type CompetitionPrizeApi struct {
}

var CPService = service.ServiceGroupApp.ApplySystemServiceGroup.CompetitionPrizeService
var MUService = service.ServiceGroupApp.ApplySystemServiceGroup.MaterialUploadService

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
	// 检查当前学号是否正确，是否存在

	// 上传多个图片文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["uploads"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	materialList := []ApplySystem.MaterialUploadModel{}
	for _, file := range fileList {
		var material ApplySystem.MaterialUploadModel
		err = MaterialUpload(&material, file, c)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		materialList = append(materialList, material)
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
	// 删除对应的材料
	for _, material := range CP.MaterialUploadModels {
		err = MaterialDelete(&material)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
			return
		}
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
	// 删除材料，

	if err := CPService.DeleteCompetitionPrizeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

type CompetitionPrizeUpdateRequest struct {
	Student_id       string     `json:"student_id"`
	Student_name     string     `json:"student_name"`
	Competition_name string     `json:"competition_name"`
	Award_time       *time.Time `json:"award_time"`
	Award_type       string     `json:"award_type"`
	Award_level      string     `json:"award_level"`
	Competition_type string     `json:"competition_type"`
	Description      string     `json:"description"`
	ID               uint       `json:"id"`
}

// UpdateCompetitionPrizeStudent 更新比赛获奖申报
// @Tags CompetitionPrize
// @Summary 更新比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ApplySystem.CompetitionPrize true "更新比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CP/updateCompetitionPrize [put]
func (CPApi *CompetitionPrizeApi) UpdateCompetitionPrizeStudent(c *gin.Context) {
	// 要更新的参数
	var cr CompetitionPrizeUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.GVA_LOG.Error("读入参数失败", zap.Error(err))
		res.FailWithError(err, cr, c)
		return
	}
	// 上传了新的资料
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["uploads"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	// 根据当前该条的申请表url中的id查找到原来的申请证明材料
	var CompentitionPrizeOld ApplySystem.CompetitionPrize
	CompentitionPrizeOld, err = CPService.GetCompetitionPrize(cr.ID)
	if err != nil {
		global.GVA_LOG.Error("文章不存在", zap.Error(err))
		response.FailWithMessage("文章不存在", c)
		return
	}

	// 如果当前的状态是已通过1或者已提交0那就都不能修改，只有是未通过才能修改
	if CompentitionPrizeOld.Audit_status == 0 || CompentitionPrizeOld.Audit_status == 1 {
		global.GVA_LOG.Error("不能修改", zap.Error(err))
		response.FailWithMessage("不能修改", c)
		return
	}

	// 判断有没有传新的证明材料，如果有就要把旧的删掉
	materialList := []ApplySystem.MaterialUploadModel{} // 材料
	if len(fileList) > 0 {
		// 材料表对应部分删掉
		for _, material := range CompentitionPrizeOld.MaterialUploadModels {
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
	competitionPrizeNew := ApplySystem.CompetitionPrize{
		UpdatedAt:            time.Now().Format("2006-01-02 15:04:05"),
		Student_id:           cr.Student_id,
		Student_name:         cr.Student_name,
		Competition_name:     cr.Competition_name,
		Award_time:           cr.Award_time,
		Award_type:           cr.Award_type,
		Award_level:          cr.Award_level,
		Competition_type:     cr.Competition_type,
		Description:          cr.Description,
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

	if err := CPService.UpdateCompetitionPrize(cr.ID, dataMap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
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
	// 查询具体到每一条的我的申报，那么我要显示的内容，
	var cr ApplySystemReq.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWitheCode(res.ArgumentError, c)
		return
	}
	if reCP, err := CPService.GetCompetitionPrize(cr.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reCP, c)
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

	//id判断

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
