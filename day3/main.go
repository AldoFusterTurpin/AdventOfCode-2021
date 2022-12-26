package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// filePath := "day3/data/input.txt"
	// inputString := getContentOfFile(filePath)

	// //
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
