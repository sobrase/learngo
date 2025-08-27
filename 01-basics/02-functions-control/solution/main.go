package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function with no parameters and no return value
func printHeader() {
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    TEXT ANALYSIS REPORT                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
}

// Function with parameters and return value
func analyzeText(text string) (int, int, int, float64) {
	length := len(text)
	words := strings.Fields(text)
	wordCount := len(words)

	var totalWordLength int
	for _, word := range words {
		totalWordLength += len(word)
	}

	var avgWordLength float64
	if wordCount > 0 {
		avgWordLength = float64(totalWordLength) / float64(wordCount)
	}

	return length, wordCount, totalWordLength, avgWordLength
}

// Function with multiple return values
func getWordFrequency(text string) (map[string]int, int) {
	words := strings.Fields(strings.ToLower(text))
	frequencies := make(map[string]int)

	for _, word := range words {
		// Remove punctuation
		word = strings.Trim(word, ".,!?;:\"()[]{}")
		if word != "" {
			frequencies[word]++
		}
	}

	return frequencies, len(frequencies)
}

// Function with named return values
func filterWords(words []string, minLength int) (filtered []string) {
	for _, word := range words {
		if len(word) >= minLength {
			filtered = append(filtered, word)
		}
	}
	return
}

// Variadic function
func calculateTextStats(text string, stopWords ...string) map[string]interface{} {
	stats := make(map[string]interface{})

	// Basic stats
	length, wordCount, totalLength, avgLength := analyzeText(text)
	stats["totalCharacters"] = length
	stats["wordCount"] = wordCount
	stats["totalWordLength"] = totalLength
	stats["averageWordLength"] = avgLength

	// Word frequency
	frequencies, uniqueWords := getWordFrequency(text)
	stats["uniqueWords"] = uniqueWords
	stats["wordFrequencies"] = frequencies

	// Find most common word
	var mostCommonWord string
	var maxFreq int
	for word, freq := range frequencies {
		if freq > maxFreq {
			maxFreq = freq
			mostCommonWord = word
		}
	}
	stats["mostCommonWord"] = mostCommonWord
	stats["maxFrequency"] = maxFreq

	return stats
}

// Function to get word size for visualization
func getWordSize(frequency int, maxFreq int) string {
	if maxFreq == 0 {
		return "normal"
	}

	percentage := float64(frequency) / float64(maxFreq)
	switch {
	case percentage >= 0.8:
		return "large"
	case percentage >= 0.5:
		return "medium"
	default:
		return "small"
	}
}

// Function to print word cloud
func printWordCloud(frequencies map[string]int) {
	// Sort words by frequency
	type wordFreq struct {
		word  string
		count int
	}

	var wordList []wordFreq
	for word, count := range frequencies {
		wordList = append(wordList, wordFreq{word, count})
	}

	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].count > wordList[j].count
	})

	// Find max frequency
	var maxFreq int
	for _, wf := range wordList {
		if wf.count > maxFreq {
			maxFreq = wf.count
		}
	}

	fmt.Println("║  WORD CLOUD:                                                 ║")
	fmt.Println("║                                                              ║")

	// Print word cloud in rows
	wordsPerRow := 5
	for i := 0; i < len(wordList); i += wordsPerRow {
		fmt.Print("║        ████████████████████████████████████████████████████  ║\n")
		fmt.Print("║        █")

		for j := 0; j < wordsPerRow && i+j < len(wordList); j++ {
			word := strings.ToUpper(wordList[i+j].word)
			size := getWordSize(wordList[i+j].count, maxFreq)

			switch size {
			case "large":
				fmt.Printf("  %-12s", word)
			case "medium":
				fmt.Printf("  %-10s", word)
			default:
				fmt.Printf("  %-8s", word)
			}
			fmt.Print("  █")
		}
		fmt.Println()
		fmt.Print("║        ████████████████████████████████████████████████████  ║\n")
		fmt.Println("║                                                              ║")
	}
}

// Function to print word frequency table
func printFrequencyTable(frequencies map[string]int) {
	fmt.Println("║  WORD FREQUENCIES:                                           ║")
	fmt.Println("║  ┌─────────────┬─────────────┬─────────────┬─────────────┐   ║")
	fmt.Println("║  │ Word        │ Frequency   │ Length      │ Percentage  │   ║")
	fmt.Println("║  ├─────────────┼─────────────┼─────────────┼─────────────┤   ║")

	// Sort words alphabetically for table
	var words []string
	for word := range frequencies {
		words = append(words, word)
	}
	sort.Strings(words)

	totalWords := 0
	for _, count := range frequencies {
		totalWords += count
	}

	for _, word := range words {
		count := frequencies[word]
		percentage := float64(count) / float64(totalWords) * 100
		fmt.Printf("║  │ %-15s │     %2d       │     %2d       │   %1.2f%%     │   ║\n",
			word, count, len(word), percentage)
	}

	fmt.Println("║  └─────────────┴─────────────┴─────────────┴─────────────┘   ║")
}

func main() {
	fmt.Println("Exercise 2: Text Analysis & Word Cloud Generator")
	fmt.Println("===============================================")

	// Sample text for analysis
	sampleText := "The quick brown fox jumps over the lazy dog. Programming in Go is fun and efficient! The quick fox loves programming."

	// Task 1: Analyze text
	length, wordCount, _, avgLength := analyzeText(sampleText)
	frequencies, uniqueWords := getWordFrequency(sampleText)

	// Task 2: Calculate comprehensive statistics
	stats := calculateTextStats(sampleText, "the", "and", "or", "but", "in", "on", "at", "to", "for", "of", "with", "by")

	// Task 3: Print the beautiful analysis report
	printHeader()

	fmt.Printf("║  Original Text: \"%s\"\n", sampleText)
	fmt.Printf("║  dog. Programming in Go is fun and efficient! The quick fox loves programming.\"              ║\n")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  TEXT STATISTICS:                                            ║")
	fmt.Printf("║  • Total characters: %d                                      ║\n", length)
	fmt.Printf("║  • Word count: %d                                            ║\n", wordCount)
	fmt.Printf("║  • Unique words: %d                                          ║\n", uniqueWords)
	fmt.Printf("║  • Average word length: %.1f                                  ║\n", avgLength)
	fmt.Printf("║  • Most common word: \"%s\" (%d occurrences)                  ║\n",
		stats["mostCommonWord"], stats["maxFrequency"])

	// Task 4: Print word cloud and frequency table
	printWordCloud(frequencies)
	printFrequencyTable(frequencies)

	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	// Demonstrate defer usage
	defer func() {
		fmt.Println("\n📊 Analysis completed successfully!")
	}()

	// Demonstrate switch statement
	fmt.Println("\n🔍 Text Analysis Summary:")
	switch {
	case wordCount > 20:
		fmt.Println("• Long text detected")
	case wordCount > 10:
		fmt.Println("• Medium text detected")
	default:
		fmt.Println("• Short text detected")
	}
}
