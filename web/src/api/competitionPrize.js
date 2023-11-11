import service from '@/utils/request'

// @Tags CompetitionPrize
// @Summary 创建比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionPrize true "创建比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /CP/createCompetitionPrize [post]
export const createCompetitionPrize = (data) => {
  return service({
    url: '/CP/createCompetitionPrize',
    method: 'post',
    data
  })
}

// @Tags CompetitionPrize
// @Summary 删除比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionPrize true "删除比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CP/deleteCompetitionPrize [delete]
export const deleteCompetitionPrize = (data) => {
  return service({
    url: '/CP/deleteCompetitionPrize',
    method: 'delete',
    data
  })
}

// @Tags CompetitionPrize
// @Summary 批量删除比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CP/deleteCompetitionPrize [delete]
export const deleteCompetitionPrizeByIds = (data) => {
  return service({
    url: '/CP/deleteCompetitionPrizeByIds',
    method: 'delete',
    data
  })
}

// @Tags CompetitionPrize
// @Summary 更新比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionPrize true "更新比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CP/updateCompetitionPrize [put]
export const updateCompetitionPrize = (data) => {
  return service({
    url: '/CP/updateCompetitionPrize',
    method: 'put',
    data
  })
}

// @Tags CompetitionPrize
// @Summary 用id查询比赛获奖申报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompetitionPrize true "用id查询比赛获奖申报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CP/findCompetitionPrize [get]
export const findCompetitionPrize = (params) => {
  return service({
    url: '/CP/findCompetitionPrize',
    method: 'get',
    params
  })
}

// @Tags CompetitionPrize
// @Summary 分页获取比赛获奖申报列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取比赛获奖申报列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CP/getCompetitionPrizeList [get]
export const getCompetitionPrizeList = (params) => {
  return service({
    url: '/CP/getCompetitionPrizeList',
    method: 'get',
    params
  })
}
