package main

var prefix map[string]bool = make(map[string]bool)
var wordsSet map[string]bool = make(map[string]bool)
var found map[string]bool = make(map[string]bool)
var mark [12][12]int
var current []byte
var m, n int

func findWords(board [][]byte, words []string) []string {

	defer func() {
		prefix = make(map[string]bool)
		wordsSet = make(map[string]bool)
		found = make(map[string]bool)
		current = []byte{}
		m, n = 0, 0
	}()

	if len(board) == 0 {
		return []string{}
	}

	for _, w := range words {
		wordsSet[w] = true
		for i := 1; i < len(w); i++ {
			prefix[w[:i]] = true
		}
	}

	m = len(board)
	n = len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			searchBoard(board, i, j)
		}
	}

	var result []string
	for k, v := range found {
		if v {
			result = append(result, k)
		}
	}
	return result
}

func searchBoard(board [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= m || y >= n || mark[x][y] == 1 {
		return
	}

	current = append(current, board[x][y])
	defer func() { current = current[:len(current)-1] }()
	if wordsSet[string(current)] {
		found[string(current)] = true
	}
	if len(current) > 10 || !prefix[string(current)] {
		return
	}

	mark[x][y] = 1
	searchBoard(board, x+1, y)
	searchBoard(board, x-1, y)
	searchBoard(board, x, y+1)
	searchBoard(board, x, y-1)
	mark[x][y] = 0
}
