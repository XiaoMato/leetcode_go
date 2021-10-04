package main

func isValidSudoku(board [][]byte) bool {
	for _, v := range board {
		if !isLineValid(v) {
			return false
		}
	}

	for i := range board {
		c := []byte{}
		for _, v := range board {
			c = append(c, v[i])
		}
		if !isLineValid(c) {
			return false
		}
	}

	for _, x := range []int{0, 3, 6} {
		for _, y := range []int{0, 3, 6} {
			square := [][]byte{}
			square = append(square, board[x][y:y+3])
			square = append(square, board[x+1][y:y+3])
			square = append(square, board[x+2][y:y+3])
			if !isSquareValid(square) {
				return false
			}
		}
	}
	return true
}

func isLineValid(line []byte) bool {
	has := make(map[byte]bool)
	for _, v := range line {
		if v != '.' && has[v] {
			return false
		}
		has[v] = true
	}
	return true
}

func isSquareValid(square [][]byte) bool {
	has := make(map[byte]bool)
	for _, row := range square {
		for _, v := range row {
			if v != '.' && has[v] {
				return false
			}
			has[v] = true
		}
	}
	return true
}

func main() {
	println(isValidSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}
