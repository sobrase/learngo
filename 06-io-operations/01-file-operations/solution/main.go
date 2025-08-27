package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Exercise 11: File I/O Operations")
	fmt.Println("===============================")

	// 1. Create and write to file
	fmt.Println("Creating and writing to file:")
	content := "Hello, World!\nThis is a sample file.\nCreated with Go file operations.\n\nEnd of file.\n"

	err := os.WriteFile("sample.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("File 'sample.txt' created successfully")
	fmt.Printf("Content written: %d bytes\n\n", len(content))

	// 2. Read file content
	fmt.Println("Reading file content:")
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("Line %d: %s\n", i+1, line)
		}
	}
	fmt.Println()

	// 3. File statistics
	fmt.Println("File statistics:")
	fileInfo, err := os.Stat("sample.txt")
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %s\n", fileInfo.Mode())
	fmt.Printf("Is directory: %t\n\n", fileInfo.IsDir())

	// 4. Buffered reading
	fmt.Println("Buffered reading:")
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fmt.Printf("Processing line: %s\n", line)
		}
		lineNum++
	}
	fmt.Println()

	// 5. CSV Processing
	fmt.Println("CSV Processing:")
	csvData := `Name,Age,City
John,25,New York
Jane,30,Los Angeles
Bob,35,Chicago`

	err = os.WriteFile("data.csv", []byte(csvData), 0644)
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		return
	}

	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Printf("Error opening CSV: %v\n", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	fmt.Println("Parsed CSV data:")
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		fmt.Printf("%s is %s years old from %s\n", record[0], record[1], record[2])
	}
	fmt.Println()

	// 6. File copying
	fmt.Println("File copying:")
	sourceFile, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("Error opening source: %v\n", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("sample_copy.txt")
	if err != nil {
		fmt.Printf("Error creating destination: %v\n", err)
		return
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}

	fmt.Printf("Source file: sample.txt (%d bytes)\n", bytesCopied)
	fmt.Printf("Destination file: sample_copy.txt (%d bytes)\n", bytesCopied)
	fmt.Println("Copy completed successfully")
	fmt.Println()

	// 7. Directory operations
	fmt.Println("Directory operations:")

	// Create directory
	err = os.MkdirAll("test_dir/subdir", 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Println("Created directory: test_dir")
	fmt.Println("Created subdirectory: test_dir/subdir")

	// List directory contents
	entries, err := os.ReadDir("test_dir")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	fmt.Println("Listing test_dir contents:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  %s/ (directory)\n", entry.Name())
		} else {
			info, _ := entry.Info()
			fmt.Printf("  %s (%d bytes)\n", entry.Name(), info.Size())
		}
	}
	fmt.Println()

	// 8. File path operations
	fmt.Println("File path operations:")
	path := filepath.Join("test_dir", "subdir", "file.txt")
	fmt.Printf("Joined path: %s\n", path)

	dir := filepath.Dir(path)
	fmt.Printf("Directory: %s\n", dir)

	base := filepath.Base(path)
	fmt.Printf("Base name: %s\n", base)

	ext := filepath.Ext(path)
	fmt.Printf("Extension: %s\n", ext)
	fmt.Println()

	// 9. Cleanup
	fmt.Println("Cleanup:")
	filesToRemove := []string{
		"sample.txt",
		"sample_copy.txt",
		"data.csv",
	}

	for _, file := range filesToRemove {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", file, err)
		} else {
			fmt.Printf("Removed: %s\n", file)
		}
	}

	err = os.RemoveAll("test_dir")
	if err != nil {
		fmt.Printf("Error removing test_dir: %v\n", err)
	} else {
		fmt.Println("Removed test_dir directory")
	}
	fmt.Println("Removed test files and directories")
}
