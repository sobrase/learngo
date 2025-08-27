package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// Book struct for JSON/XML serialization
type Book struct {
	ID        int      `json:"id" xml:"id"`
	Title     string   `json:"title" xml:"title"`
	Author    string   `json:"author" xml:"author"`
	Year      int      `json:"year" xml:"year"`
	Price     float64  `json:"price" xml:"price"`
	Available bool     `json:"available" xml:"available"`
	Tags      []string `json:"tags" xml:"tags>tag"`
}

// Catalog struct for XML
type Catalog struct {
	Books       []Book    `xml:"book"`
	TotalBooks  int       `xml:"total_books"`
	LastUpdated time.Time `xml:"last_updated"`
}

// CustomBook with custom marshaling
type CustomBook struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func (cb CustomBook) MarshalJSON() ([]byte, error) {
	type Alias CustomBook
	return json.Marshal(&struct {
		Alias
		FormattedPrice string `json:"formatted_price"`
	}{
		Alias:          Alias(cb),
		FormattedPrice: fmt.Sprintf("USD %.2f", cb.Price),
	})
}

func main() {
	fmt.Println("Exercise 10: JSON and XML Serialization")
	fmt.Println("=======================================")

	// 1. Create book data
	books := []Book{
		{
			ID:        1,
			Title:     "The Go Programming Language",
			Author:    "Alan Donovan",
			Year:      2015,
			Price:     39.99,
			Available: true,
			Tags:      []string{"programming", "go", "computer-science"},
		},
		{
			ID:        2,
			Title:     "Effective Go",
			Author:    "Go Team",
			Year:      2020,
			Price:     0.00,
			Available: true,
			Tags:      []string{"programming", "go", "free"},
		},
	}

	// 2. JSON Serialization
	fmt.Println("JSON Serialization:")
	catalog := map[string]interface{}{
		"books":        books,
		"total_books":  len(books),
		"last_updated": time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println()

	// 3. XML Serialization
	fmt.Println("XML Serialization:")
	catalogXML := Catalog{
		Books:       books,
		TotalBooks:  len(books),
		LastUpdated: time.Now(),
	}

	xmlData, err := xml.MarshalIndent(catalogXML, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling XML: %v\n", err)
		return
	}
	fmt.Println(xml.Header + string(xmlData))
	fmt.Println()

	// 4. Custom Marshaling
	fmt.Println("Custom Marshaling:")
	customBook := CustomBook{
		ID:    3,
		Title: "Custom Book",
		Price: 25.50,
	}

	customJSON, err := json.MarshalIndent(customBook, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling custom JSON: %v\n", err)
		return
	}
	fmt.Println(string(customJSON))
	fmt.Println()

	// 5. JSON Unmarshaling
	fmt.Println("JSON Unmarshaling:")
	jsonStr := `{"id":4,"title":"New Book","author":"Unknown","year":2024,"price":29.99,"available":true,"tags":["new","fiction"]}`

	var newBook Book
	err = json.Unmarshal([]byte(jsonStr), &newBook)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	fmt.Printf("Unmarshaled book: %+v\n", newBook)
	fmt.Println()

	// 6. Streaming JSON
	fmt.Println("Streaming JSON:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	fmt.Println("Processing books...")
	for i, book := range books {
		fmt.Printf("Processing book %d...\n", i+1)
		encoder.Encode(book)
	}
	fmt.Println("All books processed successfully")
}
