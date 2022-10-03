package main

import "fmt"

func area(m [][]int, x, y int) int {
	var max int
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if m[x][y] == 1 {
				cur := fill(m, x, y)
				if max < cur {
					max = cur
				}
			}

		}
	}
	return max
}
func fill(m [][]int, x, y int) int {
	if 0 <= x && x < len(m) && 0 <= y && y < len(m[0]) && m[x][y] == 1 {
		m[x][y] = 0
		return 1 + fill(m, x-1, y-1) +
			fill(m, x, y-1) +
			fill(m, x+1, y-1) +
			fill(m, x+1, y) +
			fill(m, x+1, y+1) +
			fill(m, x, y+1) +
			fill(m, x-1, y+1) +
			fill(m, x-1, y)
	}
	return 0
}
func main() {
	m := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(area(m, 0, 0))
}
