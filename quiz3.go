package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindIndex(data []string, element string) int {
	for index, value := range data {
		// check the curr value is really unique from the unique array value
		if value == element {
			return index
		}
	}

	return -1
}

func FindAllIndex(data []string, element string) []int {
	allIndex := []int{}

	for index, value := range data {
		// check the curr value is really unique from the unique array value
		if value == element {
			allIndex = append(allIndex, index)
		}
	}

	return allIndex
}

func ArrayMerge(arrayA, arrayB []string) []string {
	// merge the array then sort by value (asc)
	mergedArr := append(arrayA, arrayB...)

	// validate the empty merged array
	if len(mergedArr) == 0 {
		return []string{}
	}

	// make an empty array to append the unique value
	var uniqueArr []string
	var isUniqueArr bool = true
	var lastValue string = mergedArr[len(mergedArr)-1]

	for i := 0; i < len(mergedArr)-1; i++ {
		var currValue string = mergedArr[i]
		var nextValue string = mergedArr[i+1]

		// only for the diff value from current and next
		if currValue != nextValue {
			// double checking for the unique value
			var indexOf = FindIndex(uniqueArr, currValue)

			// when the curr value is found in any index number of unique array then change the is unique arr to false then break
			if indexOf > -1 {
				isUniqueArr = false
				break
			}

			// only for true condition then append
			if isUniqueArr {
				uniqueArr = append(uniqueArr, currValue)
			}
		}
	}

	// only for the last condition
	var lastIndexOf = FindIndex(uniqueArr, lastValue)

	// check the index of last value and unique arr
	if lastIndexOf > -1 {
		isUniqueArr = false
	} else {
		isUniqueArr = true
	}

	// only for true condition then append
	if isUniqueArr {
		uniqueArr = append(uniqueArr, lastValue)
	}

	// return the unique arr val
	return uniqueArr
}

func Mapping(slice []string) map[string]int {
	// empty map for append it later
	var stringMapping = map[string]int{}

	for _, keyMap := range slice {
		// check the valid key and value for map
		_, isValid := stringMapping[keyMap]

		if !isValid {
			// create the map key value for the first time
			stringMapping[keyMap] = 1
		} else {
			// increment by 1 (based on map before)
			stringMapping[keyMap] += 1
		}
	}

	return stringMapping
}

func MunculSekali(angka string) []int {
	// split the string into the array format
	splitStr := strings.Split(angka, "")
	// for collecting the once value
	showOnce := []int{}

	for i := 0; i < len(splitStr)-1; i++ {
		// get currentVal and findAllIndex
		currentVal := splitStr[i]
		findAllIndex := FindAllIndex(splitStr, currentVal)

		// only check where the findAllIndex having 1 value (not twice or more)
		if len(findAllIndex) == 1 {
			// converting from string to integer
			intCurrentVal, err := strconv.Atoi(currentVal)

			// error handling
			if err == nil {
				// append it to showOnce
				showOnce = append(showOnce, intCurrentVal)
			}
		}

	}

	// for the last value (because odd array length)
	lastValue := splitStr[len(splitStr)-1]
	findLastValue := FindAllIndex(splitStr, lastValue)

	// only check where the findAllIndex having 1 value (not twice or more)
	if len(findLastValue) == 1 {
		// converting from string to integer
		intLastValue, err := strconv.Atoi(lastValue)

		// error handling
		if err == nil {
			// append it to showOnce
			showOnce = append(showOnce, intLastValue)
		}

	}

	// return the result
	return showOnce
}

func main() {
	// 1. Merge Array
	firstExample := ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"})
	secondExample := ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"})
	thirdExample := ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})
	fourthExample := ArrayMerge([]string{}, []string{"devil jin", "sergei"})
	fifthExample := ArrayMerge([]string{"hwoarang"}, []string{})
	sixthExample := ArrayMerge([]string{}, []string{})

	fmt.Println(firstExample)
	fmt.Println(secondExample)
	fmt.Println(thirdExample)
	fmt.Println(fourthExample)
	fmt.Println(fifthExample)
	fmt.Println(sixthExample)

	fmt.Println()
	// 2. Mapping
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))

	fmt.Println()
	// 3. Muncul Sekali
	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
}
