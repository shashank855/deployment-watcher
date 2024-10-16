package main

import "fmt"

//Finding all subset of {"abc"}
// ex: input {abc} output: {"","a","b","c","ab","ac","bc","abc"}
func printSubset(mystr []string, output []string) {
	if len(mystr) > 0 {
		printSubset(mystr[1:], append(output, mystr[0]))
		printSubset(mystr[1:], output)

	} else {
		fmt.Println(output)
	}

}

func main() {
	var mystr = []string{"a", "b", "c"}
	printSubset(mystr, []string{})
}
