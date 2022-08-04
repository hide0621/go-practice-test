package main

import "fmt"

func IsEven(num int) bool {
	return num%2 == 0
}

func main() {
	x := IsEven(-2)
	fmt.Println(x)
}
