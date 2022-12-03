package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	inventory := readInventoryFromFilePath(fileName)
	fmt.Println(inventory)
}

func readInventoryFromFilePath(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("can not read input file: %v", fileName)
		os.Exit(1)
	}

	calories := []int{}
	inventory := [][]int{}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			inventory = append(inventory, calories)
			calories = make([]int, 0)
		} else {
			textInt, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			calories = append(calories, textInt)
		}
	}
	// add remaining data
	inventory = append(inventory, calories)

	file.Close()
	return inventory
}
