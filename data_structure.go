package main

import "fmt"

func main() {
	// struktur data

	// Array
	// var scoreData [5]int = [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(scoreData); i++ {
	// 	fmt.Println(scoreData[i])
	// }

	// // for range
	// for i, v := range scoreData {
	// 	fmt.Println(i, v)
	// }

	// Slice (banyak data, index, dinamis)
	// scoreData := make([]int, 10)

	// var scoreData []int = []int{1, 2, 3}
	// scoreData = append(scoreData, 10)
	// scoreData[0] = 10

	// fmt.Println(scoreData)

	// Slices
	// scoreData := []int{1, 3, 4, 5, 6}
	// scoreData = slices.Delete(scoreData, 2, 4)
	// fmt.Println(scoreData)

	// array to slice
	// arrayData := [5]int{1, 2, 3, 4, 5}
	// sliceData := arrayData[:]
	// sliceData = append(sliceData, 10)

	// arrayData[0] = 100
	// fmt.Println(sliceData)

	// scoreData[0] = 11
	// scoreData[4] = 33

	// fmt.Println(scoreData)
	// fmt.Println(scoreData[3])

	// MAP
	// var dataPenjualan = map[string]int{}
	var dataPenjualan = make(map[string]int)

	dataPenjualan["januari"] = 10000
	dataPenjualan["februari"] = 20000

	// check key is valid or not
	_, isValid := dataPenjualan["aosidjaoisd"]

	fmt.Println(dataPenjualan, isValid)
}
