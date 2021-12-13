package gen

//go:generate $GOPATH/bin/stringer -type Code -linecomment  -output int_x_gen.go
type Code int
