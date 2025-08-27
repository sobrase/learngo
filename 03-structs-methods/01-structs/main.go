package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise 5: Library Management System")
	fmt.Println("====================================")

	// TODO: Task 1 - Create book and author structs
	// Create structs for:
	// - Author with name, birthYear, nationality, genres
	// - Book with title, author, ISBN, year, pages, price, available
	// - Library with books, members, transactions
	// - Member with ID, name, email, membershipType, borrowedBooks

	// TODO: Task 2 - Implement library operations
	// Implement methods for:
	// - Book.DisplayInfo() - Show book details
	// - Book.IsAvailable() - Check availability
	// - Book.CalculateLateFee(days int) - Calculate late fees
	// - Library.AddBook(book Book) - Add book to library
	// - Library.FindBook(isbn string) - Search for book
	// - Library.CheckoutBook(isbn, memberID string) - Checkout process
	// - Library.ReturnBook(isbn, memberID string) - Return process

	// TODO: Task 3 - Create advanced features
	// Create methods for:
	// - Library.GetBooksByAuthor(authorName string) - Find author's books
	// - Library.GetOverdueBooks() - Find overdue books
	// - Library.GenerateReport() - Generate library statistics
	// - Member.GetBorrowingHistory() - Get member's history

	// TODO: Task 4 - Print library dashboard
	// Generate a beautiful dashboard that shows:
	// - Library statistics and catalog
	// - Member information and status
	// - Search results and recommendations
	// - Overdue book alerts
	// - Transaction history
	// - Financial summary

	fmt.Println("\n📚 Your library catalog dashboard will appear here!")
	fmt.Println("Complete the TODO tasks above to generate a beautiful library dashboard.")
}
