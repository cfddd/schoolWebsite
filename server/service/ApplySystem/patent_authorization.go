package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	ApplySystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PatentAuthorizationService struct {
}

// CreatePatentAuthorization 创建PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) CreatePatentAuthorization(PA *ApplySystem.PatentAuthorization) (err error) {
	err = global.GVA_DB.Create(PA).Error
	return err
}

// DeletePatentAuthorization 删除PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) DeletePatentAuthorization(PA ApplySystem.PatentAuthorization) (err error) {
	err = global.GVA_DB.Delete(&PA).Error
	return err
}

// DeletePatentAuthorizationByIds 批量删除PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) DeletePatentAuthorizationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ApplySystem.PatentAuthorization{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePatentAuthorization 更新PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) UpdatePatentAuthorization(PA ApplySystem.PatentAuthorization) (err error) {
	err = global.GVA_DB.Save(&PA).Error
	return err
}

// GetPatentAuthorization 根据id获取PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) GetPatentAuthorization(id uint) (PA ApplySystem.PatentAuthorization, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&PA).Error
	return
}

// GetPatentAuthorizationInfoList 分页获取PatentAuthorization记录
// Author [piexlmax](https://github.com/piexlmax)
func (PAService *PatentAuthorizationService) GetPatentAuthorizationInfoList(info ApplySystemReq.PatentAuthorizationSearch) (list []ApplySystem.PatentAuthorization, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ApplySystem.PatentAuthorization{})
	var PAs []ApplySystem.PatentAuthorization
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Student_id != "" {
		db = db.Where("student_id = ?", info.Student_id)
	}
	if info.Student_name != "" {
		db = db.Where("student_name = ?", info.Student_name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PAs).Error
	return PAs, total, err
}
