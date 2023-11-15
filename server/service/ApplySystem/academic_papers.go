package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AcademicPapersService struct {
}

// CreateAcademicPapers 创建学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) CreateAcademicPapers(AP *ApplySystem.AcademicPapers) (err error) {
	err = global.GVA_DB.Create(AP).Error
	return err
}

// DeleteAcademicPapers 删除学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) DeleteAcademicPapers(AP ApplySystem.AcademicPapers) (err error) {
	err = global.GVA_DB.Delete(&AP).Error
	return err
}

// DeleteAcademicPapersByIds 批量删除学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) DeleteAcademicPapersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ApplySystem.AcademicPapers{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAcademicPapers 更新学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) UpdateAcademicPapers(id uint, maps map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&ApplySystem.AcademicPapers{}).Where("id = ?", id).Updates(maps).Error
	return err
}

// GetAcademicPapers 根据id获取学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) GetAcademicPapers(id uint) (AP ApplySystem.AcademicPapers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&AP).Error
	return
}

// GetAcademicPapersInfoList 分页获取学术论文记录
// Author [piexlmax](https://github.com/piexlmax)
func (APService *AcademicPapersService) GetAcademicPapersInfoList(info ApplySystemReq.AcademicPapersSearch) (list []ApplySystem.AcademicPapers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ApplySystem.AcademicPapers{})
	var APs []ApplySystem.AcademicPapers
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
	if info.Papers_name != "" {
		db = db.Where("papers_name LIKE ?", "%"+info.Papers_name+"%")
	}
	if info.StartPublish_time != nil && info.EndPublish_time != nil {
		db = db.Where("publish_time BETWEEN ? AND ? ", info.StartPublish_time, info.EndPublish_time)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&APs).Error
	return APs, total, err
}
