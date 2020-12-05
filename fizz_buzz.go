package main

import "fmt"

var values = make(map[int]string)

func main() {

	var start = 1
	var end = 100

	values[3] = "Fizz"
	values[5] = "Buzz"
	values[7] = "Bizz"

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
