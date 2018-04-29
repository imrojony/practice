//chapter 4
//package main
//
//import "fmt"
//
//func add(x,y int) int {
//	return x+y
//
//}
//func main()  {
//	fmt.Println(add(13,23))
//
//}
//chapter 5
//package main
//
//import "fmt"
//
//func add(x, y int) int {
//	return x + y
//}
//
//
//func main() {
//	fmt.Println(add(42, 13))
//}
//chapter 6
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
