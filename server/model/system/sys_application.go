package system

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SysApplication 应用管理结构体
type SysApplication struct {
	ID          string         `gorm:"primarykey;type:varchar(36)" json:"id"`                      // 主键UUID
	CreatedAt   time.Time      `json:"createdAt"`                                                  // 创建时间
	UpdatedAt   time.Time      `json:"updatedAt"`                                                  // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`                                            // 删除时间
	Name        string         `json:"name" gorm:"comment:应用名称;not null" validate:"required"`        // 应用名称
	ApiKey      string         `json:"apiKey" gorm:"comment:API密钥;not null" validate:"required"`     // API密钥(加密存储)
	BaseUrl     string         `json:"baseUrl" gorm:"comment:基础URL;not null" validate:"required"`   // 基础URL
	Description string         `json:"description" gorm:"comment:应用描述"`                           // 应用描述
	Status      int            `json:"status" gorm:"comment:状态;default:1"`                         // 状态 1:启用 0:禁用
}

// BeforeCreate 创建前钩子，自动生成UUID
func (s *SysApplication) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// TableName 设置表名
func (SysApplication) TableName() string {
	return "sys_applications"
}