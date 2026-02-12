package main

import (
	"fmt"
	"strings"
)

func isUpper(a string) string {
	return strings.ToUpper(a)
}

func uniqueStringChars(str string) int {
	seen := make(map[rune]bool)
	countSeen := make(map[rune]int)

	for _, r := range str {
		seen[r] = true
	}
	fmt.Println(seen)
	for _, r := range str {
		//Check if key exists in map
		if _, exists := countSeen[r]; exists {
			countSeen[r] = countSeen[r] + 1
		} else {
			countSeen[r] = 1
		}
	}
	fmt.Println("Count of each char inside string " + str)
	for char, count := range countSeen {
		fmt.Printf("    '%c' : '%d' \n", char, count)
	}
	fmt.Printf("Count of each char is %v \n", countSeen)
	return len(seen)
}

func getVowels(str string) {
	vowel := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
	for _, r := range str {
		if _, exists := vowel[r]; exists {
			vowel[r] = vowel[r] + 1
		}
	}
	totalVowels := 0
	for char, count := range vowel {
		fmt.Printf("    '%c' : '%d' \n", char, count)
		totalVowels += count
	}
	fmt.Printf("Total vowels in string %s is %d \n", str, totalVowels)
}
func getMaximum(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	fmt.Printf("Sorted list: %v\n", list)
	fmt.Printf("Maximum value: %d\n", list[len(list)-1])
}
func getMaximum2(list1 []int) {
	max := list1[0]
	for i := range len(list1) {
		if list1[i] > max {
			max = list1[i]
		}
	}
	fmt.Printf("Max value in list1 %v is %d \n", list1, max)
}
func main() {
	fmt.Println(isUpper("Hello"))
	fmt.Printf("Unique chars in string %s is %d \n", "Hello", uniqueStringChars("Hello"))
	fmt.Printf("Vowels inside string %s is as follows\n", "Hello")
	getVowels("Hello")
	getMaximum([]int{5, 4, 7, 6, 1})
	getMaximum2([]int{3, 4, 9, 1, 3, 10})
}
