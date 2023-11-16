package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
)

type Test_uploadService struct {
}

// CreateTest_upload 创建test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService) CreateTest_upload(test_upload *test.Test_upload) (err error) {
	err = global.GVA_DB.Create(test_upload).Error
	return err
}

// DeleteTest_upload 删除test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService)DeleteTest_upload(test_upload test.Test_upload) (err error) {
	err = global.GVA_DB.Delete(&test_upload).Error
	return err
}

// DeleteTest_uploadByIds 批量删除test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService)DeleteTest_uploadByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]test.Test_upload{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTest_upload 更新test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService)UpdateTest_upload(test_upload test.Test_upload) (err error) {
	err = global.GVA_DB.Save(&test_upload).Error
	return err
}

// GetTest_upload 根据id获取test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService)GetTest_upload(id uint) (test_upload test.Test_upload, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&test_upload).Error
	return
}

// GetTest_uploadInfoList 分页获取test_upload记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_uploadService *Test_uploadService)GetTest_uploadInfoList(info testReq.Test_uploadSearch) (list []test.Test_upload, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&test.Test_upload{})
    var test_uploads []test.Test_upload
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&test_uploads).Error
	return  test_uploads, total, err
}
