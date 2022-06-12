package main

import "fmt"

var x int

type customType int

var y customType

func main() {
	fmt.Println("hello golang")
	z := "first string variable"
	x = 12345
	y = 98765
	fmt.Printf("%T %v \n", x, x)
	fmt.Printf("%T %v \n", y, y)
	fmt.Printf("%T %v \n", z, z)
	foo()

}
func foo() {
	fmt.Println("foo func calling")
}
