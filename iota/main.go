package main

import "log"

// 递增常量
const (
	Apple  = iota // 0
	Banana        // 1
	Cherry        // 2
)

// 枚举模拟
type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// 表达式计算
const (
	_  = iota
	KB = 1 << (10 * iota) // 1024 = 1 << 10
	MB = 1 << (10 * iota) // 1048576 = 1 << 20
	GB = 1 << (10 * iota) // 1073741824 = 1 << 30
	TB = 1 << (10 * iota) // 1099511627776 = 1 << 40
)

// 位运算
const (
	FlagNone  = 0         // 0
	FlagRead  = 1 << iota // 2
	FlagWrite             // 4
	FlagExec              // 8
)

// 起始值
const (
	A = 5
	B = 4
	C = iota // 2
	D        // 3
)

// 跳值
const (
	X = iota // 0
	_        //
	Y        // 2
)

func main() {
	log.Println("----递增常量----")
	log.Println(Apple)  // 0
	log.Println(Banana) // 1
	log.Println(Cherry) // 2

	log.Println("----枚举模拟----")
	printSeason(Autumn)

	log.Println("----表达式计算----")
	log.Println(KB)
	log.Println(MB)
	log.Println(GB)
	log.Println(TB)

	log.Println("----位运算----")
	log.Println(FlagNone)
	log.Println(FlagRead)
	log.Println(FlagWrite)
	log.Println(FlagExec)

	log.Println("----起始值----")
	log.Println(A)
	log.Println(B)
	log.Println(C)
	log.Println(D)

	log.Println("----跳值----")
	log.Println(X)
	log.Println(Y)
}

func printSeason(s Season) {
	switch s {
	case Spring:
		log.Println("Spring")
	case Summer:
		log.Println("Summer")
	case Autumn:
		log.Println("Autumn")
	case Winter:
		log.Println("Winter")
	}
}
