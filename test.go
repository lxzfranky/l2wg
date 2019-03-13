package main

import (
	"fmt"
	"math"
)

func main() {

	var radius float64 = 10

	var area = math.Pow(radius, 2) * math.Pi

	fmt.Println(area)
}
