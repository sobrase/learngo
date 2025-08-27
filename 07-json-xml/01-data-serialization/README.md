# Exercise 10: JSON and XML Serialization

## Objective
Learn about Go's `encoding/json` and `encoding/xml` packages for data serialization.

## Problem
Create a program that demonstrates JSON and XML marshaling/unmarshaling with structs.

## Requirements
1. Use `encoding/json` package for JSON operations:
   - `json.Marshal()` and `json.Unmarshal()`
   - `json.NewEncoder()` and `json.NewDecoder()`
   - Struct tags for JSON field mapping
2. Use `encoding/xml` package for XML operations:
   - `xml.Marshal()` and `xml.Unmarshal()`
   - XML struct tags
3. Handle custom marshaling/unmarshaling
4. Work with nested structures and arrays
5. Handle different data types and conversions

## Tasks
1. Create structs for a book catalog system
2. Implement JSON serialization/deserialization
3. Implement XML serialization/deserialization
4. Create custom marshaling for special cases
5. Handle different data formats and conversions

## Expected Output
```
Book Catalog System:

JSON Serialization:
{
  "books": [
    {
      "id": 1,
      "title": "The Go Programming Language",
      "author": "Alan Donovan",
      "year": 2015,
      "price": 39.99,
      "available": true,
      "tags": ["programming", "go", "computer-science"]
    },
    {
      "id": 2,
      "title": "Effective Go",
      "author": "Go Team",
      "year": 2020,
      "price": 0.00,
      "available": true,
      "tags": ["programming", "go", "free"]
    }
  ],
  "total_books": 2,
  "last_updated": "2024-01-01T12:00:00Z"
}

XML Serialization:
<?xml version="1.0" encoding="UTF-8"?>
<catalog>
  <book>
    <id>1</id>
    <title>The Go Programming Language</title>
    <author>Alan Donovan</author>
    <year>2015</year>
    <price>39.99</price>
    <available>true</available>
    <tags>
      <tag>programming</tag>
      <tag>go</tag>
      <tag>computer-science</tag>
    </tags>
  </book>
  <book>
    <id>2</id>
    <title>Effective Go</title>
    <author>Go Team</author>
    <year>2020</year>
    <price>0.00</price>
    <available>true</available>
    <tags>
      <tag>programming</tag>
      <tag>go</tag>
      <tag>free</tag>
    </tags>
  </book>
  <total_books>2</total_books>
  <last_updated>2024-01-01T12:00:00Z</last_updated>
</catalog>

Custom Marshaling:
Book with custom price formatting:
{
  "id": 3,
  "title": "Custom Book",
  "price": "$25.50",
  "formatted_price": "USD 25.50"
}

Streaming JSON:
Processing book 1...
Processing book 2...
Processing book 3...
All books processed successfully
```

## Hints
- Use `json:"field_name"` tags for JSON field mapping
- Use `xml:"field_name"` tags for XML field mapping
- Use `json:"-"` to omit fields from JSON output
- Use `omitempty` option to skip empty fields
- Implement `json.Marshaler` and `json.Unmarshaler` for custom behavior
- Use `encoding/json` for streaming operations
- Remember to handle errors from marshaling/unmarshaling
- Use `time.Time` for date/time fields
