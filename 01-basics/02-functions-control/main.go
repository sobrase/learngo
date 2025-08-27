package main

import (
	"fmt"
	"strings"
)

func analyzeText(text string) []any {
	stringss := strings.Split(text, " ")
	wordcount := map[string]int{}
	words := len(stringss)
	a := map[string]int{
		"wordnumber":      0,
		"characternumber": 0,
	}
	returnvalue := []any{}
	for _, i := range stringss {
		a["wordnumber"] += 1
		a["characternumber"] += len(i)
		wc := strings.Split(i, "")
		for _, j := range wc {
			exists := wordcount[j]
			if exists == 0 {
				wordcount[j] = 1
			} else {
				wordcount[j] += 1
			}
		}
	}
	returnvalue = append(returnvalue, words)
	returnvalue = append(returnvalue, a["characternumber"]/a["wordnumber"])
	returnvalue = append(returnvalue, wordcount)
	return returnvalue
	// v := []byte(text)
	// for _, i := range v {
	// 	fmt.Printf("string is : %d %X %s \n", i, i, string(i))
	// 	o := []byte{0x00, 0x01, 0b11111111}
	// 	for _, j := range o {
	// 		fmt.Printf("%d xor %d equal %d \n", i, j, i^j)
	// 		fmt.Printf("%08b xor %08b equal %08b \n", i, j, i^j)
	// 		fmt.Printf("%s xor %s equal %s \n", string(i), string(j), string(i^j))
	// 	}
	// 	return
	//}
}

func filterWords(text string) {
	stopwords := []string{"the", "and", "or", "but", "in", "on", "at", "to", "for", "of", "with", "by"}
	words := strings.Split(text, " ")
	for n, i := range words {
		for _, j := range stopwords {
			if i == j {
				words = append(words[:n], words[n+1:]...)
			}
		}
	}
	fmt.Print(words)

}

func main() {
	fmt.Println("Exercise 2: Text Analysis & Word Cloud Generator")
	fmt.Println("===============================================")

	// Sample text for analysis
	sampleText := "The quick brown fox jumps over the lazy dog. Programming in Go is fun and efficient! The quick fox loves programming."
	table := analyzeText(sampleText)
	fmt.Println(table)
	// TODO: Task 1 - Create text analysis functions
	// Create functions to analyze text:
	// - analyzeText(text string) - Returns word count, character count, average word length
	// - getWordFrequency(text string) - Returns map of words and their frequencies
	// - filterWords(words []string, minLength int) - Returns filtered words
	// - calculateTextStats(text string) - Returns comprehensive statistics

	// TODO: Task 2 - Create word cloud generation functions
	// Create functions to generate the word cloud:
	// - generateWordCloud(frequencies map[string]int) - Creates visual representation
	// - getWordSize(frequency int, maxFreq int) - Determines visual size
	// - printWordCloud(wordMap map[string]int) - Displays the cloud

	// TODO: Task 3 - Implement text processing
	// - Remove common stop words (the, and, or, but, in, on, at, to, for, of, with, by)
	// - Convert to lowercase
	// - Remove punctuation
	// - Sort words by frequency

	// TODO: Task 4 - Print the word cloud
	// Generate a beautiful word cloud that shows:
	// - Text statistics analysis
	// - Visual word cloud representation
	// - Word frequency table
	// - Formatted analysis report

	fmt.Printf("\n📝 Analyzing text: \"%s\"\n", sampleText)
	fmt.Println("Complete the TODO tasks above to generate a beautiful word cloud!")
	filterWords(sampleText)
}
