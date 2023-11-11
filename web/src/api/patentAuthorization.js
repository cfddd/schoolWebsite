import service from '@/utils/request'

// @Tags PatentAuthorization
// @Summary 创建PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PatentAuthorization true "创建PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PA/createPatentAuthorization [post]
export const createPatentAuthorization = (data) => {
  return service({
    url: '/PA/createPatentAuthorization',
    method: 'post',
    data
  })
}

// @Tags PatentAuthorization
// @Summary 删除PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PatentAuthorization true "删除PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PA/deletePatentAuthorization [delete]
export const deletePatentAuthorization = (data) => {
  return service({
    url: '/PA/deletePatentAuthorization',
    method: 'delete',
    data
  })
}

// @Tags PatentAuthorization
// @Summary 批量删除PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PA/deletePatentAuthorization [delete]
export const deletePatentAuthorizationByIds = (data) => {
  return service({
    url: '/PA/deletePatentAuthorizationByIds',
    method: 'delete',
    data
  })
}

// @Tags PatentAuthorization
// @Summary 更新PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PatentAuthorization true "更新PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PA/updatePatentAuthorization [put]
export const updatePatentAuthorization = (data) => {
  return service({
    url: '/PA/updatePatentAuthorization',
    method: 'put',
    data
  })
}

// @Tags PatentAuthorization
// @Summary 用id查询PatentAuthorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PatentAuthorization true "用id查询PatentAuthorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PA/findPatentAuthorization [get]
export const findPatentAuthorization = (params) => {
  return service({
    url: '/PA/findPatentAuthorization',
    method: 'get',
    params
  })
}

// @Tags PatentAuthorization
// @Summary 分页获取PatentAuthorization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PatentAuthorization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PA/getPatentAuthorizationList [get]
export const getPatentAuthorizationList = (params) => {
  return service({
    url: '/PA/getPatentAuthorizationList',
    method: 'get',
    params
  })
}
