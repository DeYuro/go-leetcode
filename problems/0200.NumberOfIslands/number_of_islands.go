package problem200

func numIslands(grid [][]byte) int {

	var counter int

	col := len(grid[0])
	row := len(grid)

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				mark(grid, r, c)
				counter++
			}
		}
	}

	return counter
}

func mark(grid [][]byte, row, col int)  {

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}

	grid[row][col] = '6'

	mark(grid, row, col+1)
	mark(grid, row, col-1)
	mark(grid, row+1, col)
	mark(grid, row-1, col)

}
