package models

import "github.com/jinzhu/gorm"

// Auth 验证信息
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检查帐号
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}
