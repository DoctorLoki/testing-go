package main

// To run this program:
//  go build main.go; ./main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	timingTests()
}

// A timing test. Which is better?
// * Implementing set-inclusion by linear iteration over a slice of strings, or
// * Implementing set-inclusion by first construction a map of strings?
//
// Spoiler: For small numbers of strings below about 50 or sometimes up to 100,
// the linear iteration is faster. This is because linear searching avoids memory
// allocation (which always requires iterating over all the strings).
func timingTests() {
	fmt.Printf("Timing tests on arrays and maps.\n")

	for i := 10; i <= 100; i += 10 {
		fmt.Printf("Examine string slices of length %d.\n", i)
		timingTestLinearSearchVsMaps(i)
		fmt.Printf("-----\n")
	}
}

func timingTestLinearSearchVsMaps(numStrings int) {
	sliceB := makeShortStringSlice(numStrings, 2)
	sliceC := makeShortStringSlice(numStrings, 3)
	sliceD := makeShortStringSlice(numStrings, 4)
	sliceE := makeShortStringSlice(numStrings, 5)
	sliceF := makeShortStringSlice(numStrings, 6)

	for i := 3; i <= 6; i++ {
		n := int(math.Pow(10, float64(i)))
		timeLookups(n, "B,C", sliceB, sliceC)
		timeLookups(n, "B,D", sliceB, sliceD)
		timeLookups(n, "B,E", sliceB, sliceE)
		timeLookups(n, "B,F", sliceB, sliceF)
		timeLookups(n, "C,D", sliceC, sliceD)
		timeLookups(n, "C,E", sliceC, sliceE)
		timeLookups(n, "C,F", sliceC, sliceF)
	}
}

func makeShortStringSlice(numStrings int, moduloSkip int) []string {
	var slice []string
	for i := 0; i < numStrings; i++ {
		unique := 1000 + i // Produces a short unique string.
		if unique % moduloSkip == 0 {
			continue // Skip certain strings deterministically to make a variety of test cases.
		}
		slice = append(slice, fmt.Sprintf("%d", unique))
	}
	return slice
}

func timeLookups(n int, testname string, slice1 []string, slice2 []string) {
	fmt.Printf("iterations: %d\ttest: %s\t", n, testname)
	elapsed1 := linearLookups(n, slice1, slice2)
	fmt.Printf("linear: %s\t", elapsed1)
	elapsed2 := strmapLookups(n, slice1, slice2)
	fmt.Printf("strmap: %s\t", elapsed2)
	fmt.Printf("linear < strmap: %v\n", elapsed1 < elapsed2)
}

func linearLookups(n int, slice1 []string, slice2 []string) time.Duration {
	start := time.Now()

	for i := 0; i < n; i++ {
		StringSliceInStringSlice(slice1, slice2)
	}

	return time.Since(start)
}

func strmapLookups(n int, slice1 []string, slice2 []string) time.Duration {
	start := time.Now()

	for i := 0; i < n; i++ {
		StringSliceInStringSliceUsingMap(slice1, slice2)
	}

	return time.Since(start)
}

// StringInStringSlice returns true iff the string is within the slice.
// This is implemented as an O(n) linear search through the given slice.
func StringInStringSlice(s string, slice []string) bool {
	for _, s2 := range slice {
		if s == s2 {
			return true
		}
	}
	return false
}

// StringSliceInStringSlice returns true iff every string within slice1
// occurs within slice2.
// This is implemented as a linear search using StringInStringSlice.
// Accordingly, it is O(n^2) in time complexity and allocated no memory.
func StringSliceInStringSlice(slice1 []string, slice2 []string) bool {
	for _, s := range slice1 {
		if !StringInStringSlice(s, slice2) {
			return false
		}
	}
	return true
}

// StringInStringSliceUsingMap returns true iff the string is within the slice.
// This is implemented as an O(n) map construction step followed by an O(1) lookup.
// It uses O(n) memory due to the memory allocation requirements.
func StringInStringSliceUsingMap(s string, slice []string) bool {
	stringmap := make(map[string]struct{})
	for _, s2 := range slice {
		stringmap[s2] = struct{}{}
	}
	if _, ok := stringmap[s]; ok {
		return true
	}
	return false
}

// StringSliceInStringSliceUsingMap returns true iff every string within slice1
// occurs within slice2.
// This is implemented as a O(n) map construction followed by N O(1) lookups.
// Accordingly, it is O(n) time complexity and uses O(n) space.
func StringSliceInStringSliceUsingMap(slice1 []string, slice2 []string) bool {
	stringmap := make(map[string]struct{})
	for _, s2 := range slice2 {
		stringmap[s2] = struct{}{}
	}
	for _, s1 := range slice1 {
		if _, ok := stringmap[s1]; !ok {
			return false
		}
	}
	return true
}


