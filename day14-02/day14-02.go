package main

import (
	"fmt"

	"github.com/jwl/adventofcode2020/aocutils"
)

func main() {
	fmt.Println("dayXX-YY started")

	input := []int{1, 2, 3}

	fmt.Printf("input contains 3: %#v", aocutils.Contains(input, 3))
}
