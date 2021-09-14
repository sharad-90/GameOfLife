package main

func main() {
	arr := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(arr)
}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	copyBoard := make([][]int, m)
	for i := range board {
		copyBoard[i] = make([]int, len(board[i]))
		copy(copyBoard[i], board[i])
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total := findNeighbours(copyBoard, i, j, m, n)
			if copyBoard[i][j] == 1 && (total < 2||  total > 3)  {
					board[i][j] = 0
			} else if copyBoard[i][j] == 0 && total == 3 {
					board[i][j] = 1
			}
		}
	}
}

func findNeighbours(board [][]int, x, y, m, n int) int {
	var total int
	for i := -1; i <= 1; i = i + 1 {
		for j := -1; j <= 1; j = j + 1 {
			if i == 0 && j == 0 {
				continue
			}
			if i+x >= 0 && i+x < m &&
				j+y >= 0 && j+y < n && board[i+x][j+y]==1 {
				total = total + 1
			}
		}
	}
	return total
}
