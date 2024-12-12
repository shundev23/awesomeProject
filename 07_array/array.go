package main

import "fmt"

func main() {
	// 通常の書き方
	// int型配列、サイズが2
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// ブランケットでまとめて定義
	// var b [2]int = [2]int{100, 200}
	// fmt.Println(b)

	// スライスはサイズ未指定の配列と覚える
	// これにより、サイズ変更が可能になる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
