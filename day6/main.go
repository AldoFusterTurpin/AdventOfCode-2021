package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day6/packet"
)

func main() {
	filePath := "day6/data/input.txt"
	inputString := getContentOfFile(filePath)

	// For part 1
	// result, err := packet.DetectStartOfPacketMarker(inputString)

	// For part 2
	result, err := packet.DetectStartOfMessageMarker(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
