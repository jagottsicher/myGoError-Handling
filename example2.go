package main

import "fmt"

func main() {
	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Name:", answer)
}
