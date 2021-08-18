package response

const (
	ERR_OK     = 0 // 成功
	ERR_FAIL   = 1 // 失败
	ERR_STRUCT = 2 // struct解析错误

	// auth 400-500
	ERR_AUTH_SIGN_UP = 401 // 注册失败
)
