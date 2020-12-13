package aocutils

import (
	"bufio"
	"log"
	"os"
)

// DuplicateSlice duplicates and returns a slice of strings.
func DuplicateSlice(src []string) []string {
	tmp := make([]string, len(src))
	copy(tmp, src)
	return tmp
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// LoadInputIntoListOfStrings takes a filename, reads it
// and returns it as a slice of strings, one per line.
func LoadInputIntoListOfStrings(filename string) []string {
	input := []string{}

	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		input = append(input, snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return input
}

// AreSlicesEqual returns whether the two given slices are equivalent.
func AreSlicesEqual(p1 []string, p2 []string) bool {
	if len(p1) != len(p2) {
		return false
	}
	for y, row := range p1 {
		for x, c := range row {
			if string(c) != string(p2[y][x]) {
				return false
			}
		}
	}
	return true
}

// Contains returns whether slice s contains the int e.
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Remove returns a slice with the element at index i removed.
func Remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// GetLargestElement returns the value of the largest element in list.
func GetLargestElement(list []int) int {
	largest := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > largest {
			largest = list[i]
		}
	}
	return largest
}

// GetLargestElement64 returns the index and value of the largest element in list.
func GetLargestElement64(list []int64) int64 {
	largest := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > largest {
			largest = list[i]
		}
	}
	return largest
}
