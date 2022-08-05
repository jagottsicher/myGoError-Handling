package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, 世界")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(n, "Bytes written.")
}

