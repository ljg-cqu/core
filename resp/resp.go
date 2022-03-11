package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/utils"
	"net/http"
)

const (
	// Fields around account
	UserId   FieldKey = "user_id"
	UserName FieldKey = "user_name"

	// Fields around authentication
	AccessJwt  = "access_jwt"
	RefreshJwt = "refresh_jwt"

	// Fields around user request
	RequestId  FieldKey = "request_id"
	RequestUrl FieldKey = "request_url"

	// Fields around server
	ServerAddress FieldKey = "server_addr"
	ServerPort    FieldKey = "server_port"

	// Fields around Antchain

	// Fields around Esign
)

// HTTP status codes as registered with IANA. See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 7231, 6.3.1
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently = 301 // RFC 7231, 6.4.2
	StatusFound            = 302 // RFC 7231, 6.4.3
	StatusSeeOther         = 303 // RFC 7231, 6.4.4
	StatusNotModified      = 304 // RFC 7232, 4.1
	StatusUseProxy         = 305 // RFC 7231, 6.4.5

	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var (
	OK = &ErrCodeMsg{200, "ok", nil}

	ErrBadRequest   = &ErrCodeMsg{StatusBadRequest, http.StatusText(StatusBadRequest), nil}     // 400
	ErrUnauthorized = &ErrCodeMsg{StatusUnauthorized, http.StatusText(StatusUnauthorized), nil} // 401

	ErrInternalServerError = &ErrCodeMsg{StatusInternalServerError, http.StatusText(StatusInternalServerError), nil} // 500
)

//var (
//	// logon status
//	ErrNotLoginYet  = &ErrCodeMsg{11, `还未登录或登录已过期！`}
//	ErrTokenInvalid = &ErrCodeMsg{12, `登录令牌验证失败！`}
//	ErrTokenExpired = &ErrCodeMsg{13, `登录令牌已过期，请重新登录！`}
//
//	//
//	ErrAccessDenied = &ErrCodeMsg{21, `未授权的访问请求！`}
//
//	// db
//	ErrDbInsert  = &ErrCodeMsg{1001, `网络开小差了，请重试(1001)`} //数据插入错误
//	ErrDbDelete  = &ErrCodeMsg{1002, `网络开小差了，请重试(1002)`} //数据删除错误
//	ErrDbUpdate  = &ErrCodeMsg{1003, `网络开小差了，请重试(1003)`} //数据修改错误
//	ErrDbQuery   = &ErrCodeMsg{1004, `网络开小差了，请重试(1004)`} //数据查询错误
//	ErrDbGenID   = &ErrCodeMsg{1005, `网络开小差了，请重试(1005)`} //数据生成ID错误
//	ErrDbConnect = &ErrCodeMsg{1011, `网络开小差了，请重试(1011)`} //数据暂不可用，请稍后再试！
//	ErrNotFound  = &ErrCodeMsg{1012, `网络开小差了，请重试(1012)`} //数据不存在！
//
//	//redis 2001~2999
//	ErrRedisConn  = &ErrCodeMsg{2001, `网络开小差了，请重试(2001)`} //
//	ErrRedisSave  = &ErrCodeMsg{2002, `网络开小差了，请重试(2002)`} //
//	ErrRedisQuery = &ErrCodeMsg{2003, `网络开小差了，请重试(2003)`} //
//	ErrRedisParse = &ErrCodeMsg{2004, `网络开小差了，请重试(2004)`} //
//
//)
//
//// errors about web app gateway
//var (
//	ErrParamsInvalid       = &ErrCodeMsg{10001, `请求数据错误！`}
//	ErrPhoneInvalid        = &ErrCodeMsg{10002, `手机号格式不正确！`}
//	ErrSMSCodeInvalid      = &ErrCodeMsg{10003, `验证码格式不正确！`}
//	ErrPasswordLength      = &ErrCodeMsg{10011, `请设置一个6~18位长度的密码！`}
//	ErrPasswordInclude     = &ErrCodeMsg{10012, `密码需要包含至少一个字母及一个数字！`}
//	ErrLoginType           = &ErrCodeMsg{10021, `暂不支持该登录方式！`}
//	ErrNameInvalid         = &ErrCodeMsg{10101, `用户名格式不正确，长度为3-16字节！`}
//	ErrAddAcc              = &ErrCodeMsg{10102, `账号添加失败！`}
//	ErrAccNoExist          = &ErrCodeMsg{10103, `账号不存在！`}
//	ErrPassword            = &ErrCodeMsg{10104, `登录密码错误！`}
//	ErrConfigUndefine      = &ErrCodeMsg{10201, `未定义的配置类型！`}
//	ErrConfigJsonDeCode    = &ErrCodeMsg{10202, `配置字符串json解析失败！`}
//	ErrConfigFloat64DeCode = &ErrCodeMsg{10202, `配置字符串float64解析失败！`}
//)

type ErrCodeMsg struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Fileds Fields `json:"fileds"`
}

type RespObj struct {
	ErrCodeMsg
	Data string `json:"data"`
}

type FieldKey string
type Fields map[FieldKey]string

func (e *ErrCodeMsg) AppendMsg(msg string) *ErrCodeMsg {
	e.Msg = e.Msg + ":" + msg
	return e
}

func (e *ErrCodeMsg) WithFields(fields map[FieldKey]string) *ErrCodeMsg {
	e.Fileds = fields
	return e
}

func Succ(c *gin.Context, data interface{}) {
	dataStr, _ := utils.Json.MarshalToString(data)
	resp(c, &RespObj{
		ErrCodeMsg: *OK,
		Data:       dataStr,
	})
}

func Err(c *gin.Context, codeMsg *ErrCodeMsg) {
	resp(c, &RespObj{
		ErrCodeMsg: *codeMsg,
	})
}

func resp(c *gin.Context, respObj *RespObj) {
	c.JSON(http.StatusOK, respObj)
}
