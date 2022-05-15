package ErrMsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 10** 用户模块错误
	ERROR_PASSWORD_WRONG      = 1001
	ERROR_USERNAME_EXIST      = 1002
	ERROR_TOKEN_WRONG         = 1011
	ERROR_TOKEN_RUNTIME       = 1012
	ERROR_USERLIST_WRONG      = 1021
	ERROR_USEREDIT_WRONG      = 1031
	ERROR_USERDELETE_NOTFOUND = 1041
	ERROR_USERDELETE_WRONG    = 1042
	// code = 20** jwt模块相关
	ERROR_JWT_MISSING     = 2001
	ERROR_JWT_TYPEWRONG   = 2002
	ERROR_JWT_TIMEEXPIRED = 2003
	ERROR_JWT_NOTVALID    = 2004

	// code = 200*** 文章模块错误
)

var codemsg = map[int]string{
	SUCCESS:                   "OK",
	ERROR:                     "FAIL",
	ERROR_PASSWORD_WRONG:      "用户名或密码错误",
	ERROR_USERNAME_EXIST:      "用户已存在",
	ERROR_TOKEN_WRONG:         "用户TOKEN错误",
	ERROR_TOKEN_RUNTIME:       "用户TOKEN过期",
	ERROR_USERLIST_WRONG:      "获取用户列表失败",
	ERROR_USEREDIT_WRONG:      "用户更新失败",
	ERROR_USERDELETE_NOTFOUND: "用户不存在",
	ERROR_USERDELETE_WRONG:    "用户删除失败",
	ERROR_JWT_MISSING:         "Token缺失",
	ERROR_JWT_TYPEWRONG:       "Token类型错误",
	ERROR_JWT_TIMEEXPIRED:     "Token过期",
	ERROR_JWT_NOTVALID:        "Token验证失败",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
