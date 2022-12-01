package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sum = 0
var maxSum = 0

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		DayOne(text)
	}

	fmt.Println(maxSum)
}

func DayOne(currLine string) {
	if currLine == "" {
		if sum > maxSum {
			maxSum = sum
		}
		sum = 0
		return
	}

	line, err := strconv.ParseInt(currLine, 10, 64)
	if err != nil {
		panic(err)
	}

	sum += int(line)
}
