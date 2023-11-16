import service from '@/utils/request'

// @Tags Test_upload
// @Summary 创建test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_upload true "创建test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /test_upload/createTest_upload [post]
export const createTest_upload = (data) => {
  return service({
    url: '/test_upload/createTest_upload',
    method: 'post',
    data
  })
}

// @Tags Test_upload
// @Summary 删除test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_upload true "删除test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_upload/deleteTest_upload [delete]
export const deleteTest_upload = (data) => {
  return service({
    url: '/test_upload/deleteTest_upload',
    method: 'delete',
    data
  })
}

// @Tags Test_upload
// @Summary 批量删除test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_upload/deleteTest_upload [delete]
export const deleteTest_uploadByIds = (data) => {
  return service({
    url: '/test_upload/deleteTest_uploadByIds',
    method: 'delete',
    data
  })
}

// @Tags Test_upload
// @Summary 更新test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_upload true "更新test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_upload/updateTest_upload [put]
export const updateTest_upload = (data) => {
  return service({
    url: '/test_upload/updateTest_upload',
    method: 'put',
    data
  })
}

// @Tags Test_upload
// @Summary 用id查询test_upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Test_upload true "用id查询test_upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_upload/findTest_upload [get]
export const findTest_upload = (params) => {
  return service({
    url: '/test_upload/findTest_upload',
    method: 'get',
    params
  })
}

// @Tags Test_upload
// @Summary 分页获取test_upload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取test_upload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_upload/getTest_uploadList [get]
export const getTest_uploadList = (params) => {
  return service({
    url: '/test_upload/getTest_uploadList',
    method: 'get',
    params
  })
}
