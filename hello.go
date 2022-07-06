package main

import (
	"fmt"
	"math"
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

	var d = strconv.Itoa(i) //convert int to string

	fmt.Printf("Val: %v Type: %T\n", d, d)

	sb := []byte(s) //convert string to byte slice

	fmt.Printf("Val: %v Type: %T\n", sb, sb)
	//prints them as ASCII value array.

	//"" for strings, '' for runes

	r := 'a'
	fmt.Printf("Val: %v Type: %T\n", r, r)

	//constants

	const pi float32 = 3.14

	fmt.Printf("Val: %v Type: %T\n", pi, pi)

	const actualPi = math.Pi

	fmt.Printf("Val: %v Type: %T\n", actualPi, actualPi)

	//arrays

	var arr [2]int

	arr[0] = 42
	arr[1] = 43

	fmt.Printf("Val: %v Type: %T\n", arr, arr)

	arr2 := [3]float32{1.1, 2.2, 3.3}

	fmt.Printf("Val: %v Type: %T\n", arr2, arr2)

	arr3 := [...]int{1, 2, 3} //dynamic array to fit the given data only

	arr3[0] = 42

	fmt.Printf("Val: %v Type: %T\n", arr3, arr3)

	arrCopy := arr3 //copy of arr3

	arrCopy[0] = 43

	fmt.Printf("Val: %v Type: %T\n", arr3, arr3)
	fmt.Printf("Val: %v Type: %T\n", arrCopy, arrCopy)

	//pointers

	arrPointer := &arr3

	arrPointer[0] = 44

	fmt.Printf("Val: %v Type: %T\n", arr3, arr3)
	fmt.Printf("Val: %v Type: %T\n", arrPointer, arrPointer)
}
