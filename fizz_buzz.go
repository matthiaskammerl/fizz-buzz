package main

import "fmt"

var values = make(map[int]string)

var start = 1
var end = 100

func main() {
	values[3] = "Fizz"
	values[5] = "Buzz"

	random(values)
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

func checkAndAdd(count int, value int, word string) string {
	if count%value == 0 {
		return word
	}
	return ""
}
