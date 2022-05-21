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
	// code = 30** 分类模块错误
	ERROR_CATEGORY_EXIST       = 3001
	ERROR_CATEGORY_ADD_FAIL    = 3002
	ERROR_CATEGORY_DELETE_FAIL = 3003
	ERROR_CATEGORY_EDIT_FAIL   = 3004
	ERROR_CATEGORY_NOT_EXIST   = 3005
	ERROR_CATEGORY_LIST_FAIL   = 3006
	// code = 40** 菜单栏相关
	ERROR_MENU_ADD = 4001
	// code = 50** 权限相关
	ERROR_ROLE_ADD = 5001
)

var codemsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	//用户相关
	ERROR_PASSWORD_WRONG:      "用户名或密码错误",
	ERROR_USERNAME_EXIST:      "用户已存在",
	ERROR_TOKEN_WRONG:         "用户TOKEN错误",
	ERROR_TOKEN_RUNTIME:       "用户TOKEN过期",
	ERROR_USERLIST_WRONG:      "用户列表获取失败",
	ERROR_USEREDIT_WRONG:      "用户更新失败",
	ERROR_USERDELETE_NOTFOUND: "用户不存在",
	ERROR_USERDELETE_WRONG:    "用户删除失败",
	// JWT token相关
	ERROR_JWT_MISSING:     "Token缺失",
	ERROR_JWT_TYPEWRONG:   "Token类型错误",
	ERROR_JWT_TIMEEXPIRED: "Token过期",
	ERROR_JWT_NOTVALID:    "Token验证失败",
	// 分类相关
	ERROR_CATEGORY_EXIST:       "分类已存在",
	ERROR_CATEGORY_ADD_FAIL:    "分类新建失败",
	ERROR_CATEGORY_DELETE_FAIL: "分类删除失败",
	ERROR_CATEGORY_EDIT_FAIL:   "分类编辑失败",
	ERROR_CATEGORY_NOT_EXIST:   "分类不存在",
	ERROR_CATEGORY_LIST_FAIL:   "分类列表获取失败",
	// Menu相关
	ERROR_MENU_ADD: "菜单添加错误",
	// 权限相关
	ERROR_ROLE_ADD: "权限添加错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
