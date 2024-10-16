package main

import (
	"fmt"
)

func checkSmallString(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		return str2
	} else {
		return str1
	}
}
func commonPrefix(str1 string, str2 string) string {
	var chArr1 = []byte(str1)
	var chArr2 = []byte(str2)
	var commonString string
	iterator := checkSmallString(str1, str2)

	for i := 0; i < len(iterator); i++ {
		if chArr1[i] == chArr2[i] {
			commonString = commonString + string(chArr1[i])
		} else {
			break
		}

	}
	if len(commonString) == 0 {
		return ""
	} else {
		return commonString
	}
}

func longestCommonPrefix(strs []string) string {
	finalPrefix := strs[0]
	for _, value := range strs {
		finalPrefix = commonPrefix(finalPrefix, value)
	}
	return finalPrefix
}

func main() {
	var num int
	var str string
	fmt.Println("enter size of array")
	fmt.Scanf("%d", &num)
	var arr = make([]string, num)
	for i := 0; i < num; i++ {
		fmt.Printf("Enter %d string", i)
		fmt.Scanf("%s", &str)
		arr[i] = str
	}
	fmt.Println(longestCommonPrefix(arr))
}
