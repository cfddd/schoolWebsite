import service from '@/utils/request'

// @Tags AcademicPapers
// @Summary 创建学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AcademicPapers true "创建学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AP/createAcademicPapers [post]
export const createAcademicPapers = (data) => {
  return service({
    url: '/AP/createAcademicPapers',
    method: 'post',
    data
  })
}

// @Tags AcademicPapers
// @Summary 删除学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AcademicPapers true "删除学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AP/deleteAcademicPapers [delete]
export const deleteAcademicPapers = (data) => {
  return service({
    url: '/AP/deleteAcademicPapers',
    method: 'delete',
    data
  })
}

// @Tags AcademicPapers
// @Summary 批量删除学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AP/deleteAcademicPapers [delete]
export const deleteAcademicPapersByIds = (data) => {
  return service({
    url: '/AP/deleteAcademicPapersByIds',
    method: 'delete',
    data
  })
}

// @Tags AcademicPapers
// @Summary 更新学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AcademicPapers true "更新学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AP/updateAcademicPapers [put]
export const updateAcademicPapers = (data) => {
  return service({
    url: '/AP/updateAcademicPapers',
    method: 'put',
    data
  })
}

// @Tags AcademicPapers
// @Summary 用id查询学术论文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AcademicPapers true "用id查询学术论文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AP/findAcademicPapers [get]
export const findAcademicPapers = (params) => {
  return service({
    url: '/AP/findAcademicPapers',
    method: 'get',
    params
  })
}

// @Tags AcademicPapers
// @Summary 分页获取学术论文列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取学术论文列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AP/getAcademicPapersList [get]
export const getAcademicPapersList = (params) => {
  return service({
    url: '/AP/getAcademicPapersList',
    method: 'get',
    params
  })
}
