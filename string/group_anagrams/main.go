package main

import (
	"fmt"
	"sort"
)

func sort_my_str(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func groupAnagrams(strs []string) {
	my_map := make(map[string][]string)
	for _, val := range strs {
		sorted_string := sort_my_str(val)
		my_map[sorted_string] = append(my_map[sorted_string], val)
	}
	fmt.Println(my_map)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)

}
