package main

import "fmt"

// A type struct
type A struct {
	name string
}

func (a A) GetName() string {
	a.name = "hello world!"
	return a.name
}

func main()  {
	a := A{} // type variable

	// 类型变量的调用方法
	s1 := a.GetName()

	// 类型调用方法传参类型变量
	s2 := A.GetName(a)

	fmt.Println(s1)
	fmt.Println(s2)
}
