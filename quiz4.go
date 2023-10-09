package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) EstimateDestination() float64 {
	var fuelEffeciency = 1.5 // 1.5 L/mil

	return car.FuelIn / fuelEffeciency
}

type Student struct {
	Name  []string
	Score []int
}

func (student Student) AverageScore() string {
	totalScore := 0

	for _, score := range student.Score {
		totalScore += score
	}

	averageScore := "Average Score : " + fmt.Sprintf("%d", totalScore/len(student.Score))

	return averageScore
}

func (student Student) MinMaxScore() (string, string) {
	minScoreName := student.Name[0]
	minScore := student.Score[0]
	maxScoreName := student.Name[0]
	maxScore := student.Score[0]

	for i := 1; i < len(student.Score); i++ {
		// check the minScore value is the small value than current student score
		if student.Score[i] < minScore {
			// if the minScore less then the next score then replace it
			minScoreName = student.Name[i]
			minScore = student.Score[i]
		}

		if student.Score[i] > maxScore {
			// if the
			maxScoreName = student.Name[i]
			maxScore = student.Score[i]
		}
	}

	resultMinScore := fmt.Sprintf("Min Score of Students: %s (%d)", minScoreName, minScore)
	resultMaxScore := fmt.Sprintf("\nMax Score of Students: %s (%d)", maxScoreName, maxScore)

	return resultMinScore, resultMaxScore
}

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for i := 0; i < len(numbers); i++ {
		// check min score
		if *numbers[i] < min {
			min = *numbers[i]
		}

		// check max score
		if *numbers[i] > max {
			max = *numbers[i]
		}
	}

	return min, max
}

func main() {
	// 1.
	fmt.Println("Jawaban No 1.")

	sedan := Car{"Sedan", 2}
	estimatedDistance := sedan.EstimateDestination()

	fmt.Printf("Mobil tipe %s dengan %0.2f liter bahan bakar bisa ditempuh sekitar %0.2f mill\n", sedan.Type, sedan.FuelIn, estimatedDistance)

	fmt.Println("")

	// 2.
	fmt.Println("Jawaban No 2.")

	studentScore := Student{}
	totalStudent := 5

	for i := 0; i < totalStudent; i++ {
		// var name string
		// var score int

		// fmt.Printf("Input student name %d: ", i+1)
		// fmt.Scanln(&name)
		// fmt.Printf("Input student score %d: ", i+1)
		// fmt.Scanln(&score)

		// studentScore.Name = append(studentScore.Name, name)
		// studentScore.Score = append(studentScore.Score, score)

		studentScore.Name = append(studentScore.Name, "")
		studentScore.Score = append(studentScore.Score, 0)

		fmt.Printf("Input student name %d: ", i+1)
		fmt.Scanln(&studentScore.Name[i])
		fmt.Printf("Input student score %d: ", i+1)
		fmt.Scanln(&studentScore.Score[i])

	}

	fmt.Println(studentScore.AverageScore())
	fmt.Println(studentScore.MinMaxScore())

	fmt.Println("")

	// 3.
	fmt.Println("Jawaban No 3.")

	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min", min)
	fmt.Println("Nilai max", max)
}
