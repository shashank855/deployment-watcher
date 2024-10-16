package main

import "fmt"

const no_of_queen int = 6

func getPossibleIndex(placed_element []int) (bool, []int) {
	var possible_index []int
	index_map := make(map[int]bool)
	if len(placed_element) == 0 {
		for i := 0; i < no_of_queen; i++ {
			possible_index = append(possible_index, i)
		}
		return true, possible_index
	} else {
		for i := 0; i < len(placed_element); i++ {
			if i == len(placed_element)-1 {
				index_map[placed_element[i]] = true
				if placed_element[i]-1 >= 0 {
					index_map[placed_element[i]-1] = true
				}
				if placed_element[i]+1 <= no_of_queen {
					index_map[placed_element[i]+1] = true
				}
			} else {
				index_map[placed_element[i]] = true
			}

		}
		for i := 0; i < no_of_queen; i++ {
			if index_map[i] == false {
				possible_index = append(possible_index, i)
			}
		}
		if len(possible_index) == 0 {
			return false, possible_index
		} else {
			return true, possible_index
		}
	}
}

func findNqueen(placed_element []int) {
	//fmt.Println("Printing placed element")
	//fmt.Println(placed_element)
	if len(placed_element) == no_of_queen {
		fmt.Println(placed_element)
		return
	}
	ok, possibleIndex := getPossibleIndex(placed_element)

	//fmt.Println("Printing possible index")
	//fmt.Println(possibleIndex)
	//fmt.Println("============")

	if ok {
		for _, val := range possibleIndex {
			placed_element_tmp := append(placed_element, val)
			findNqueen(placed_element_tmp)
		}
	}
}

func main() {
	// it stores the column index of the queen
	var placed_element []int
	placed_element = append(placed_element, 2)
	placed_element = append(placed_element, 0)
	placed_element = append(placed_element, 3)
	findNqueen(placed_element)
}
