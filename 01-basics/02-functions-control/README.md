# Exercise 2: Text Analysis & Word Cloud Generator

## 🎯 **Objective**
Create a text analysis tool that processes text and generates a visual word cloud representation using Go functions and control structures.

## 🎨 **Goal: Print a Text Word Cloud**

You will create a program that analyzes text input and displays a beautiful word cloud showing word frequencies with visual representations.

## 📋 **Requirements**
1. Create functions with different signatures:
   - Function with no parameters and no return value
   - Function with parameters and return value
   - Function with multiple return values
   - Function with named return values
   - Variadic function
2. Use control structures:
   - `if` statements with initialization
   - `for` loops (traditional, range-based, while-style)
   - `switch` statements
   - `defer` statements
3. Use the `strings` package for string manipulation
4. Use `strconv` package for string conversions

## 🎯 **Specific Tasks**

### **Task 1: Text Analysis Functions**
Create functions to analyze text:
- `analyzeText(text string)` - Returns word count, character count, average word length
- `getWordFrequency(text string)` - Returns map of words and their frequencies
- `filterWords(words []string, minLength int)` - Returns filtered words
- `calculateTextStats(text string)` - Returns comprehensive statistics

### **Task 2: Word Cloud Generation**
Create functions to generate the word cloud:
- `generateWordCloud(frequencies map[string]int)` - Creates visual representation
- `getWordSize(frequency int, maxFreq int)` - Determines visual size
- `printWordCloud(wordMap map[string]int)` - Displays the cloud

### **Task 3: Text Processing**
- Remove common stop words (the, and, or, but, in, on, at, to, for, of, with, by)
- Convert to lowercase
- Remove punctuation
- Sort words by frequency

### **Task 4: Print the Word Cloud**
Generate a beautiful word cloud that looks like this:

```
╔══════════════════════════════════════════════════════════════╗
║                    TEXT ANALYSIS REPORT                      ║
╠══════════════════════════════════════════════════════════════╣
║  Original Text: "The quick brown fox jumps over the lazy    ║
║  dog. Programming in Go is fun and efficient!"              ║
╠══════════════════════════════════════════════════════════════╣
║  TEXT STATISTICS:                                            ║
║  • Total characters: 89                                      ║
║  • Word count: 12                                            ║
║  • Unique words: 10                                          ║
║  • Average word length: 4.2                                  ║
║  • Most common word: "the" (2 occurrences)                  ║
╠══════════════════════════════════════════════════════════════╣
║  WORD CLOUD:                                                 ║
║                                                              ║
║        ████████████████████████████████████████████████████  ║
║        █  THE  █  QUICK  █  BROWN  █  FOX  █  JUMPS  █     ║
║        ████████████████████████████████████████████████████  ║
║                                                              ║
║        ████████████████████████████████████████████████████  ║
║        █  OVER  █  LAZY  █  DOG  █  PROGRAMMING  █  GO  █   ║
║        ████████████████████████████████████████████████████  ║
║                                                              ║
║        ████████████████████████████████████████████████████  ║
║        █  IS  █  FUN  █  AND  █  EFFICIENT  █              ║
║        ████████████████████████████████████████████████████  ║
║                                                              ║
║  WORD FREQUENCIES:                                           ║
║  ┌─────────────┬─────────────┬─────────────┬─────────────┐   ║
║  │ Word        │ Frequency   │ Length      │ Percentage  │   ║
║  ├─────────────┼─────────────┼─────────────┼─────────────┤   ║
║  │ the         │     2       │     3       │   16.7%     │   ║
║  │ quick       │     1       │     5       │    8.3%     │   ║
║  │ brown       │     1       │     5       │    8.3%     │   ║
║  │ fox         │     1       │     3       │    8.3%     │   ║
║  │ jumps       │     1       │     5       │    8.3%     │   ║
║  │ over        │     1       │     4       │    8.3%     │   ║
║  │ lazy        │     1       │     4       │    8.3%     │   ║
║  │ dog         │     1       │     3       │    8.3%     │   ║
║  │ programming │     1       │    11       │    8.3%     │   ║
║  │ go          │     1       │     2       │    8.3%     │   ║
║  │ is          │     1       │     2       │    8.3%     │   ║
║  │ fun         │     1       │     3       │    8.3%     │   ║
║  │ and         │     1       │     3       │    8.3%     │   ║
║  │ efficient   │     1       │     9       │    8.3%     │   ║
║  └─────────────┴─────────────┴─────────────┴─────────────┘   ║
╚══════════════════════════════════════════════════════════════╝
```

## 🎯 **Expected Output**
Your program should produce:
- Text statistics analysis
- Visual word cloud representation
- Word frequency table
- Formatted analysis report

## 💡 **Hints**
- Use `strings.Fields()` to split text into words
- Use `strings.ToLower()` for case-insensitive analysis
- Use `strings.Trim()` to remove punctuation
- Use `fmt.Printf` with width specifiers for alignment
- Use Unicode box characters for the visual elements
- Use `defer` for cleanup operations

## 🚀 **Bonus Challenges**
1. Add color coding based on word frequency
2. Create different word cloud shapes (circle, rectangle, triangle)
3. Add word sentiment analysis (positive/negative words)
4. Generate word cloud from file input
5. Add interactive word cloud (clickable words)
6. Create animated word cloud (words appearing one by one)

## 🎨 **Learning Outcomes**
By completing this exercise, you will:
- Master Go function signatures and return values
- Understand control structures and loops
- Learn string manipulation techniques
- Create visual data representations
- Practice real-world text processing

This exercise transforms text analysis into a beautiful, visual experience! 📊
