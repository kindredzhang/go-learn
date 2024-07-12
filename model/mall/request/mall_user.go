package request

type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}
