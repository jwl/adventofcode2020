package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwl/adventofcode2020/aocutils"
)

func getIDs(rawIDs string) []int64 {
	stringIDs := strings.Split(rawIDs, ",")
	intIDs := []int64{}

	for _, e := range stringIDs {
		if e != "x" {
			id, _ := strconv.ParseInt(e, 10, 64)
			intIDs = append(intIDs, id)
		}
	}

	return intIDs
}

func getIDsAndOrder(rawIDs string) []int64 {
	stringIDs := strings.Split(rawIDs, ",")
	intIDs := []int64{}

	for _, e := range stringIDs {
		if e != "x" {
			id, _ := strconv.ParseInt(e, 10, 64)
			intIDs = append(intIDs, id)
		} else {
			intIDs = append(intIDs, 1)
		}
	}

	return intIDs
}

func getNextBus(earliestTimestamp int64, busID int64) int64 {
	return ((earliestTimestamp / busID) + 1) * busID
}

func checkTimestamp(t int64, ids []int64) bool {
	// check if timestamp t divides cleanly into
	// each id plus its individual offset
	for offset, id := range ids {
		if (t+int64(offset))%id != 0 {
			return false
		}
	}
	return true
}

func getMax(ids []int64) int64 {
	max := int64(1)
	for _, id := range ids {
		max = max * id
	}
	return max
}

func partTwoSlow(ids []int64) int64 {
	fmt.Printf("in partTwoSlow, ids: %#v\n", ids)
	i := int64(1)
	max := getMax(ids)
	t := ids[0]
	fmt.Printf("first timestamp to check is: %d, max is %d, i is %d\n", t, max, i)
	for t < max {
		t = ids[0] * i
		if checkTimestamp(t, ids) {
			fmt.Printf("Solution found, t=%d, i=%d\n", t, i)
			return t
		}
		i++
	}

	return t
}

func partTwoFast(ids map[int]int) int {
	minValue := 0
	runningProduct := 1
	fmt.Printf("ids: %#v\n", ids)
	for k, v := range ids {
		for (minValue+v)%k != 0 {
			minValue += runningProduct
		}
		runningProduct *= k
	}
	fmt.Println(minValue)
	return minValue
}

func main() {
	fmt.Println("day13-01 started")

	input := aocutils.LoadInputIntoListOfStrings("input")

	// partTwoSlow
	// ids := getIDsAndOrder(input[1])
	// fmt.Printf("Earliest timestamp t using partTwoSlow is: %d\n", partTwoSlow(ids))

	// partTwoFast
	lines := strings.Split(input[1], ",")
	lineOffsetMap := make(map[int]int)
	for i, l := range lines {
		if l == "x" {
			continue
		}
		fmt.Printf("i is %d, l is %s\n", i, l)
		lineNo, _ := strconv.Atoi(l)
		lineOffsetMap[lineNo] = i
	}
	fmt.Printf("lineOffsetMap is: %#v\n", lineOffsetMap)
	fmt.Printf("Earliest timestamp t using partTwoFast is: %d\n", partTwoFast(lineOffsetMap))
}
