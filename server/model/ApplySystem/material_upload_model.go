package ApplySystem

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

type MaterialUploadModel struct {
	global.GVA_MODEL
	Path               string           `json:"path"`
	CompetitionPrizeID uint             // 外键
	CompetitionPrize   CompetitionPrize // 属于

}

var MUService = service.ServiceGroupApp.ApplySystemServiceGroup.MaterialUploadService

func (mat *MaterialUploadModel) Delete() (err error) {
	// 材料表对应部分删掉
	err = MUService.DeleteMaterialUpload(*mat)
	if err != nil {
		return
	} else {
		// 本地的材料要删掉
		err = os.Remove(mat.Path)
		if err != nil {
			return
		}
	}
	return nil
}

func (mat *MaterialUploadModel) Upload(file *multipart.FileHeader, c *gin.Context) (err error) {
	// 存到本地去
	mat.Path = fmt.Sprintf("server/uploads/file" + file.Filename)
	err = c.SaveUploadedFile(file, mat.Path)
	if err != nil {
		return
	} else {
		// 存到数据库中
		err = MUService.CreateMaterialUpload(mat)
		if err != nil {
			return
		}
	}
	return nil
}
