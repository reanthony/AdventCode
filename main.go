package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum = 0
var maxSum = 0

func main() {
	// Day One
	maxSum, err := DayOne()
	if err != nil {
		panic(err)
	}
	fmt.Println(maxSum)

	//Day Two
	totalPoints, err := DayTwo()
	if err != nil {
		panic(err)
	}
	fmt.Println(totalPoints)

	// Day four
	result, err := DayFour()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func DayTwo() (points int, err error) {
	var priceMap map[string]int = make(map[string]int, 9)

	priceMap["A"], priceMap["X"] = 1, 1
	priceMap["B"], priceMap["Y"] = 2, 2
	priceMap["C"], priceMap["Z"] = 3, 3

	file, err := os.Open("./Input/DayTwo/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()
		currLine = strings.TrimSpace(currLine)
		opponent := priceMap[currLine[0:1]]
		me := priceMap[currLine[2:3]]
		pointsAwarded, err := calculateDayTwoPoints(opponent, me)
		if err != nil {
			panic(err)
		}
		points += pointsAwarded
	}
	return points, nil
}

func calculateDayTwoPoints(opponent int, me int) (pointsAwsrded int, err error) {
	//if draw
	if opponent == me {
		pointsAwsrded += me + 3
		return pointsAwsrded, nil
	}
	// declare winner
	won := declareWinnerDayTwo(opponent, me)
	if won {
		pointsAwsrded += 6 + me
	} else {
		pointsAwsrded += me
	}
	return pointsAwsrded, nil
}

func declareWinnerDayTwo(opponent int, me int) bool {
	if me == 1 {
		if opponent == 2 {
			return false
		}
		if opponent == 3 {
			return true
		}
	}
	if me == 2 {
		if opponent == 1 {
			return true
		}
		if opponent == 3 {
			return false
		}

	}
	if me == 3 {
		if opponent == 1 {
			return false
		}
		if opponent == 2 {
			return true
		}
	}
	return false
}

func DayOne() (maxSum int, err error) {
	file, err := os.Open("./Input/DayOne/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()
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
	return maxSum, nil
}

func DayFour() (totalSize int, err error) {
	file, err := os.Open("./Input/DayFour/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		currLine := scanner.Text()
		parts := strings.Split(currLine, ",")
		first := parts[0]
		second := parts[1]
		firstParts := strings.Split(first, "-")
		firstBegin := firstParts[0]
		firstEnd := firstParts[1]
		secondParts := strings.Split(second, "-")
		secondBegin := secondParts[0]
		secondEnd := secondParts[1]
		firstArry, secondArry := makeArrays(firstBegin, firstEnd, secondBegin, secondEnd)
		result = DayFourHelper(firstArry, secondArry, result)
	}

	return result, nil

}

func makeArrays(firstBegin string, firstEnd string, secondBegin string, secondEnd string) (firstArray []int, secondArray []int) {
	firstBeginInt, err := strconv.Atoi(firstBegin)
	if err != nil {
		panic(err)
	}
	firstEndInt, err := strconv.Atoi(firstEnd)
	if err != nil {
		panic(err)
	}
	secondBeginInt, err := strconv.Atoi(secondBegin)
	if err != nil {
		panic(err)
	}
	secondEndInt, err := strconv.Atoi(secondEnd)
	if err != nil {
		panic(err)
	}

	indx := 0
	firstArray = make([]int, (firstEndInt - firstBeginInt + 1))
	for i := firstBeginInt; i <= firstEndInt; i++ {
		firstArray[indx] = i
		indx++
	}

	indx = 0
	secondArray = make([]int, (secondEndInt - secondBeginInt + 1))
	for i := secondBeginInt; i <= secondEndInt; i++ {
		secondArray[indx] = i
		indx++
	}

	return firstArray, secondArray

}

func DayFourHelper(first []int, second []int, result int) int {

	//Check first is contained in second
	if isSubset(first, second) {
		result++
		return result
	}
	if isSubset(second, first) {
		result++
	}

	//Check second is contained in first
	return result
}

func isSubset(first []int, second []int) bool {
	var myMap map[int]bool = make(map[int]bool, len(first))
	for _, f := range first {
		myMap[f] = true
	}

	for _, s := range second {
		if _, ok := myMap[s]; !ok {
			return false
		}
	}

	// second is fully contained in first
	// second is a full subset of first
	return true
}
