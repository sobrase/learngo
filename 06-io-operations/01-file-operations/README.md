# Exercise 11: File I/O Operations

## Objective
Learn about Go's file I/O operations using `os`, `io`, and `bufio` packages.

## Problem
Create a program that demonstrates various file operations, reading, writing, and text processing.

## Requirements
1. Use `os` package for file operations:
   - `os.Open()` and `os.Create()`
   - `os.ReadFile()` and `os.WriteFile()`
   - File permissions and modes
2. Use `io` package for I/O primitives:
   - `io.Copy()`, `io.ReadAll()`
   - `io.Reader` and `io.Writer` interfaces
3. Use `bufio` package for buffered I/O:
   - `bufio.NewReader()` and `bufio.NewWriter()`
   - Line-by-line reading
4. Work with directories and file paths
5. Handle errors and implement proper cleanup

## Tasks
1. Create a simple text file and write content
2. Read and process text files line by line
3. Copy files and directories
4. Implement a simple log file processor
5. Work with different file formats (CSV, config files)

## Expected Output
```
File Operations Demo:

Creating and writing to file:
File 'sample.txt' created successfully
Content written: 156 bytes

Reading file content:
Line 1: Hello, World!
Line 2: This is a sample file.
Line 3: Created with Go file operations.
Line 4: 
Line 5: End of file.

File statistics:
Size: 156 bytes
Mode: -rw-r--r--
Is directory: false

Buffered reading:
Processing line: Hello, World!
Processing line: This is a sample file.
Processing line: Created with Go file operations.
Processing line: 
Processing line: End of file.

CSV Processing:
Name,Age,City
John,25,New York
Jane,30,Los Angeles
Bob,35,Chicago

Parsed CSV data:
John is 25 years old from New York
Jane is 30 years old from Los Angeles
Bob is 35 years old from Chicago

File copying:
Source file: sample.txt (156 bytes)
Destination file: sample_copy.txt (156 bytes)
Copy completed successfully

Directory operations:
Created directory: test_dir
Created subdirectory: test_dir/subdir
Listing test_dir contents:
  subdir/ (directory)
  sample.txt (156 bytes)
  sample_copy.txt (156 bytes)

Cleanup:
Removed test files and directories
```

## Hints
- Use `defer file.Close()` to ensure files are closed
- Use `os.Stat()` to get file information
- Use `filepath.Join()` for cross-platform path handling
- Use `bufio.Scanner` for line-by-line reading
- Use `encoding/csv` for CSV file processing
- Remember to handle file permissions
- Use `os.MkdirAll()` for creating nested directories
- Use `filepath.Walk()` for directory traversal
