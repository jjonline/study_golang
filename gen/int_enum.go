package gen

//go:generate echo "first"

// 通过为自定义类型Code的值定义时给注释备注而指定该code对应的文案
// int_x_gen.go 即为生成的go源码文件，可以看到里面实现了fmt.String接口
const (
	CodeErr        Code = iota + 1 // 错误
	CodeMsg                        // 提示
	CodePwdError                   // 密码错误
	CodeAccError                   // 账户错误
	CodeValidError                 // 参数错误
	CodeToRedirect                 // 跳转至
	CodeSuccess    Code = 500      // 登录成功
	CodeForbidden  Code = 600      // 禁止操作
	CodeNotAllow   Code = 700      // 无操作权限
)
