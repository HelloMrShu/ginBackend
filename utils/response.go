package utils

var code int
var msg string

//APIResponse 定义返回格式
func APIResponse(err error) (code int, msg string) {
	if err == nil {
		code = 200
		msg = "success"
	} else {
		code = 500
		msg = "failed"
	}

	return code, msg
}
