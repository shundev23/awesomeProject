package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// keyがstring型、valueがint型の空のmapを作成
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// メモリー確保していないのでエラー
	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	s := []int{}
	if s == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println("slice is not nil")
	}

	// マップからキー"apple"の値を取得
	// vに値、okに存在するかどうかの真偽値が入る
	v, ok = m["apple"]
	fmt.Println(v, ok) // 出力: 100 true

	// 存在しないキー"orange"の値を取得
	v2, ok2 = m["orange"]
	fmt.Println(v2, ok2) // 出力: 0 false
}
