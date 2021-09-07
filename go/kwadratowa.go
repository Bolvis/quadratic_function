package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := getFloat()
	b := getFloat()
	c := getFloat()

	positiveDelta, r1, r2, delta := kwadratowa(a, b, c)
	if positiveDelta {
		fmt.Printf("x1 = %f\n", r1)
		fmt.Printf("x2 = %f\n", r2)
	}
	fmt.Printf("Δ = %f\n", delta)
}

func kwadratowa(a, b, c float64) (bool, float64, float64, float64) {
	var r1 float64
	var r2 float64
	var positiveDelta bool

	delta := math.Pow(b, 2) - (4 * a * c)
	if delta >= 0 {
		r1 = ((-1 * b) - math.Sqrt(delta)) / (2 * a)
		r2 = ((-1 * b) + math.Sqrt(delta)) / (2 * a)
		positiveDelta = true
	}
	
	return positiveDelta, r1, r2, delta
}

func getFloat() float64 {
set:
	fmt.Print("Set a ->")
	var x string
	fmt.Scanln(&x)
	num, err := strconv.ParseFloat(x, 64)
	if err != nil || num == 0 {
		fmt.Println("Something's wrong, try again")
		goto set
	}
	fmt.Printf("You've set number to %f\n", num)
	return num
}
