package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// filePath := "day7/data/input.txt"
	// inputString := getContentOfFile(filePath)

	// maxSizePerDirectory := 100000
	// result, err :=
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Solution: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
