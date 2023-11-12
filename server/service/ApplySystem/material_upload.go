package ApplySystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
)

type MaterialUploadService struct {
}

// CreateMaterialUpload 创建材料
// Author [piexlmax](https://github.com/piexlmax)
func (MUService *MaterialUploadService) CreateMaterialUpload(MU *ApplySystem.MaterialUploadModel) (err error) {
	err = global.GVA_DB.Create(MU).Error
	return err
}

// DeleteMaterialUpload 删除材料
// Author [piexlmax](https://github.com/piexlmax)
func (MUService *MaterialUploadService) DeleteMaterialUpload(MU ApplySystem.MaterialUploadModel) (err error) {
	err = global.GVA_DB.Delete(&MU).Error
	return err
}

// DeleteMaterialUploadByIds 批量删除材料
// Author [piexlmax](https://github.com/piexlmax)
func (MUService *MaterialUploadService) DeleteMaterialUploadByIds(ids []uint) (err error) {
	err = global.GVA_DB.Delete(&[]ApplySystem.MaterialUploadModel{}, "id in ?", ids).Error
	return err
}
