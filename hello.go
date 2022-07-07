package main

import (
	"fmt"
	"math"
	"reflect"
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

	//slices
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Val: %v Type: %T Len: %v Capacity: %v \n", slice, slice, len(slice), cap(slice))

	//copies of slices all point to the master slice. This is the same as the pointer to the master slice. Same case for maps.

	//maps

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}

	map1["d"] = 4
	map1["e"] = 5

	delete(map1, "b")

	//return order of maps is not guaranteed.

	fmt.Printf("Val: %v Type: %T\n", map1, map1)

	k, ok := map1["b"] //ok is true if the key is in map.

	fmt.Printf("Val: %v Type: %T\n", k, k)
	fmt.Printf("Val: %v Type: %T\n", ok, ok)

	//structs

	type Person struct {
		name string
		age  int
		id   int
	}

	p := Person{"John", 42, 1}
	p2 := Person{
		name: "Jane",
		age:  43,
		id:   2,
	}

	fmt.Printf("Val: %v Type: %T\n", p, p)
	fmt.Printf("Val: %v Type: %T\n", p2, p2)
	fmt.Printf("Val: %v Type: %T\n", p.name, p.name)

	//anonymous structs

	aPerson := struct {
		name string
		age  int
		id   int
	}{
		name: "John",
		age:  42,
		id:   1,
	}

	fmt.Printf("Val: %v Type: %T\n", aPerson, aPerson)

	//composition (similar to inheritance)

	type Lifeform struct {
		name        string
		age         int
		carbonBased bool
	}

	type Bird struct {
		Lifeform  //embedded struct
		legs      int
		beakColor string
	}

	blueJay := Bird{
		Lifeform: Lifeform{
			name:        "Blue Jay",
			age:         3,
			carbonBased: true,
		},
		legs:      2,
		beakColor: "blue",
	}

	fmt.Printf("Val: %v Type: %T\n", blueJay, blueJay)

	//struct tags

	type Person2 struct {
		name string `required max:"50"` //validation
		age  int    `min:"18"`
	}

	p3 := Person2{
		name: "John",
		age:  42,
	}

	t := reflect.TypeOf(p3)
	v := reflect.ValueOf(p3)

	fmt.Printf("Val: %v Type: %T\n", p3, p3)
	fmt.Printf("Val: %v Type: %T\n", t, t)
	fmt.Printf("Val: %v Type: %T\n", v, v)

	if val, ok := map1["b"]; ok {
		fmt.Printf("Val: %v Type: %T\n", val, val)
	} else {
		fmt.Println("Key not found")
	}
}
