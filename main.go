package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	testValidity("23-ab-48-caba-56-haha")
	fmt.Println(averageNumber("23-ab-48-caba-56-haha"))
	fmt.Println(wholeStory("23-ab-48-caba-56-haha"))
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
