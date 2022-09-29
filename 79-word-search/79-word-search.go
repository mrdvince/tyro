func exist(board [][]byte, word string) bool {
    rows, columns := len(board), len(board[0])
    
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, columns)
    }
    for i := 0; i < rows; i++ {
        // ["A"]
        for j := 0; j < columns;  j++ {
            if dfs(board,visited, word, 0, i, j) {
                return true
            }
        }
    } 
    return false
}

func dfs(board [][]byte,visited [][]bool, word string, curr int, i int, j int) bool {
    // "A"
    if !visited[i][j] && board[i][j] == word[curr] {
        if curr == len(word) - 1 {
            return true
        }
        visited[i][j] = true
//     top
    if i-1 >=0 && dfs(board,visited, word, curr+1, i-1, j) {
        return true
    }
//     bottom
    if i+1 < len(board) && dfs(board, visited, word, curr+1, i+1, j) {
        return true
    }
//     right
    if j+1 < len(board[0]) && dfs(board, visited, word, curr+1, i, j+1) {
        return true
    }
//     left
    if j-1 >= 0 && dfs(board, visited, word, curr+1, i, j-1) {
        return true
    }
        visited[i][j] = false
    }
    
    return false
}