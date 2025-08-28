package main

import (
	"encoding/json"
	"fmt"
)

type Shapes interface{
	area() int
}

type Rectangle struct{
	lenghth int;
	width int;
}

type Square struct{
	width int;
}

type Person struct {
	Name string `json:"name"`;
	Age int `json:"age"`;
	Email string `json:"email,omitempty"`;
}

func test(s Shapes){
	fmt.Println("test")
}

func(r Rectangle) area() int{
	return r.lenghth*r.width
}
func(s Square) area() int{
	return s.width*s.width
}

func(p Person) info(){
	fmt.Printf("name: %s \nage: %d \nemail: %s \n",p.Name,p.Age,p.Email)
}
func(p *Person) changeName(){
	p.Name="changed"
}
func main() {
	fmt.Println("Exercise 5: Library Management System")
	fmt.Println("====================================")
	person:=Person{
		Name: "test",
		Age: 13,
		Email: "mdkljfqmlkj@fldmjf.com",

	}
	person.info()
	person.changeName()
	fmt.Println(person)
	rectangle:=Rectangle{
		width: 16,
		lenghth: 16,
	}
	test(rectangle)
	jsonperson,err:=json.Marshal(person)
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
	fmt.Println(string(jsonperson))
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
