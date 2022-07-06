package main

import "fmt"

//outside of main function, variables must be properly declared
var gloabl_i float32 = 10
var External_i int = 20 //exposed outside package

//All variables with a starting lowercase character in name are not exposed outside of package

//block variable declaration

var (
	a         = 42
	b         = 43
	c float32 = 44.44
) //all vars

func main() {

	var i int //explicitly declaring the type of variable to integer

	i = 42

	x := 43 //declare and will be type casted to integer

	fmt.Println("Hello There", i, x)
	fmt.Printf("Val: %v Type: %T\n", i, i)
	fmt.Printf("Val: %v Type: %T\n", x, x)
	fmt.Printf("Val: %v Type: %T\n", gloabl_i, gloabl_i)
	fmt.Printf("Val: %v Type: %T\n", a, a)
	fmt.Printf("Val: %v Type: %T\n", b, b)
	fmt.Printf("Val: %v Type: %T\n", c, c)

	i = i + 1
	fmt.Printf("Val: %v Type: %T\n", i, i)
	fmt.Printf("Val: %v Type: %T\n", External_i, External_i)
}
