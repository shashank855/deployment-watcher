package main

import "fmt"

func traverseAndMark(row int, col int, grid *[][]byte) {
	if (*grid)[row][col] == 0 || (*grid)[row][col] == 2 {
		return
	} else {
		// When row or column is zero increment row also and column also
		if row == 0 || col == 0 {
			(*grid)[row][col] = 2 //Setting 2 for visited item
			traverseAndMark(row, (col+1)%(len((*grid)[0])-1), grid)
			traverseAndMark((row+1)%(len((*grid)[0])-1), col, grid)
		}
		// if col == 0 {
		// 	(*grid)[row][col] = 2
		// 	traverseAndMark((row+1)%(len((*grid)[0])-1), col, grid)
		// }

		// When row and column both are agreater than zero do search in all four direction
		if row > 0 && col > 0 {
			(*grid)[row][col] = 2
			traverseAndMark((row+1)%(len((*grid)[0])-1), col, grid)
			traverseAndMark((row-1)%(len((*grid)[0])-1), col, grid)
			traverseAndMark((col+1)%(len((*grid)[0])-1), col, grid)
			traverseAndMark((col-1)%(len((*grid)[0])-1), col, grid)
		}
	}

}

func numIslands(grid [][]byte) int {
	var num_of_islands int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				num_of_islands += 1
				traverseAndMark(i, j, &grid)
				// fmt.Println("row:", i, "col:", j)
				// fmt.Println(grid)
			}
		}
	}
	return num_of_islands
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	islands := numIslands(grid)
	fmt.Println(islands)

}
