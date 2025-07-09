package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysApplicationApi struct{}

var applicationService = service.ServiceGroupApp.SystemServiceGroup.ApplicationService

// CreateSysApplication 创建应用
// @Tags SysApplication
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.CreateSysApplicationRequest true "应用信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /application/createSysApplication [post]
func (sysApplicationApi *SysApplicationApi) CreateSysApplication(c *gin.Context) {
	var req systemReq.CreateSysApplicationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证请求参数
	if err := utils.Verify(req, utils.Rules{
		"Name":    {utils.NotEmpty()},
		"ApiKey":  {utils.NotEmpty()},
		"BaseUrl": {utils.NotEmpty()},
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 转换为数据模型
	application := system.SysApplication{
		Name:        req.Name,
		ApiKey:      req.ApiKey,
		BaseUrl:     req.BaseUrl,
		Description: req.Description,
	}

	if err := applicationService.CreateSysApplication(&application); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysApplication 删除应用
// @Tags SysApplication
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByUUID true "应用ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /application/deleteSysApplication [delete]
func (sysApplicationApi *SysApplicationApi) DeleteSysApplication(c *gin.Context) {
	var reqId request.GetByUUID
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if reqId.ID == "" {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	if err := applicationService.DeleteSysApplication(reqId.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysApplicationByIds 批量删除应用
// @Tags SysApplication
// @Summary 批量删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UUIDsReq true "应用ID列表"
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /application/deleteSysApplicationByIds [delete]
func (sysApplicationApi *SysApplicationApi) DeleteSysApplicationByIds(c *gin.Context) {
	var IDS request.UUIDsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := applicationService.DeleteSysApplicationByIds(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysApplication 更新应用
// @Tags SysApplication
// @Summary 更新应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApplication true "应用信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /application/updateSysApplication [put]
func (sysApplicationApi *SysApplicationApi) UpdateSysApplication(c *gin.Context) {
	var req systemReq.UpdateSysApplicationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(req, utils.SysApplicationVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 构建应用对象
	sysApplication := system.SysApplication{
		ID:          req.ID,
		Name:        req.Name,
		ApiKey:      req.ApiKey,
		BaseUrl:     req.BaseUrl,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := applicationService.UpdateSysApplication(sysApplication); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysApplication 用id查询应用
// @Tags SysApplication
// @Summary 用id查询应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByUUID true "应用ID"
// @Success 200 {object} response.Response{data=systemRes.SysApplicationResponse,msg=string} "查询成功"
// @Router /application/findSysApplication [get]
func (sysApplicationApi *SysApplicationApi) FindSysApplication(c *gin.Context) {
	var reqId request.GetByUUID
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if reqId.ID == "" {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	if resysApplication, err := applicationService.GetSysApplication(reqId.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysApplication": resysApplication}, c)
	}
}

// GetSysApplicationList 分页获取应用列表
// @Tags SysApplication
// @Summary 分页获取应用列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysApplicationSearch true "查询参数"
// @Success 200 {object} response.Response{data=systemRes.SysApplicationListResponse,msg=string} "获取成功"
// @Router /application/getSysApplicationList [get]
func (sysApplicationApi *SysApplicationApi) GetSysApplicationList(c *gin.Context) {
	var pageInfo systemReq.SysApplicationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := applicationService.GetSysApplicationInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysApplicationListResponse{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}