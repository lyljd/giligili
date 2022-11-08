package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Error string      `json:"error,omitempty"`
}

// 只要服务器接收到请求都统一返回200，至于状态码在返回的code中体现
// 0是正常响应（默认值，可以不用设置），4开头的都是客户端出了问题，而5开头的都是服务器出了问题
const (
	CodeNoLogin           = 401   // CodeNoLogin 没有通过身份验证
	CodeNoPower           = 403   // CodeNoPower 通过身份验证，但没有权限操作该资源
	CodeNotFound          = 404   // CodeNotFound 通过身份验证，但没有找到需要的资源
	CodeRefreshTokenErr   = 4001  // CodeRefreshTokenErr 通过RefreshToken刷新Token时出现问题
	CodeResourceExpireErr = 4003  // CodeResourceExpireErr 不验证身份，请求的资源已过期
	CodeCliErr            = 40000 // CodeCliErr 客户端出现问题通用报错
	CodeCliParErr         = 40001 // CodeCliParErr 客户端传入参数有问题
	CodeSerErr            = 50000 // CodeSerErr 服务器出现问题通用报错
	CodeSerDbErr          = 50001 // CodeSerDbErr 服务器操作数据库出现问题
)

// Err 通用错误处理（为了在生产环境下隐藏Err）
func Err(code int, msg string, err error) Response {
	res := Response{
		Code: code,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// NoLogin 没有通过身份验证
func NoLogin(msg string, err error) Response {
	if msg == "" {
		msg = "请登录"
	}
	return Err(CodeNoLogin, msg, err)
}

// NoPower 通过身份验证，但没有权限操作该资源
func NoPower(msg string, err error) Response {
	if msg == "" {
		msg = "权限不足"
	}
	return Err(CodeNoPower, msg, err)
}

// NotFound 通过身份验证，但没有找到需要的资源
func NotFound(msg string, err error) Response {
	if msg == "" {
		msg = "未找到相应资源"
	}
	return Err(CodeNotFound, msg, err)
}

// RefreshTokenErr 通过RefreshToken刷新Token时出现问题
func RefreshTokenErr(msg string, err error) Response {
	if msg == "" {
		msg = "请登录"
	}
	return Err(CodeRefreshTokenErr, msg, err)
}

// ResourceExpireErr 不验证身份，请求的资源已过期
func ResourceExpireErr(err error) Response {
	return Err(CodeResourceExpireErr, "资源已过期", err)
}

// CliErr 客户端出现问题通用报错
func CliErr(err error) Response {
	return Err(CodeCliErr, "客户端未知错误", err)
}

// CliParErr 客户端传入参数有问题
func CliParErr(msg string, err error) Response {
	if msg == "" {
		msg = "传入参数有误"
	}
	return Err(CodeCliParErr, msg, err)
}

// SerErr 服务器出现问题通用报错
func SerErr(err error) Response {
	return Err(CodeSerErr, "服务器未知错误", err)
}

// SerDbErr 服务器操作数据库出现问题
func SerDbErr(err error) Response {
	return Err(CodeSerDbErr, "操作数据库失败", err)
}
