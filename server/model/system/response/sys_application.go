package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

// SysApplicationResponse 应用响应结构体
type SysApplicationResponse struct {
	Application system.SysApplication `json:"application"`
}

// SysApplicationListResponse 应用列表响应结构体
type SysApplicationListResponse struct {
	List     []system.SysApplication `json:"list"`
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
}