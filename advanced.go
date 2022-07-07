package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://skillshack.dev/robots.txt")
	if err != nil {
		log.Fatal(err)
		// panic(err.Error())
		panic("Error in Get")
	}
	defer res.Body.Close()//close body after function is done

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		// panic(err.Error())//equivalent of throwing an exception
		panic("Error in Parsing Body")
	}

	fmt.Println(string(body))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in function", r)
		}
	}()

	panic("no reason to panic, but I'm panicing anyway")
}

