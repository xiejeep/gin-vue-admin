package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApplicationRouter struct{}

// InitApplicationRouter 初始化应用管理路由信息
func (s *ApplicationRouter) InitApplicationRouter(Router *gin.RouterGroup) {
	applicationRouter := Router.Group("application").Use(middleware.OperationRecord())
	applicationRouterWithoutRecord := Router.Group("application")
	{
		applicationRouter.POST("createSysApplication", sysApplicationApi.CreateSysApplication)   // 新建应用
		applicationRouter.DELETE("deleteSysApplication", sysApplicationApi.DeleteSysApplication) // 删除应用
		applicationRouter.DELETE("deleteSysApplicationByIds", sysApplicationApi.DeleteSysApplicationByIds) // 批量删除应用
		applicationRouter.PUT("updateSysApplication", sysApplicationApi.UpdateSysApplication)    // 更新应用
	}
	{
		applicationRouterWithoutRecord.GET("findSysApplication", sysApplicationApi.FindSysApplication)        // 根据ID获取应用
		applicationRouterWithoutRecord.GET("getSysApplicationList", sysApplicationApi.GetSysApplicationList) // 获取应用列表
	}
}