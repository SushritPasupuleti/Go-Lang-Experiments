package main

import (
	"fmt"
	"strconv"
)

//outside of main function, variables must be properly declared
var globalI float32 = 10
var ExternalI int = 20 //exposed outside package

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
	fmt.Printf("Val: %v Type: %T\n", globalI, globalI)
	fmt.Printf("Val: %v Type: %T\n", a, a)
	fmt.Printf("Val: %v Type: %T\n", b, b)
	fmt.Printf("Val: %v Type: %T\n", c, c)

	i = i + 1
	fmt.Printf("Val: %v Type: %T\n", i, i)
	fmt.Printf("Val: %v Type: %T\n", ExternalI, ExternalI)

	//type coercion

	c = float32(c)

	fmt.Printf("Type Casted 'c' Val: %v Type: %T\n", c, c)

	//working with strings

	s := "Hello There"
	fmt.Println(s)

	var d = strconv.Itoa(i)	//convert int to string

	fmt.Printf("Val: %v Type: %T\n", d, d)

	sb := []byte(s)	//convert string to byte slice

	fmt.Printf("Val: %v Type: %T\n", sb, sb)
	//prints them as ASCII value array.
}
