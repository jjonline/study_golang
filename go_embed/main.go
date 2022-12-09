package main

import (
	"fmt"
	"github.com/jjonline/study_golang/go_embed/three"
)

func main() {
	s, e := three.T3.ReadFile("b-.txt")
	fmt.Println(string(s))
	fmt.Println(e)
}
