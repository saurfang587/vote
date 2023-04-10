package tools

const (
	OK          = 0
	NotLogin    = 10001 //您还没有登录
	UserInfoErr = 10002 //用户信息错误
	DoErr       = 10003
)

type HttpCode struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
