# Exercise 3: Game Board Generator & Visualizer

## 🎯 **Objective**
Create a game board generator that demonstrates Go slice operations, manipulation, and the `sort` package.

## 🎨 **Goal: Print a Dynamic Game Board**

You will create a program that generates and displays various game boards (like chess, tic-tac-toe, or a custom game) using slice operations and visual formatting.

## 📋 **Requirements**
1. Create and initialize slices in different ways:
   - Using `make()`
   - Using literal syntax
   - Using `append()`
2. Perform slice operations:
   - Slicing (using `[:]` syntax)
   - Appending elements
   - Copying slices
   - Removing elements
3. Use the `sort` package to sort slices
4. Work with multi-dimensional slices
5. Use `len()` and `cap()` functions

## 🎯 **Specific Tasks**

### **Task 1: Chess Board Generator**
Create a function that generates an 8x8 chess board:
- Initialize board with pieces in starting positions
- Use 2D slices to represent the board
- Display pieces using Unicode chess symbols

### **Task 2: Tic-Tac-Toe Game Board**
Create a 3x3 game board with:
- Empty cells, X's, and O's
- Functions to place moves
- Win condition checking
- Board state visualization

### **Task 3: Dynamic Board Operations**
Implement functions for:
- `addRow(board [][]string)` - Add new row to board
- `addColumn(board [][]string)` - Add new column to board
- `removeRow(board [][]string, index int)` - Remove specific row
- `rotateBoard(board [][]string)` - Rotate board 90 degrees
- `flipBoard(board [][]string)` - Flip board horizontally/vertically

### **Task 4: Print Beautiful Game Boards**
Generate visually appealing game boards that look like this:

```
╔══════════════════════════════════════════════════════════════╗
║                        CHESS BOARD                           ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║     a   b   c   d   e   f   g   h                           ║
║   ┌───┬───┬───┬───┬───┬───┬───┬───┐                         ║
║ 8 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │ 8                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 7 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ 7                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 6 │   │   │   │   │   │   │   │   │ 6                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 5 │   │   │   │   │   │   │   │   │ 5                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 4 │   │   │   │   │   │   │   │   │ 4                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 3 │   │   │   │   │   │   │   │   │ 3                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 2 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ 2                       ║
║   ├───┼───┼───┼───┼───┼───┼───┼───┤                         ║
║ 1 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │ 1                       ║
║   └───┴───┴───┴───┴───┴───┴───┴───┘                         ║
║     a   b   c   d   e   f   g   h                           ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║                        TIC-TAC-TOE                           ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║        1   2   3                                             ║
║      ┌───┬───┬───┐                                           ║
║    A │ X │ O │   │                                           ║
║      ├───┼───┼───┤                                           ║
║    B │ O │ X │   │                                           ║
║      ├───┼───┼───┤                                           ║
║    C │   │   │ X │                                           ║
║      └───┴───┴───┘                                           ║
║                                                              ║
║  Status: X wins!                                             ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║                    DYNAMIC BOARD DEMO                        ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Original Board:                                             ║
║  ┌───┬───┬───┐                                               ║
║  │ 1 │ 2 │ 3 │                                               ║
║  ├───┼───┼───┤                                               ║
║  │ 4 │ 5 │ 6 │                                               ║
║  ├───┼───┼───┤                                               ║
║  │ 7 │ 8 │ 9 │                                               ║
║  └───┴───┴───┘                                               ║
║                                                              ║
║  After adding row:                                           ║
║  ┌───┬───┬───┐                                               ║
║  │ 1 │ 2 │ 3 │                                               ║
║  ├───┼───┼───┤                                               ║
║  │ 4 │ 5 │ 6 │                                               ║
║  ├───┼───┼───┤                                               ║
║  │ 7 │ 8 │ 9 │                                               ║
║  ├───┼───┼───┤                                               ║
║  │ A │ B │ C │                                               ║
║  └───┴───┴───┘                                               ║
║                                                              ║
║  After rotating 90°:                                         ║
║  ┌───┬───┬───┬───┐                                           ║
║  │ 3 │ 6 │ 9 │ C │                                           ║
║  ├───┼───┼───┼───┤                                           ║
║  │ 2 │ 5 │ 8 │ B │                                           ║
║  ├───┼───┼───┼───┤                                           ║
║  │ 1 │ 4 │ 7 │ A │                                           ║
║  └───┴───┴───┴───┘                                           ║
╚══════════════════════════════════════════════════════════════╝
```

## 🎯 **Expected Output**
Your program should produce:
- Chess board with proper piece placement
- Tic-tac-toe game board with moves
- Dynamic board operations demonstration
- Formatted board visualizations

## 💡 **Hints**
- Use `make([][]string, rows, cols)` for 2D slices
- Use Unicode chess symbols: ♔♕♖♗♘♙♚♛♜♝♞♟
- Use `copy()` to duplicate board states
- Use `append()` to add rows/columns
- Use `fmt.Printf` with width specifiers for alignment
- Use Unicode box characters for borders

## 🚀 **Bonus Challenges**
1. Add move validation for chess pieces
2. Create a game loop for tic-tac-toe
3. Add board themes (different colors/styles)
4. Implement board save/load functionality
5. Add animation effects for moves
6. Create a board editor mode

## 🎨 **Learning Outcomes**
By completing this exercise, you will:
- Master slice creation and manipulation
- Understand 2D slice operations
- Learn sorting and filtering techniques
- Create visual data structures
- Practice real-world game development concepts

This exercise transforms slice operations into an engaging visual game board system! 🎮
