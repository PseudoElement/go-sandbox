package funcs

func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	visited := make(map[[2]int]struct{})

	for row := range len(board) {
		for col := 0; col < len(board[0]); col++ {
			if backtrack(board, row, col, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, row, col int, word string, wordIdx int, visited map[[2]int]struct{}) bool {
	// Found the entire word
	if wordIdx == len(word) {
		return true
	}

	// Bounds check
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}

	key := [2]int{row, col}
	_, isVisited := visited[key]
	// Already visited or character doesn't match
	if isVisited || board[row][col] != word[wordIdx] {
		return false
	}

	// Mark as visited
	visited[key] = struct{}{}

	// Explore all 4 directions
	found := backtrack(board, row+1, col, word, wordIdx+1, visited) ||
		backtrack(board, row-1, col, word, wordIdx+1, visited) ||
		backtrack(board, row, col+1, word, wordIdx+1, visited) ||
		backtrack(board, row, col-1, word, wordIdx+1, visited)

	// BACKTRACK: unmark the cell so other paths can use it
	delete(visited, key)

	return found
}
