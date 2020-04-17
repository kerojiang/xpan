// Package pcscaptcha 验证码处理包
// TODO: 直接打开验证码
package pcscaptcha

import (
	"os"
	"path/filepath"
	"xpan/internal/pcsconfig"
)

const (
	// CaptchaName 验证码文件名称
	CaptchaName = "captcha.png"
)

// RemoveOldCaptchaPath 移除旧的验证码路径
func RemoveOldCaptchaPath() error {
	return os.Remove(filepath.Join(pcsconfig.GetConfigDir(), CaptchaName))
}

// RemoveCaptchaPath 移除验证码路径
func RemoveCaptchaPath() error {
	return os.Remove(CaptchaPath())
}
