package main

import "fmt"

// mainより先に呼ばれる。
// 初期設定で使用する
/*func init() {
	fmt.Println("Init")
}
*/
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	// ショートは関数内でしか定義できない
	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
