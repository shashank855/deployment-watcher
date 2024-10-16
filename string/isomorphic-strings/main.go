package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		var isisomorphic bool
		map_for_str1 := make(map[string]int)
		map_for_str2 := make(map[string]int)
		str1 := []byte(s)
		str2 := []byte(t)

		for i := 0; i <= len(str1)-1; i++ {
			if map_for_str1[string(str1[i])] == 0 {
				map_for_str1[string(str1[i])] = i + 1
			}
			if map_for_str2[string(str2[i])] == 0 {
				map_for_str2[string(str2[i])] = i + 1
			}
		}
		fmt.Println(map_for_str1)
		fmt.Println(map_for_str2)

		for i := 0; i <= len(str1)-1; i++ {
			if map_for_str1[string(str1[i])] != map_for_str2[string(str2[i])] {
				isisomorphic = false
				break
			} else {
				isisomorphic = true
			}
		}
		return isisomorphic
	}
}

func main() {
	s1 := "paper"
	s2 := "title"
	fmt.Println(isIsomorphic(s1, s2))
}
