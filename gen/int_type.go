package gen

//go:generate $GOPATH/bin/stringer -type Code -linecomment  -output int_x_gen.go
//go:generate ls -ll
//go:generate echo "two"

// Code 自定义INT型类型
// 通过stringer指令为Code添加了String方法实现
// 那么直接通过fmt.Println(c)时会调用输出String方法而不是c值本身
type Code int

// Error 为 Code 类型实现 error 接口，这样 Code 类型变量可以直接错误 error 返回契合go语言的返回错误特征
//  stringer 自动生成为 Code 类型实现了 fmt.Stringer 接口 自动具备 String 方法
func (c Code) Error() string {
	return c.String()
}

// Code 为 Code 类型的值转换数字错误编码提供快捷方法
func (c Code) Code() int {
	return int(c)
}
