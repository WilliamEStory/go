package main

import (
	"fmt"
)

func main() {
	ten := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range ten {
		if value%2 == 0 {
			fmt.Printf("%v is even \n", value)
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
