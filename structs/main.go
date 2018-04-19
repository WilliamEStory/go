package main

import (
	"fmt"
)

type person struct {
	fistName string
	lastName string
}

func main() {
	alex := person{fistName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
