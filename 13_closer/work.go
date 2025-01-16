package main

import "fmt"

/*
クロージャーの例1: incrementGenerator
- 外部変数 x を保持し続ける
- 関数が呼ばれるたびに x の値を更新
- x の状態は呼び出し間で維持される
*/
func incrementGenerator() func() int {
	x := 0 // この変数xはクロージャーによって保持される
	return func() int {
		// この無名関数は x にアクセスでき、その値を記憶し続ける
		// ブレークポイント1: クロージャが呼ばれる前のxの値を確認
		currentX := x
		fmt.Printf("Before increment: x = %d\n", currentX)

		// ブレークポイント2: インクリメント処理
		x++

		// ブレークポイント3: インクリメント後のxの値を確認
		fmt.Printf("After increment: x = %d\n", x)
		return x
	}
}

/*
クロージャーの例2: circleArea
- 外部変数 pi を保持
- 異なる pi 値で複数のインスタンスを作成可能
- 半径を受け取り面積を計算
*/
func circleArea(pi float64) func(radius float64) float64 {
	// pi の値はクロージャーによって保持される
	return func(radius float64) float64 {
		return pi * radius * radius // 保持された pi を使用
	}
}

func main() {
	// incrementGeneratorの使用例
	// ブレークポイント4: カウンター関数の生成
	counter := incrementGenerator()
	// 各呼び出しで状態（x）が維持される

	fmt.Println("--- First call ---")
	// ブレークポイント5: 1回目の呼び出し
	result1 := counter()
	fmt.Println(result1)

	fmt.Println("--- Second call ---")
	// ブレークポイント6: 2回目以降の呼び出し
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// circleAreaの使用例
	c1 := circleArea(3.14) // pi=3.14 のインスタンス
	fmt.Println(c1(2))     // 半径2で計算

	c2 := circleArea(3) // pi=3 の別インスタンス
	fmt.Println(c2(2))  // 同じ半径でも異なる結果
}
