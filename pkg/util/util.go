package util

import "github.com/EDDYCJY/go-gin-example/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
