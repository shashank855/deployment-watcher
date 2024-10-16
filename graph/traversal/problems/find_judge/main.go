package main

import "fmt"

func contains(arr []int, judge_name int) bool {
	for _, val := range arr {
		if val == judge_name {
			return true
		}
	}
	return false
}

func check_trust(my_map map[int][]int, judge_candidates []int) int {
	final_judge := -1
	for _, judge_name := range judge_candidates {
	inner:
		for key, mapVal := range my_map {
			if key != judge_name {
				// fmt.Println("key:", key)
				// fmt.Println("judge_name:", judge_name)
				// fmt.Println("map:", mapVal)
				if contains(mapVal, judge_name) == false {
					judge_name = -1
					break inner
				}
			}
		}
		final_judge = judge_name

	}
	return final_judge
}

func find_judge(my_map map[int][]int) int {
	null_count := 0
	var judge_candidates []int
	for key, _ := range my_map {
		if len(my_map[key]) == 0 {
			judge_candidates = append(judge_candidates, key)
			null_count = null_count + 1
		}
	}
	fmt.Println("judge_candidates:", judge_candidates)
	if null_count >= 2 {
		return -1
	}
	return check_trust(my_map, judge_candidates)
}

func main() {
	//[[1,3],[2,3],[3,1]]
	my_arr := [][]int{
		// {1, 2},
		// {2, 3},
	}
	my_map := make(map[int][]int)
	for _, val := range my_arr {
		my_map[val[0]] = append(my_map[val[0]], val[1])
		if len(my_map[val[1]]) == 0 {
			my_map[val[1]] = append(my_map[val[1]])
		}
	}
	fmt.Println(len(my_map))
	fmt.Println(find_judge(my_map))

}
