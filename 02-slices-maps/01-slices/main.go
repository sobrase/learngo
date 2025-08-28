package main

import (
	"fmt"
)

func chessBoard() {
	chessBoard := [][]string{
		{ "♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
		{ "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
		{ "♟", "♟", "♟", "♟", "♟", "♟ ", "♟", "♟"},
		{ "♟", "♟", "♟", "j", "♟", "♟", "♟", "♟"},
		{ "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
		{ "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
		{ "♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"},
		
}
fmt.Println("  a   b   c   d   e   f   g   h ")
for n,_ := range chessBoard {
	fmt.Printf("%d ",n)
	for _,j := range chessBoard[n]{

			fmt.Printf("│ %-3s ",j)
	
	}
	fmt.Printf("| %3d \n",n)
}
fmt.Println("  a   b   c   d   e   f   g   h ")
}

func main() {
	fmt.Println("Exercise 3: Game Board Generator & Visualizer")
	fmt.Println("============================================")
chessBoard()
	// TODO: Task 1 - Create chess board generator
	// Create a function that generates an 8x8 chess board:
	// - Initialize board with pieces in starting positions
	// - Use 2D slices to represent the board
	// - Display pieces using Unicode chess symbols

	// TODO: Task 2 - Create tic-tac-toe game board
	// Create a 3x3 game board with:
	// - Empty cells, X's, and O's
	// - Functions to place moves
	// - Win condition checking
	// - Board state visualization

	// TODO: Task 3 - Implement dynamic board operations
	// Implement functions for:
	// - addRow(board [][]string) - Add new row to board
	// - addColumn(board [][]string) - Add new column to board
	// - removeRow(board [][]string, index int) - Remove specific row
	// - rotateBoard(board [][]string) - Rotate board 90 degrees
	// - flipBoard(board [][]string) - Flip board horizontally/vertically

	// TODO: Task 4 - Print beautiful game boards
	// Generate visually appealing game boards that show:
	// - Chess board with proper piece placement
	// - Tic-tac-toe game board with moves
	// - Dynamic board operations demonstration
	// - Formatted board visualizations

	fmt.Println("\n🎮 Your game boards will appear here!")
	fmt.Println("Complete the TODO tasks above to generate beautiful game boards.")
}
