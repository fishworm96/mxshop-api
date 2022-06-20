package forms

type PassWordListForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号格式有管饭可循 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}