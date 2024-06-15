package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

var name, greeting string

func main() {
	// slide 10
	fmt.Println(quote.Go())

	// Slide 12
	name = "Beege"
	greeting = "Salutations"
	statement := greeting + ", " + name + "!"
	fmt.Println(statement)
	const a rune = 'a'
	stringA := string(a)
	fmt.Println("a:", a)
	fmt.Println("stringA:", stringA)

	// Slide 13
	// Function defs below first
	addResult := add(2, 3)
	fmt.Println("addResult:", addResult)
	addMultiplySum, addMultiplyProduct := addMultiply(4, 5)
	fmt.Println("addMultiply Sum:", addMultiplySum, "--- Product:", addMultiplyProduct)

	// Slide 14
	x := 0
	for x < 3 {
		fmt.Println("x:", x)
		x++
	}
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	// Slide 14 but function below first
	sumResult, averageResult := sumAverage(3, 4, 5)
	fmt.Println("sumAverage sumResult:", sumResult, "--- averageResult:", averageResult)

	// Slide 15
	if sumResult < 5 {
		fmt.Println("Meh sum")
	} else if sumResult < 15 {
		fmt.Println("Bigger sum")
	} else {
		fmt.Println("Big boi sum")
	}

	if tempSum := add(2, 3); tempSum > 4 {
		fmt.Println("tempSum > 4")
	}
	// the below does not exist in this scope
	//fmt.Println(tempSum)

	// Slide 16
	someNums := []int{1, 2, 3, 4}
	fmt.Println("some nums:", someNums)
	fmt.Println("len:", len(someNums), "-- cap:", cap(someNums))
	someNums = someNums[:2]
	fmt.Println("half some nums:", someNums)
	fmt.Println("len:", len(someNums), "-- cap:", cap(someNums))
	someNums = someNums[1:]
	fmt.Println("reduced capacity -- len:", len(someNums), "-- cap:", cap(someNums))
	someNums = make([]int, 0, 3)
	fmt.Println("make slice:", someNums, "-- len:", len(someNums), "-- cap:", cap(someNums))
	someNums = append(someNums, 20, 10, 5)
	fmt.Println("appended slice:", someNums)
	for index, value := range someNums {
		fmt.Println("index:", index, "-- value:", value)
	}

	// Slide 17
	// map[key type]value type
	ages := make(map[string]uint8)
	ages["Beege"] = 21
	fmt.Println("ages:", ages)
	ages = map[string]uint8{
		"Beege": 21,
		"Susan": 21,
	}
	fmt.Println("ages again:", ages)
	if ageVal, exists := ages["Gildong"]; exists {
		fmt.Println("ageVal:", ageVal)
	} else {
		fmt.Println("Gildong not in ages")
	}

	// Slide 18
	defer fmt.Println("deferred")
	for i := 0; i < 5; i++ {
		fmt.Println("not deferred")
	}

	// Slide 19
	integerValue, err := strconv.Atoi("-4")
	fmt.Println("value:", integerValue, "--- err:", err)
	integerValue, err = strconv.Atoi("a")
	fmt.Println("errored value:", integerValue, "--- err:", err)

	// Slide 19 but will panic
	// zero := 0
	// val := 7 / zero
	// fmt.Println(val)
}

// Slide 13
func add(a, b int) int {
	return a + b
}

func addMultiply(a, b int) (int, int) {
	return a + b, a * b
}

// Slide 14
func sumAverage(nums ...int) (sum, average int) {
	for _, num := range nums {
		sum += num
	}
	average = sum / len(nums)
	return
}
