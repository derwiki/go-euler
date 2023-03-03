package main

import "fmt"

func main() {
	/*
		var grid [1001][1001]int

		y := 500
		x := 500
	*/
	var grid [5][5]int

	y := 2
	x := 2
	i := 1
	grid[y][x] = i

	// move right and go down as fas as you can
	grid[y][x+1] = 2
	grid[y+1][x+1] = 3

	// now on bottom row, go left as far as you can
	grid[y+1][x] = 4
	grid[y+1][x-1] = 5

	// when you're as far left, start going up
	grid[y][x-1] = 6
	grid[y-1][x-1] = 7

	// now that we're on top, go right as far as you can
	grid[y-1][x] = 8
	grid[y-1][x+1] = 9

	fmt.Println(grid)
	for _, row := range grid {
		fmt.Println(row)
	}
}
