package system

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type ApplicationService struct{}

// CreateSysApplication 创建应用
func (applicationService *ApplicationService) CreateSysApplication(application *system.SysApplication) (err error) {
	// 加密API密钥
	encryptedApiKey, err := applicationService.encryptApiKey(application.ApiKey)
	if err != nil {
		return err
	}
	application.ApiKey = encryptedApiKey

	// 设置默认状态
	if application.Status == 0 {
		application.Status = 1
	}

	err = global.GVA_DB.Create(application).Error
	return err
}

// DeleteSysApplication 删除应用
func (applicationService *ApplicationService) DeleteSysApplication(id string) (err error) {
	err = global.GVA_DB.Delete(&system.SysApplication{}, "id = ?", id).Error
	return err
}

// DeleteSysApplicationByIds 批量删除应用
func (applicationService *ApplicationService) DeleteSysApplicationByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysApplication{}, "id in ?", ids).Error
	return err
}

// UpdateSysApplication 更新应用
func (applicationService *ApplicationService) UpdateSysApplication(sysApplication system.SysApplication) (err error) {
	err = global.GVA_DB.Save(&sysApplication).Error
	return err
}

// GetSysApplication 根据id获取应用记录
func (applicationService *ApplicationService) GetSysApplication(id string) (sysApplication system.SysApplication, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysApplication).Error
	return
}

// GetSysApplicationInfoList 分页获取应用列表
func (applicationService *ApplicationService) GetSysApplicationInfoList(info systemReq.SysApplicationSearch) (list []system.SysApplication, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysApplication{})
	var applications []system.SysApplication
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.BaseUrl != "" {
		db = db.Where("base_url LIKE ?", "%"+info.BaseUrl+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&applications).Error

	// 解密API密钥
	for i := range applications {
		decryptedApiKey, decryptErr := applicationService.decryptApiKey(applications[i].ApiKey)
		if decryptErr != nil {
			// 如果解密失败，显示加密后的部分内容
			applications[i].ApiKey = "****" + applications[i].ApiKey[len(applications[i].ApiKey)-4:]
		} else {
			// 显示部分明文API密钥
			if len(decryptedApiKey) > 8 {
				applications[i].ApiKey = decryptedApiKey[:4] + "****" + decryptedApiKey[len(decryptedApiKey)-4:]
			} else {
				applications[i].ApiKey = "****"
			}
		}
	}

	return applications, total, err
}

// encryptApiKey 加密API密钥
func (applicationService *ApplicationService) encryptApiKey(apiKey string) (string, error) {
	// 使用AES加密
	key := []byte("gin-vue-admin-32-byte-secret-key") // 32字节密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(apiKey), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryptApiKey 解密API密钥
func (applicationService *ApplicationService) decryptApiKey(encryptedApiKey string) (string, error) {
	key := []byte("gin-vue-admin-32-byte-secret-key") // 32字节密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(encryptedApiKey)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}