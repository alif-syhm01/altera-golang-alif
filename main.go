package main

import "fmt"

func main() {
	// determine the prime number
	fmt.Println(isPrime(5))
	fmt.Println(isPrime(12))

	fmt.Println()

	// determine the number of multiple seven
	fmt.Println(multipleSeven(7))
	fmt.Println(multipleSeven(49))
	fmt.Println(multipleSeven(0))

	fmt.Println()

	// find the area of the trapezoid in cm
	fmt.Println(areaOfTheTrapezoid(10, 8, 20))
	fmt.Println(areaOfTheTrapezoid(16, 10, 8))
}

func isPrime(number int) string {
	var counter uint8 = 0
	var result string

	for divider := 1; divider <= number; divider++ {
		if number%divider == 0 {
			counter++
		}
	}

	if counter == 2 {
		result = fmt.Sprintf("%d", number) + " adalah bilangan prima"
	} else {
		result = fmt.Sprintf("%d", number) + " bukan bilangan prima"
	}

	return result
}

func multipleSeven(number int) string {
	var result string

	if number%7 == 0 && number >= 1 {
		result = fmt.Sprintf("%d", number) + " adalah kelipatan 7"
	} else {
		result = fmt.Sprintf("%d", number) + " bukan kelipatan 7"
	}

	return result
}

func areaOfTheTrapezoid(a int, b int, height int) string {
	const trapezoidalGeometry float32 = 0.5
	var areaFormula float32 = trapezoidalGeometry * float32(height) * (float32(a) + float32(b))

	result := "Luas = " + fmt.Sprintf("%f", areaFormula) + "cm"

	return result
}

// func main() {
// 	// Loop
// 	for j := 1; j <= 5 ; j++ {
// 		// nested loop
// 		for i := 1; i <= j; i++ {
// 			fmt.Print("*")
// 			// fmt.Println("Hello World", i)
// 		}

// 		fmt.Println("")
// 	}

// 	// Manual
// 	// fmt.Println("Hello World")
// 	// fmt.Println("Hello World")
// 	// fmt.Println("Hello World")
// 	// fmt.Println("Hello World")
// 	// fmt.Println("Hello World")

// 	// declare
// 	// var age uint = 10

// 	// if age == 17 && age == 10 {
// 	// 	fmt.Println("Sweet seventeen")
// 	// } else if age > 5 {
// 	// 	fmt.Println("Kid")
// 	// } else {
// 	// 	fmt.Println("Invalid Age")
// 	// }

// 	// a := 7
// 	// b := 3

// 	// Operasi Bitwise >>, <<
// 	// result := 2 << 1
// 	// 2 = 10 (biner) > 0 * 2^0 + 1 * 2^1 = 0 + 2 = 2
// 	// ? = 100 > 0 * 2^0 + 0 * 2^1 + 1 * 2^2 = 0 + 0 + 4 = 4

// 	// fmt.Println(result)

// 	// // operasi logika &&, ||
// 	// // result := a == 7 && a > b // true && true = true
// 	// result := a == 5 && a > b // false && true = false

// 	// fmt.Println(result)

// 	// kondisi >, <, >=, <=, ==, !=, !
// 	// result := a >= b

// 	// fmt.Println(result)

// 	// result :Print(a)
// 	// result := a * b
// 	// result := a % b

// 	// fmt.Println(resufmt

// 	// var name = "Alta"

// 	// // short declaration
// 	// const phi float32 = 3.14

// 	// // Assignment
// 	// name = "Malang"

// 	// // showing the output
// 	// fmt.Println(age)
// }
