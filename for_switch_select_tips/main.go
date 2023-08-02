package main

import "fmt"

// range 遍历数组的拷贝
func rangeTest() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a) // original a = [1 2 3 4 5]

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r) // after for range loop, r = [1 2 3 4 5]
	fmt.Println("after for range loop, a =", a) // after for range loop, a = [1 12 13 4 5]
}

func main() {
	rangeTest()
	breakInnerTest()
	breakOuterTest()
}

// break 中断内层循环
func breakInnerTest() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	// find first even number of the interger slice
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break
		case 1:
			// do nothing
		}
	}
	println(firstEven) // 12
}

func breakOuterTest() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	// find first even number of the interger slice
outer_loop:
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break outer_loop
		case 1:
			// do nothing
		}
	}
	println(firstEven) // 6
}
