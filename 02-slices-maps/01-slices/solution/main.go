package main

import (
	"fmt"
	"sort"
)

// Chess piece symbols
const (
	whiteKing   = "♔"
	whiteQueen  = "♕"
	whiteRook   = "♖"
	whiteBishop = "♗"
	whiteKnight = "♘"
	whitePawn   = "♙"
	blackKing   = "♚"
	blackQueen  = "♛"
	blackRook   = "♜"
	blackBishop = "♝"
	blackKnight = "♞"
	blackPawn   = "♟"
	empty       = " "
)

// Function to create chess board
func createChessBoard() [][]string {
	board := make([][]string, 8)
	for i := range board {
		board[i] = make([]string, 8)
	}

	// Initialize black pieces (top)
	board[0] = []string{blackRook, blackKnight, blackBishop, blackQueen, blackKing, blackBishop, blackKnight, blackRook}
	board[1] = []string{blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn}

	// Initialize white pieces (bottom)
	board[6] = []string{whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn}
	board[7] = []string{whiteRook, whiteKnight, whiteBishop, whiteQueen, whiteKing, whiteBishop, whiteKnight, whiteRook}

	// Empty squares
	for i := 2; i < 6; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = empty
		}
	}

	return board
}

// Function to print chess board
func printChessBoard(board [][]string) {
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        CHESS BOARD                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                              ║")
	fmt.Println("║     a   b   c   d   e   f   g   h                           ║")
	fmt.Println("║   ┌───┬───┬───┬───┬───┬───┬───┬───┐                         ║")

	for i := 0; i < 8; i++ {
		fmt.Printf("║ %d │", 8-i)
		for j := 0; j < 8; j++ {
			fmt.Printf(" %s │", board[i][j])
		}
		fmt.Printf(" %d                       ║\n", 8-i)
		if i < 7 {
			fmt.Println("║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║")
		}
	}

	fmt.Println("║   └───┴───┴───┴───┴───┴───┴───┴───┘                         ║")
	fmt.Println("║     a   b   c   d   e   f   g   h                           ║")
	fmt.Println("║                                                              ║")
}

// Function to create tic-tac-toe board
func createTicTacToeBoard() [][]string {
	board := make([][]string, 3)
	for i := range board {
		board[i] = make([]string, 3)
		for j := range board[i] {
			board[i][j] = " "
		}
	}
	return board
}

// Function to place move on tic-tac-toe board
func placeMove(board [][]string, row, col int, player string) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != " " {
		return false
	}
	board[row][col] = player
	return true
}

// Function to check win condition
func checkWin(board [][]string, player string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

// Function to print tic-tac-toe board
func printTicTacToeBoard(board [][]string) {
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║                        TIC-TAC-TOE                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                              ║")
	fmt.Println("║        1   2   3                                             ║")
	fmt.Println("║      ┌───┬───┬───┐                                           ║")

	labels := []string{"A", "B", "C"}
	for i := 0; i < 3; i++ {
		fmt.Printf("║    %s │", labels[i])
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s │", board[i][j])
		}
		fmt.Println("                                           ║")
		if i < 2 {
			fmt.Println("║      ├───┼───┼───┤                                           ║")
		}
	}

	fmt.Println("║      └───┴───┴───┘                                           ║")
	fmt.Println("║                                                              ║")
}

// Function to add row to board
func addRow(board [][]string, row []string) [][]string {
	return append(board, row)
}

// Function to add column to board
func addColumn(board [][]string, col []string) [][]string {
	if len(board) == 0 {
		return board
	}

	newBoard := make([][]string, len(board))
	for i := range board {
		newBoard[i] = append(board[i], col[i])
	}
	return newBoard
}

// Function to remove row from board
func removeRow(board [][]string, index int) [][]string {
	if index < 0 || index >= len(board) {
		return board
	}
	return append(board[:index], board[index+1:]...)
}

// Function to rotate board 90 degrees clockwise
func rotateBoard(board [][]string) [][]string {
	if len(board) == 0 {
		return board
	}

	rows := len(board)
	cols := len(board[0])
	rotated := make([][]string, cols)
	for i := range rotated {
		rotated[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = board[i][j]
		}
	}

	return rotated
}

// Function to flip board horizontally
func flipBoard(board [][]string) [][]string {
	flipped := make([][]string, len(board))
	for i := range board {
		flipped[i] = make([]string, len(board[i]))
		copy(flipped[i], board[i])
		// Reverse the row
		for j := 0; j < len(flipped[i])/2; j++ {
			flipped[i][j], flipped[i][len(flipped[i])-1-j] = flipped[i][len(flipped[i])-1-j], flipped[i][j]
		}
	}
	return flipped
}

// Function to print dynamic board demo
func printDynamicBoardDemo() {
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║                    DYNAMIC BOARD DEMO                        ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                              ║")

	// Original board
	originalBoard := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	fmt.Println("║  Original Board:                                             ║")
	printSimpleBoard(originalBoard)

	// After adding row
	newRow := []string{"A", "B", "C"}
	boardWithRow := addRow(originalBoard, newRow)
	fmt.Println("║                                                              ║")
	fmt.Println("║  After adding row:                                           ║")
	printSimpleBoard(boardWithRow)

	// After rotating
	rotatedBoard := rotateBoard(boardWithRow)
	fmt.Println("║                                                              ║")
	fmt.Println("║  After rotating 90°:                                         ║")
	printSimpleBoard(rotatedBoard)
}

// Function to print simple board
func printSimpleBoard(board [][]string) {
	fmt.Println("║  ┌───┬───┬───┐                                               ║")
	for i, row := range board {
		fmt.Print("║  │")
		for j, cell := range row {
			fmt.Printf(" %s │", cell)
			if j < len(row)-1 {
				fmt.Print("")
			}
		}
		fmt.Println("                                               ║")
		if i < len(board)-1 {
			fmt.Println("║  ├───┼───┼───┤                                               ║")
		}
	}
	fmt.Println("║  └───┴───┴───┘                                               ║")
}

func main() {
	fmt.Println("Exercise 3: Game Board Generator & Visualizer")
	fmt.Println("============================================")

	// Task 1: Create and print chess board
	chessBoard := createChessBoard()
	printChessBoard(chessBoard)

	// Task 2: Create and play tic-tac-toe
	tttBoard := createTicTacToeBoard()

	// Place some moves
	placeMove(tttBoard, 0, 0, "X")
	placeMove(tttBoard, 0, 1, "O")
	placeMove(tttBoard, 1, 1, "X")
	placeMove(tttBoard, 1, 0, "O")
	placeMove(tttBoard, 2, 2, "X")

	printTicTacToeBoard(tttBoard)

	// Check win condition
	if checkWin(tttBoard, "X") {
		fmt.Println("║  Status: X wins!                                             ║")
	} else if checkWin(tttBoard, "O") {
		fmt.Println("║  Status: O wins!                                             ║")
	} else {
		fmt.Println("║  Status: Game in progress                                    ║")
	}
	fmt.Println("║                                                              ║")

	// Task 3: Demonstrate dynamic board operations
	printDynamicBoardDemo()

	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	// Demonstrate slice operations
	fmt.Println("\n📊 Slice Operations Demo:")
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Printf("Original slice: %v\n", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Append operation
	numbers = append(numbers, 10)
	fmt.Printf("After appending 10: %v\n", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Slicing operations
	fmt.Printf("First 5 elements: %v\n", numbers[:5])
	fmt.Printf("Last 3 elements: %v\n", numbers[len(numbers)-3:])
	fmt.Printf("Elements 2-5: %v\n", numbers[2:6])

	// Sorting
	sort.Ints(numbers)
	fmt.Printf("Sorted slice: %v\n", numbers)

	// Filtering
	var evenNumbers, oddNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} else {
			oddNumbers = append(oddNumbers, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evenNumbers)
	fmt.Printf("Odd numbers: %v\n", oddNumbers)
}
