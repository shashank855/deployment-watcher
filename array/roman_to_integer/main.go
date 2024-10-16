package main

import "fmt"

var Roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	byteArr := []byte(s)
	var number, adjustSum int
	for key, val := range byteArr {
		if val == byte('I') && key+1 < len(byteArr) {
			if byteArr[key+1] == byte('V') || byteArr[key+1] == byte('X') {
				adjustSum = adjustSum - 1
			} else {
				adjustSum = adjustSum + 1
			}
		} else if val == byte('X') && key+1 < len(byteArr) {
			if byteArr[key+1] == byte('L') || byteArr[key+1] == byte('C') {
				adjustSum = adjustSum - 10
			} else {
				adjustSum = adjustSum + 10
			}
		} else if val == byte('C') && key+1 < len(byteArr) {
			if byteArr[key+1] == byte('D') || byteArr[key+1] == byte('M') {
				adjustSum = adjustSum - 100
			} else {
				adjustSum = adjustSum + 100
			}
		} else {
			number = number + Roman[byte(val)]
		}
	}
	return number + adjustSum
}

func main() {
	fmt.Println(romanToInt("III"))
	// romanToInt("XIV")
}
