package main

import (
	"fmt"
	"math"
)

func main() {
	for approx := 2.0; approx < 3; approx = math.Nextafter(approx, 3) {
		fmt.Printf("%v\n", approx)
	}
}
