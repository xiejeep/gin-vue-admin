package request

import (
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SysApplicationSearch 应用搜索结构体
type SysApplicationSearch struct {
	request.PageInfo
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Name           string     `json:"name" form:"name"`
	BaseUrl        string     `json:"baseUrl" form:"baseUrl"`
	Status         *int       `json:"status" form:"status"`
}

// CreateSysApplicationRequest 创建应用请求结构体
type CreateSysApplicationRequest struct {
	Name        string `json:"name" validate:"required" label:"应用名称"`
	ApiKey      string `json:"apiKey" validate:"required" label:"API密钥"`
	BaseUrl     string `json:"baseUrl" validate:"required" label:"基础URL"`
	Description string `json:"description" label:"应用描述"`
}

// UpdateSysApplicationRequest 更新应用请求结构体
type UpdateSysApplicationRequest struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required" label:"应用名称"`
	ApiKey      string `json:"apiKey" validate:"required" label:"API密钥"`
	BaseUrl     string `json:"baseUrl" validate:"required" label:"基础URL"`
	Description string `json:"description" label:"应用描述"`
	Status      int    `json:"status" label:"状态"`
}