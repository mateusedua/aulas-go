package main

import "fmt"

func main() {
	var array = [3]int{1, 2, 3}

	for _, v := range array {
		fmt.Println(v)
	}
}
