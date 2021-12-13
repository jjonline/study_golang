package gen

//go:generate echo "first"

const (
	CodeErr      Code = iota + 1 // 错误
	CodeMsg                      // 提示
	CodePwdError                 // 密码错误
	CodeSuccess  Code = 500      // 登录成功
)
