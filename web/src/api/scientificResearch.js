import service from '@/utils/request'

// @Tags ScientificResearch
// @Summary 创建科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScientificResearch true "创建科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SR/createScientificResearch [post]
export const createScientificResearch = (data) => {
  return service({
    url: '/SR/createScientificResearch',
    method: 'post',
    data
  })
}

// @Tags ScientificResearch
// @Summary 删除科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScientificResearch true "删除科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SR/deleteScientificResearch [delete]
export const deleteScientificResearch = (data) => {
  return service({
    url: '/SR/deleteScientificResearch',
    method: 'delete',
    data
  })
}

// @Tags ScientificResearch
// @Summary 批量删除科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SR/deleteScientificResearch [delete]
export const deleteScientificResearchByIds = (data) => {
  return service({
    url: '/SR/deleteScientificResearchByIds',
    method: 'delete',
    data
  })
}

// @Tags ScientificResearch
// @Summary 更新科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScientificResearch true "更新科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SR/updateScientificResearch [put]
export const updateScientificResearch = (data) => {
  return service({
    url: '/SR/updateScientificResearch',
    method: 'put',
    data
  })
}

// @Tags ScientificResearch
// @Summary 用id查询科研项目填报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ScientificResearch true "用id查询科研项目填报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SR/findScientificResearch [get]
export const findScientificResearch = (params) => {
  return service({
    url: '/SR/findScientificResearch',
    method: 'get',
    params
  })
}

// @Tags ScientificResearch
// @Summary 分页获取科研项目填报列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取科研项目填报列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SR/getScientificResearchList [get]
export const getScientificResearchList = (params) => {
  return service({
    url: '/SR/getScientificResearchList',
    method: 'get',
    params
  })
}
