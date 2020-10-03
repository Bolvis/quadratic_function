package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var possitiveDelta bool

func main() {
	possitiveDelta = false
seta:
	fmt.Print("Set a ->")
	var x string
	fmt.Scanln(&x)
	strings.Trim(x, "\n")
	a, err := strconv.ParseFloat(x, 64)
	if err != nil || a == 0 {
		fmt.Println("Something's wrong, try again")
		goto seta
	}
	fmt.Printf("You've set number to %f\n", a)
setb:
	fmt.Print("Set b ->")
	fmt.Scanln(&x)
	strings.Trim(x, "\n")
	b, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Something's wrong, try again")
		goto setb
	}
	fmt.Printf("You've set number to %f\n", b)
setc:
	fmt.Print("Set c ->")
	fmt.Scanln(&x)
	strings.Trim(x, "\n")
	c, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Something's wrong, try again")
		goto setc
	}
	fmt.Printf("You've set number to %f\n", c)

	r := kwadratowa(a, b, c)
	if possitiveDelta {
		fmt.Printf("x1 = %f\n", r[1])
		fmt.Printf("x2 = %f\n", r[2])
	}
	fmt.Printf("Δ = %f\n", r[0])
}

func kwadratowa(a, b, c float64) [3]float64 {
	var r [3]float64
	delta := math.Pow(b, 2) - (4 * a * c)
	r[0] = delta
	if delta >= 0 {
		r[1] = ((-1 * b) - math.Sqrt(delta)) / (2 * a)
		r[2] = ((-1 * b) + math.Sqrt(delta)) / (2 * a)
		possitiveDelta = true
	}
	return r
}
