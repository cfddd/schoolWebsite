package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/test"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Test_uploadApi struct {
}

var test_uploadService = service.ServiceGroupApp.TestServiceGroup.Test_uploadService


// CreateTest_upload 创建test_upload
// @Tags Test_upload
// @Summary 创建test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test_upload true "创建test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /test_upload/createTest_upload [post]
func (test_uploadApi *Test_uploadApi) CreateTest_upload(c *gin.Context) {
	var test_upload test.Test_upload
	err := c.ShouldBindJSON(&test_upload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Cfd":{utils.NotEmpty()},
        "Upload":{utils.NotEmpty()},
    }
	if err := utils.Verify(test_upload, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := test_uploadService.CreateTest_upload(&test_upload); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest_upload 删除test_upload
// @Tags Test_upload
// @Summary 删除test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test_upload true "删除test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_upload/deleteTest_upload [delete]
func (test_uploadApi *Test_uploadApi) DeleteTest_upload(c *gin.Context) {
	var test_upload test.Test_upload
	err := c.ShouldBindJSON(&test_upload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := test_uploadService.DeleteTest_upload(test_upload); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTest_uploadByIds 批量删除test_upload
// @Tags Test_upload
// @Summary 批量删除test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /test_upload/deleteTest_uploadByIds [delete]
func (test_uploadApi *Test_uploadApi) DeleteTest_uploadByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := test_uploadService.DeleteTest_uploadByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest_upload 更新test_upload
// @Tags Test_upload
// @Summary 更新test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test_upload true "更新test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_upload/updateTest_upload [put]
func (test_uploadApi *Test_uploadApi) UpdateTest_upload(c *gin.Context) {
	var test_upload test.Test_upload
	err := c.ShouldBindJSON(&test_upload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Cfd":{utils.NotEmpty()},
          "Upload":{utils.NotEmpty()},
      }
    if err := utils.Verify(test_upload, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := test_uploadService.UpdateTest_upload(test_upload); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest_upload 用id查询test_upload
// @Tags Test_upload
// @Summary 用id查询test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.Test_upload true "用id查询test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_upload/findTest_upload [get]
func (test_uploadApi *Test_uploadApi) FindTest_upload(c *gin.Context) {
	var test_upload test.Test_upload
	err := c.ShouldBindQuery(&test_upload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retest_upload, err := test_uploadService.GetTest_upload(test_upload.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retest_upload": retest_upload}, c)
	}
}

// GetTest_uploadList 分页获取test_upload列表
// @Tags Test_upload
// @Summary 分页获取test_upload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.Test_uploadSearch true "分页获取test_upload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_upload/getTest_uploadList [get]
func (test_uploadApi *Test_uploadApi) GetTest_uploadList(c *gin.Context) {
	var pageInfo testReq.Test_uploadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := test_uploadService.GetTest_uploadInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
