package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/bradfitz/iter"
)

func main() {
	fmt.Println(generateStrings(false))
	testValidity("23-ab-48-caba-56-haha")
	fmt.Println(averageNumber("23-ab-48-caba-56-haha"))
	fmt.Println(wholeStory("23-ab-48-caba-56-haha"))
	fmt.Println(storyStats("23-ab-48-caba-56-haha"))
}

/*
Estimated Time: 10 mins
Used Time: 5 mins able to generate from https://regexr.com/
*/
func testValidity(inputStr string) bool {
	r, err := regexp.Compile(`\d*-[A-Za-z]*`)
	if err != nil {
		fmt.Println("Failed to compile regexp", err)
	}
	return r.MatchString(inputStr)
}

/*
Estimated Time: 2 mins
Used Time: 2 mins
*/
func averageNumber(inputStr string) int {
	sum := 0
	length := 0
	stringArray := strings.Split(inputStr, "-")
	for index, val := range stringArray {
		if index%2 == 0 {
			num, _ := strconv.Atoi(val)
			sum += num
			length += 1
		}
	}
	return sum / length
}

/*
Estimated Time: 2 mins
Used Time: 2 mins
*/
func wholeStory(inputStr string) string {
	var story string
	stringArray := strings.Split(inputStr, "-")
	for index, val := range stringArray {
		if index%2 != 0 {
			story += val + " "
		}
	}
	return strings.TrimSuffix(story, " ")
}

/*
Estimated Time: 10 mins
Used Time: 15 mins
*/
func storyStats(inputStr string) (int, int, int, []string) {
	return findShortestWordLength(inputStr), findLongestWordLength(inputStr), averageWordLength(inputStr), wordsWithAverageWordLength(inputStr)
}

/*
Estimated Time: 15 mins
Used Time: 30 mins
*/
func generateStrings(flag bool) string {
	var list []string
	num := rand.Intn(15)
	if flag {
		for i := range iter.N(num) {
			fmt.Println(i)
			list = append(list, strconv.Itoa(rand.Intn(5)))
			list = append(list, randStringBytes(3))
		}
	} else {
		for i := range iter.N(num) {
			fmt.Println(i)
			list = append(list, strconv.Itoa(rand.Intn(5)))
		}
	}
	return strings.Join(list, "-")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func findShortestWordLength(s string) (c int) {
	arr := wholeStory(s)
	words := strings.Split(arr, " ")
	for _, word := range words {
		if c == 0 || len(word) < c {
			c = len(word)
		}
	}
	return
}

func findLongestWordLength(s string) (c int) {
	arr := wholeStory(s)
	words := strings.Split(arr, " ")
	for _, word := range words {
		if c == 0 || len(word) > c {
			c = len(word)
		}
	}
	return
}

func averageWordLength(s string) (c int) {
	totalWordLength := 0
	length := 0
	arr := wholeStory(s)
	words := strings.Split(arr, " ")
	for _, word := range words {
		totalWordLength += len(word)
		length += 1
	}
	return totalWordLength / length
}

func wordsWithAverageWordLength(s string) (list []string) {
	arr := wholeStory(s)
	avgWordLen := averageWordLength(s)
	words := strings.Split(arr, " ")
	for _, word := range words {
		if len(word) == avgWordLen {
			list = append(list, word)
		}
	}
	return
}
