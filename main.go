package main

import (
	"fmt"
	"strconv"
)

var (
	age  int    = 32
	name string = "senthil"
	//sex  string = "M"
)

func main() {
	fmt.Printf("%v\n%v\n", age, name)
	var (
		age  int    = 30
		name string = "senthil"
		//sex  string = "M"
	)

	test := 43.00
	test = 40.00
	var sex1 string
	name = strconv.Itoa(age)
	fmt.Printf("%v,%T\n", test, test)
	fmt.Printf("%v\n%v\n%v\n", age, name, sex1)
}
