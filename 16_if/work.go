package main

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		println("great 2")
	}

	// num := 6
	// if num%2 == 0 {
	// 	println("by 2")
	// } else if num%3 == 0 {
	// 	println("by 3")
	// } else {
	// 	println("else")
	// }

	x, y := 11, 12
	if x == 10 && y == 10 {
		println("&&")
	}

	if x == 10 || y == 10 {
		println("||")
	}
}
