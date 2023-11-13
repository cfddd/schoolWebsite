package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CompetitionPrizeService struct {
}

// CreateCompetitionPrize 创建比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) CreateCompetitionPrize(CP *ApplySystem.CompetitionPrize) (err error) {
	err = global.GVA_DB.Create(CP).Error
	return err
}

// DeleteCompetitionPrize 删除比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) DeleteCompetitionPrize(CP ApplySystem.CompetitionPrize) (err error) {
	err = global.GVA_DB.Delete(&CP).Error
	return err
}

// DeleteCompetitionPrizeByIds 批量删除比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) DeleteCompetitionPrizeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ApplySystem.CompetitionPrize{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCompetitionPrize 更新比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) UpdateCompetitionPrize(id uint, maps map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&ApplySystem.CompetitionPrize{}).Where("id = ?", id).Updates(maps).Error
	return err
}

// GetCompetitionPrize 根据id获取比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) GetCompetitionPrize(id uint) (CP ApplySystem.CompetitionPrize, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&CP).Error
	return
}

// GetCompetitionPrizeInfoList 分页获取比赛获奖申报记录
// Author [piexlmax](https://github.com/piexlmax)
func (CPService *CompetitionPrizeService) GetCompetitionPrizeInfoList(info ApplySystemReq.CompetitionPrizeSearch) (list []ApplySystem.CompetitionPrize, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ApplySystem.CompetitionPrize{})
	var CPs []ApplySystem.CompetitionPrize
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
	if info.StartAward_time != nil && info.EndAward_time != nil {
		db = db.Where("award_time BETWEEN ? AND ? ", info.StartAward_time, info.EndAward_time)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&CPs).Error

	return CPs, total, err
}
