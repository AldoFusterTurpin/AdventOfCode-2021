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

	result, err := packet.DetectStartOfPacketMarker(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSolution: %v\n", result)
}

func getContentOfFile(filePath string) string {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
