package main

import (
	"fmt"
	"sort"
)

var values = map[int]string{
	3: "Fizz",
	5: "Buzz",
}

var start = 1
var end = 100

func main() {

	sorted(values)
}

func random(values map[int]string) {
	/*
		Range over maps are returned random intentionally in go
		Order of word concatenation is therefore random as well with this method
	*/
	for i := start; i <= end; i++ {
		output := ""

		for value, word := range values {
			output += checkAndAdd(i, value, word)
		}

		if output == "" {
			output += fmt.Sprint(i)
		}

		fmt.Println(output)
	}
}

func sorted(values map[int]string) {
	/*
		Range over slices are returned sorted
		This way the order of word can be correct
	*/
	for i := start; i <= end; i++ {
		output := ""

		output = createString(i, values)

		fmt.Println(output)
	}

	fmt.Println(createString(3, values))
}

func checkAndAdd(count int, value int, word string) string {
	if count%value == 0 {
		return word
	}
	return ""
}

func createString(count int, words map[int]string) string {
	var i string

	var outputValues []int

	for value := range words {
		if count%value == 0 {
			outputValues = append(outputValues, value)
		}
	}

	sort.Ints(outputValues)

	for _, wordIndex := range outputValues {
		i += values[wordIndex]
	}

	if i == "" {
		i += fmt.Sprint(count)
	}

	return i

}
