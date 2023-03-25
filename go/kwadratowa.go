package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := getFloat(false, "a")
	b := getFloat(true, "b")
	c := getFloat(true, "c")

	positiveDelta, x1, x2, delta := kwadratowa(a, b, c)
	if positiveDelta {
		fmt.Printf("x1 = %f\n", x1)
		fmt.Printf("x2 = %f\n", x2)
	}
	fmt.Printf("Δ = %f\n", delta)
}

func kwadratowa(a, b, c float64) (bool, float64, float64, float64) {
	var x1 float64
	var x2 float64

	delta := math.Pow(b, 2) - (4 * a * c)

	positiveDelta := delta >= 0
	if positiveDelta {
		x1 = ((-1 * b) - math.Sqrt(delta)) / (2 * a)
		x2 = ((-1 * b) + math.Sqrt(delta)) / (2 * a)
	}

	return positiveDelta, x1, x2, delta
}

func getFloat(canBeZero bool, argumentName string) float64 {
	fmt.Printf("Set %s ->", argumentName)
	var x string
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Input reading error")
		return getFloat(canBeZero, argumentName)
	}

	num, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Cannot parse input to floating point number")
		return getFloat(canBeZero, argumentName)
	}
	if !canBeZero && num == 0 {
		fmt.Println("Passed number is 0 which is illegal argument")
		return getFloat(canBeZero, argumentName)
	}

	fmt.Printf("You've set %s to %f\n", argumentName, num)
	return num
}
