package main

import (
	"fmt"
	"strconv"
)

var (
	x int     = 1
	y float64 = 1.2
	s string  = "14"
)

func foo() {
	xx := float64(x)
	yy := int(y)
	h := "Hello World"

	fmt.Printf("%T %v %d\n", yy, yy, yy)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	fmt.Println(string(h[0]))
}

func main() {
	foo()
	// strconv関数のAtoiメソッドで変換
	// 文字列→数値(逆も然り)はアスキー変換で対応
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T %v\n", i, i)
}
