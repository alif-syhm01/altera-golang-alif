package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(ArrayMerge([]string{"alterra", "academy", "malang"}, []string{"academy", "jakarta"}))

	// fmt.Println(Mapping([]string{"alta", "malang", "alta", "jakarta", "malang"}))

	fmt.Println(MunculSekali("1234321"))
}

func MunculSekali(angka string) []int {
	tempMap := map[string]int{}
	result := []int{}

	for _, v := range angka {
		tempMap[string(v)] += 1
	}

	for k, v := range tempMap {
		if v == 1 {
			convertAngka, _ := strconv.Atoi(k)

			result = append(result, convertAngka)
		}
	}

	return result
}

// func Mapping(slice []string) map[string]int {
// 	var resultMapping = map[string]int{}

// 	for _, keyMap := range slice {
// 		resultMapping[keyMap]++
// 	}

// 	return resultMapping
// }

// func ArrayMerge(arrayA []string, arrayB []string) []string {
// 	mergedArr := append(arrayA, arrayB...)
// 	uniqueArr := []string{}

// 	isFound := false

// 	for i, v := range mergedArr {
// 		if len(uniqueArr) == 0 {
// 			uniqueArr = append(uniqueArr, v)
// 		} else {
// 			isFound = false
// 			for j := i + 1; j < len(mergedArr); j++ {
// 				if v == mergedArr[j] {
// 					isFound = true
// 					// break
// 				}
// 			}

// 			if !isFound {
// 				uniqueArr = append(uniqueArr, v)
// 			}
// 		}
// 	}

// 	return uniqueArr
// }
