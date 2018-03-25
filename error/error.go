package error

import "fmt"

type SteamAPIError struct {
	Code int
	Msg  string
}

func (e *SteamAPIError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

var ErrorMap = map[int]string{
	400: "错误请求,请确认所有必需参数都已发送",
	401: "未经授权,访问被拒绝。重试无效。请验证您的 key= 参数",
	403: "禁止,访问被拒绝。重试无效。请验证您的 key= 参数",
	404: "未找到,请求的API不存在",
	405: "不允许使用此方法,调用此 API 的 HTTP 方法错误",
	429: "过多请求,您受到速率限制",
	500: "内部服务器错误，发生了无法恢复的错误，请重试",
	503: "服务器暂时不可用，或过于繁忙而无法响应。请稍后重试",
}

var CommonErr = "通用错误"