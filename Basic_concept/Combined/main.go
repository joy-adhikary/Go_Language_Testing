package main

import "fmt"

func Calculate(x int, y int) (result int) {
	result = x + y
	result *= 2
	return result
}

func cal(x int) (result int) {
	result = x + 2
	return result
}

// func sum(a int, b int) (x int) {
// 	x = a + b
// 	return x
// }

func info(first string, last string, age int) (name string) {
	name = first + last
	return name
}

func main() {
	fmt.Println("finding the working pincipals of go test")
	result := Calculate(5, 10)
	fmt.Println(cal(3))
	//fmt.Println(sum(1, 1))
	fmt.Println(result)
	result1 := info("joy", "adhikary", 23)
	fmt.Println(result1)

}
