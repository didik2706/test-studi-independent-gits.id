package main

import "fmt"

func main() {
	var value int16

	fmt.Print("Masukkan value : ")
	fmt.Scan(&value)

	if value%3 == 0 && value%5 == 0 {
		fmt.Println("Hello World")
	} else if value%3 == 0 {
		fmt.Println("Hello")
	} else if value%5 == 0 {
		fmt.Println("World")
	}
}
