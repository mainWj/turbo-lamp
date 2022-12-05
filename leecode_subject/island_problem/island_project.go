package island_problem
// 此专题用来写leecode上的岛屿问题

// 1. 简单，岛屿数量，使用dfs遍历
func numIslands(grid [][]byte) int {

	line := len(grid)
	col := len(grid[0])
	if line == 0 && col == 0 {
		return 0
	}
	var dfs func(int, int, int, int, [][]byte)
	dfs = func(x int, y int, line int, col int, grid [][]byte){
		if x<0 || x >= line || y <0 || y >= col || grid[x][y] != '1'{
			return
		}
		grid[x][y] = '2'
		dfs(x + 1, y, line, col, grid)
		dfs(x - 1, y, line, col, grid)
		dfs(x, y - 1, line, col, grid)
		dfs(x, y + 1, line, col, grid)
	}
	result := 0
	for i := 0; i < line; i ++ {
		for j := 0; j < col; j ++{
			if grid[i][j] == '1'{
				result ++
				dfs(i, j, line, col, grid)
			}
		}
	}
	return result
}

// 简单  岛屿的周长
func islandPerimeter(grid [][]int) int {
	line, col := len(grid), len(grid[0])
	if line == 0  && col == 0{
		return 0
	}
	var perimeter func(int, int, [][]int)
	result := 0
	perimeter = func(x,y int, grid [][]int){
		if x < 0 || x >= line || y < 0 || y >= col || grid[x][y] == 0{
			result ++
			return
		}
		if grid[x][y] == 2{
			return
		}
		grid[x][y] = 2
		perimeter(x + 1, y, grid)
		perimeter(x - 1, y, grid)
		perimeter(x, y - 1, grid)
		perimeter(x, y + 1, grid)
	}
	for i:=0;i<line;i++ {
		for j:=0;j<col;j++{
			if grid[i][j] == 1{
				perimeter(i,j,grid)
				return result
			}
		}
	}
	return result
}

// 最大面积的岛屿，  dfs，加一个变量用来保存遍历的岛屿大小，并比较
func maxAreaOfIsland(grid [][]int) int {
	line, col := len(grid), len(grid[0])
	if line == 0 && col == 0 {
		return 0
	}
	var dfs func(int, int, [][]int)
	result, tmpcount := 0, 0
	dfs = func(x int, y int, grid [][]int){
		if x < 0 || x >= line || y < 0 || y >= col || grid[x][y] != 1{
			return
		}
		tmpcount ++
		grid[x][y] = 0
		dfs(x - 1, y, grid)
		dfs(x + 1, y, grid)
		dfs(x, y - 1, grid)
		dfs(x, y + 1, grid)
	}
	for i := 0; i < line; i++{
		for j := 0; j < col; j++{
			if grid[i][j] == 1{
				tmpcount = 0
				dfs(i, j, grid)
				if tmpcount > result{
					result = tmpcount
				}
			}
		}
	}
	return result

}