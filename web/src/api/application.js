import service from '@/utils/request'

// @Tags SysApplication
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApplication true "应用信息"
// @Success 200 {string} string "创建成功"
// @Router /application/createSysApplication [post]
export const createSysApplication = (data) => {
  return service({
    url: '/application/createSysApplication',
    method: 'post',
    data
  })
}

// @Tags SysApplication
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApplication true "应用信息"
// @Success 200 {string} string "删除成功"
// @Router /application/deleteSysApplication [delete]
export const deleteSysApplication = (data) => {
  return service({
    url: '/application/deleteSysApplication',
    method: 'delete',
    data
  })
}

// @Tags SysApplication
// @Summary 批量删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除应用"
// @Success 200 {string} string "批量删除成功"
// @Router /application/deleteSysApplicationByIds [delete]
export const deleteSysApplicationByIds = (data) => {
  return service({
    url: '/application/deleteSysApplicationByIds',
    method: 'delete',
    data
  })
}

// @Tags SysApplication
// @Summary 更新应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApplication true "应用信息"
// @Success 200 {string} string "更新成功"
// @Router /application/updateSysApplication [put]
export const updateSysApplication = (data) => {
  return service({
    url: '/application/updateSysApplication',
    method: 'put',
    data
  })
}

// @Tags SysApplication
// @Summary 用id查询应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysApplication true "用id查询应用"
// @Success 200 {string} string "查询成功"
// @Router /application/findSysApplication [get]
export const findSysApplication = (params) => {
  return service({
    url: '/application/findSysApplication',
    method: 'get',
    params
  })
}

// @Tags SysApplication
// @Summary 分页获取应用列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取应用列表"
// @Success 200 {string} string "获取成功"
// @Router /application/getSysApplicationList [get]
export const getSysApplicationList = (params) => {
  return service({
    url: '/application/getSysApplicationList',
    method: 'get',
    params
  })
}