package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetMaxSum(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		panic("Unexpected error while opening input file!")
	}
	currentMax := 0

	currentSum := 0
	id := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// skip the blank lines
			if currentSum != 0 {
				if currentSum > currentMax {
					currentMax = currentSum
				}
				currentSum = 0
				id += 1
			}
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			panic("Unexpected error while converting string to int")
		}
		currentSum += number
	}
	return currentMax
}

func GetSumOfBestThree(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		panic("Unexpected error while opening input file!")
	}

	var sums []int
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if sum != 0 {
				sums = append(sums, sum)
			}
			sum = 0
			continue
		}
		val, err := strconv.Atoi(t)
		if err != nil {
			panic("Unexpected error while converting string to int")
		}
		sum += val
	}
	maxHeap := NewMaxHeap(sums)
	sum = (*maxHeap)[0]
	maxHeap = NewMaxHeap((*maxHeap)[1:])
	sum += (*maxHeap)[0]
	maxHeap = NewMaxHeap((*maxHeap)[1:])
	sum += (*maxHeap)[0]
	return sum
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("Please provide path to input file!")
	}
	fmt.Println(args[0])

	fmt.Println("Most calories carried is ", GetMaxSum(args[0]))
	fmt.Println("Sum of top three: ", GetSumOfBestThree(args[0]))
}
