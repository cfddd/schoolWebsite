package ApplySystem

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ApplySystem"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

type ApiGroup struct {
	CompetitionPrizeApi
	ScientificResearchApi
	AcademicPapersApi
	PatentAuthorizationApi
}

var MUService = service.ServiceGroupApp.ApplySystemServiceGroup.MaterialUploadService

// MaterialDelete 删除资料
func MaterialDelete(mat *ApplySystem.MaterialUploadModel) (err error) {
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

func MaterialUpload(mat *ApplySystem.MaterialUploadModel, file *multipart.FileHeader, c *gin.Context) (err error) {
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
