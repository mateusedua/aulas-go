package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(cap(s))
	s = append(s, 12)
	fmt.Println(cap(s))
	//Forma de tralabahar com o slice
	fmt.Println(s[0:1])
}
