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

type ScientificResearchApi struct {
}

var SRService = service.ServiceGroupApp.ApplySystemServiceGroup.ScientificResearchService

type ScientificResearchRequest struct {
	Student_id              string     `json:"student_id" binding:"required" msg:"学号必填"`                //学号
	Student_name            string     `json:"student_name" binding:"required" msg:"学生姓名必填"`            //学生姓名
	Scientific_project_name string     `json:"scientific_project_name" binding:"required" msg:"科研项目必填"` //科研项目名称
	Work_id                 string     `json:"work_id"`                                                 //工号
	Project_approval_time   *time.Time `json:"project_approval_time"`                                   //项目获批时间
}

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
	var SR ScientificResearchRequest
	err := c.ShouldBindJSON(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//上传证明材料
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
		err = MaterialUpload(&material, file, c) //将材料信息入库
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		materialList = append(materialList, material)
	}

	err = SRService.CreateScientificResearch(&ApplySystem.ScientificResearch{
		Student_id:              SR.Student_id,
		Student_name:            SR.Student_name,
		Scientific_project_name: SR.Scientific_project_name,
		Work_id:                 SR.Work_id,
		Project_approval_time:   SR.Project_approval_time,
		MaterialUploadModels:    materialList,
		Audit_status:            0,
		Hint_message:            "",
	})

	if err != nil {
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
	//删除证明材料
	for _, material := range SR.MaterialUploadModels {
		err = MaterialDelete(&material)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
			return
		}
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

	// 根据IDS查找对应要删除的申请
	var SRS []ApplySystem.ScientificResearch
	for _, ID := range IDS.Ids {
		SR, err := SRService.GetScientificResearch(ID)
		if err != nil {
			global.GVA_LOG.Error("没有这条数据", zap.Error(err))
			response.FailWithMessage("没有这条数据", c)
			return
		}
		SRS = append(SRS, SR)
	}

	// 删除材料
	for _, SR := range SRS {
		for _, material := range SR.MaterialUploadModels {
			err = MaterialDelete(&material)
			if err != nil {
				global.GVA_LOG.Error("删除失败!", zap.Error(err))
				response.FailWithMessage("删除失败", c)
				return
			}
		}
	}

	if err := SRService.DeleteScientificResearchByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

type ScientificResearchUpdateRequest struct {
	Student_id              string     `json:"student_id"`
	Student_name            string     `json:"student_name"`
	Scientific_project_name string     `json:"scientific_project_name"`
	Work_id                 string     `json:"work_id" form:"work_id"`
	Project_approval_time   *time.Time `json:"project_approval_time"`
	ID                      uint       `json:"id"`
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
	//获取要更新的参数
	var SR ScientificResearchUpdateRequest
	err := c.ShouldBindJSON(&SR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 如果当前的状态是已通过1或者已提交0那就都不能修改，只有是未通过才能修改
	var ScientificResearchOld ApplySystem.ScientificResearch
	ScientificResearchOld, err = SRService.GetScientificResearch(SR.ID)
	if ScientificResearchOld.Audit_status == 0 || ScientificResearchOld.Audit_status == 1 {
		global.GVA_LOG.Error("不能修改", zap.Error(err))
		response.FailWithMessage("不能修改", c)
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

	// 判断有没有传新的证明材料，如果有就要把旧的删掉
	materialList := []ApplySystem.MaterialUploadModel{} // 材料
	if len(fileList) > 0 {
		// 材料表对应部分删掉
		for _, material := range ScientificResearchOld.MaterialUploadModels {
			err = MUService.DeleteMaterialUpload(material)
		}
		// 将新的材料部分加入材料表数据库
		for _, file := range fileList {
			var material ApplySystem.MaterialUploadModel
			err = MaterialUpload(&material, file, c)
			materialList = append(materialList, material)
		}
	}

	ScientificResearchNew := ApplySystem.ScientificResearch{
		UpdatedAt:               time.Now().Format("2006-01-02 15:04:05"),
		Student_id:              SR.Student_id,
		Student_name:            SR.Student_name,
		Scientific_project_name: SR.Scientific_project_name,
		Work_id:                 SR.Work_id,
		Project_approval_time:   SR.Project_approval_time,
		MaterialUploadModels:    materialList,
	}

	// 转化成map类型
	maps := structs.Map(&ScientificResearchNew)
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

	if err := SRService.UpdateScientificResearch(SR.ID, dataMap); err != nil {
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
	var SR ApplySystemReq.IDRequest
	err := c.ShouldBindJSON(&SR)
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
