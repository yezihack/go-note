package main

// LoginResponse 登录接口的响应参数
type LoginResponse struct {
	Errno string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data User `json:"data"`
}
// 承接前端传过来的json数据或form表单数据
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Age int `form:"age" json:"age" binding:"required"`
}