package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, x := range i {
		if x%2 == 0 {
			fmt.Printf("%v is even\n", x)
		} else {
			fmt.Printf("%v is odd\n", x)
		}
	}

}
