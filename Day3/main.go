package main

import (
	"bufio"
	"fmt"
	"os"
)

// Each rucksack has two compartment
// one type of items must go into only one compartment!!!
// items are letters of the alphabet, case matters!

// find which item is in both compartments
// sum priorities of all misplaced types of items

func getPriotity(c uint8) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func Task1(filepath string) (score int) {
	score = 0
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		half := len(line) / 2
		comp1 := line[0:half]
		comp2 := line[half:]

		fmt.Printf("1: %s, 2: %s", comp1, comp2)

		var letterMap [58]uint8
		for i := 0; i < len(comp1); i++ {
			letterMap[comp1[i]-'A'] = 1
		}
		for i := 0; i < len(comp2); i++ {
			if letterMap[comp2[i]-'A'] == 1 {
				fmt.Printf(", repeats: %c\n", comp2[i])
				score += getPriotity(comp2[i])
				break
			}
		}
	}

	return score
}

func findRepeatingCharacter(group [3]string) uint8 {
	var letterMap [58]uint8
	for i := 0; i < 3; i++ {
		s := group[i]
		for j := 0; j < len(s); j++ {
			if letterMap[s[j]-'A'] == uint8(i) {
				letterMap[s[j]-'A'] = uint8(i + 1)
			}
			if letterMap[s[j]-'A'] == 3 {
				return s[j]
			}
		}
	}
	return 0
}

// badge is the only item carried by all 3 elves in a group
// each group is representted by three line
// read 3 lines, find the item that repeats and count the priority
func Task2(filepath string) int {
	score := 0
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var group [3]string
	for i := 0; scanner.Scan(); i++ {
		group[i%3] = scanner.Text()

		if i%3 != 2 {
			continue
		}

		c := findRepeatingCharacter(group)
		score += getPriotity(c)
		fmt.Printf("Group: %s, %s, %s, repeats: %c\n", group[0], group[1], group[2], c)

	}
	return score
}

func main() {
	fmt.Println("Task1 is: ", Task1("input.txt"))
	fmt.Println("Task2: ", Task2("input.txt"))
}
