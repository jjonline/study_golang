package main

import (
	"fmt"
	"github.com/jjonline/study_golang/go_embed/three"
)

func main() {
	s, e := three.T5.ReadFile("stub/single.txt")
	fmt.Println(string(s))
	fmt.Println(e)
	s, e = three.T6.ReadFile("stubs/xyz/zyx.txt")
	fmt.Println(string(s))
	fmt.Println(e)
	fmt.Println(three.T5.ReadDir("stub"))
}
