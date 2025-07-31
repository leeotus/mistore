package verify

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

/**
 * @brief 验证码模块
 * @author leeotus (leeotus@163.com)
 * @note 这里采用的是base64Captcha,(https://github.com/mojocn/base64Captcha)
 */

// var store = base64Captcha.DefaultMemStore
var store base64Captcha.Store = RedisStore{}

/**
 * @brief 生成验证码
 */
func GenerateCaptcha() (id string, b64s string, ans string, err error) {
	var driver base64Captcha.Driver

	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuipkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 31,
			G: 98,
			B: 197,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc", "chromohv.ttf"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err = c.Generate()

	return id, b64s, ans, err
}

func VerifyCaptcha(id string, ans string) bool {
	return store.Verify(id, ans, true)
}
