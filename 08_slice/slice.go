package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	fmt.Println(n[2])
	// :は全て？って意味でいいのか？
	// :2だと、index = 2までの全て
	// →1,2
	// 2:、index = 2（を含む）残り全て
	// →3,4,5
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)
}
