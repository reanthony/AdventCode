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
