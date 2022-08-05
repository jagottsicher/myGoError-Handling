package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("a-file.txt")
	if err != nil {
		fmt.Println("An error occured:", err)
		//		log.Println("An error occured:", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}
