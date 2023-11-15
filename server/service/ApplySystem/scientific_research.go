package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ScientificResearchService struct {
}

// CreateScientificResearch 创建科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) CreateScientificResearch(SR *ApplySystem.ScientificResearch) (err error) {
	err = global.GVA_DB.Create(SR).Error
	return err
}

// DeleteScientificResearch 删除科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) DeleteScientificResearch(SR ApplySystem.ScientificResearch) (err error) {
	err = global.GVA_DB.Delete(&SR).Error
	return err
}

// DeleteScientificResearchByIds 批量删除科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) DeleteScientificResearchByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ApplySystem.ScientificResearch{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateScientificResearch 更新科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) UpdateScientificResearch(id uint, maps map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&ApplySystem.ScientificResearch{}).Where("id = ?", id).Updates(maps).Error
	return err
}

// GetScientificResearch 根据id获取科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) GetScientificResearch(id uint) (SR ApplySystem.ScientificResearch, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&SR).Error
	return
}

// GetScientificResearchInfoList 分页获取科研项目填报记录
// Author [piexlmax](https://github.com/piexlmax)
func (SRService *ScientificResearchService) GetScientificResearchInfoList(info ApplySystemReq.ScientificResearchSearch) (list []ApplySystem.ScientificResearch, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ApplySystem.ScientificResearch{})
	var SRs []ApplySystem.ScientificResearch
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Student_id != "" {
		db = db.Where("student_id = ?", info.Student_id)
	}
	if info.Student_name != "" {
		db = db.Where("student_name LIKE ?", "%"+info.Student_name+"%")
	}
	if info.Scientific_project_name != "" {
		db = db.Where("scientific_project_name LIKE ?", "%"+info.Scientific_project_name+"%")
	}
	if info.Work_id != "" {
		db = db.Where("work_id = ?", info.Work_id)
	}
	if info.StartProject_approval_time != nil && info.EndProject_approval_time != nil {
		db = db.Where("project_approval_time BETWEEN ? AND ? ", info.StartProject_approval_time, info.EndProject_approval_time)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&SRs).Error
	return SRs, total, err
}
