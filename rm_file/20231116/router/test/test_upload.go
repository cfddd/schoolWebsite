package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Test_uploadRouter struct {
}

// InitTest_uploadRouter 初始化 test_upload 路由信息
func (s *Test_uploadRouter) InitTest_uploadRouter(Router *gin.RouterGroup) {
	test_uploadRouter := Router.Group("test_upload").Use(middleware.OperationRecord())
	test_uploadRouterWithoutRecord := Router.Group("test_upload")
	var test_uploadApi = v1.ApiGroupApp.TestApiGroup.Test_uploadApi
	{
		test_uploadRouter.POST("createTest_upload", test_uploadApi.CreateTest_upload)   // 新建test_upload
		test_uploadRouter.DELETE("deleteTest_upload", test_uploadApi.DeleteTest_upload) // 删除test_upload
		test_uploadRouter.DELETE("deleteTest_uploadByIds", test_uploadApi.DeleteTest_uploadByIds) // 批量删除test_upload
		test_uploadRouter.PUT("updateTest_upload", test_uploadApi.UpdateTest_upload)    // 更新test_upload
	}
	{
		test_uploadRouterWithoutRecord.GET("findTest_upload", test_uploadApi.FindTest_upload)        // 根据ID获取test_upload
		test_uploadRouterWithoutRecord.GET("getTest_uploadList", test_uploadApi.GetTest_uploadList)  // 获取test_upload列表
	}
}
