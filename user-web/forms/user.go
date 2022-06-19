package forms

type PassWordListForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"` // 手机号格式有管饭可循 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}