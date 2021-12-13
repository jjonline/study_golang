package gen

//go:generate $GOPATH/bin/stringer -type Code -linecomment  -output int_x_gen.go
//go:generate ls -ll
type Code int
