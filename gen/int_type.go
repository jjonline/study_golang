package gen

//go:generate $GOPATH/bin/stringer -type Code -linecomment  -output int_x_gen.go
//go:generate ls -ll
//go:generate echo "two"

// Code 自定义INT型类型
// 通过stringer指令为Code添加了String方法实现
// 那么直接通过fmt.Println(c)时会调用输出String方法而不是c值本身
type Code int
