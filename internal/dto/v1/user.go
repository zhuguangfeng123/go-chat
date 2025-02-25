package v1

// UserPwdLoginReq 用户密码登录入参
type UserPwdLoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type SendLoginSmsReq struct {
	Phone string `json:"phone"`
}

type UserSmsLoginReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// UserSignupReq 用户注册入参
type UserSignupReq struct {
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
