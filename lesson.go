package main

import "fmt"

// mainより先に呼ばれる。
// 初期設定で使用する
/*func init() {
	fmt.Println("Init")
}
*/

// Goのエントリポイントはmainに限定されるため、mainで呼び出す必要あり
func bazz() {
	fmt.Println("Bazz")
}

func main() {
	// bazz()
	fmt.Println("Hello World", "test test")
}
