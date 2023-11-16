// 自动生成模板Test_upload
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// test_upload 结构体  Test_upload
type Test_upload struct {
      global.GVA_MODEL
      Cfd  string `json:"cfd" form:"cfd" gorm:"column:cfd;comment:;"`  //ok 
      Upload  string `json:"upload" form:"upload" gorm:"column:upload;comment:;"`  //上传 
}


// TableName test_upload Test_upload自定义表名 test_upload
func (Test_upload) TableName() string {
  return "test_upload"
}

