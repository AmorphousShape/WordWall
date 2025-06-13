package internal

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type Rule int32

const (
	RuleCensor Rule = iota
	RuleFilterMessage
	RuleInstantBan
)

type BannedWord struct {
	Word     string
	Regex    *regexp.Regexp
	Response Rule
}

var BannedWords = []BannedWord{}

// ObfuscatedRegex generates a regex pattern that matches a word with obfuscation.
func ObfuscatedRegex(word string) string {
	var sb strings.Builder
	sb.WriteString("(?i)") // Case-insensitive

	const noise = ".{0,2}?" // Match up to 2 of ANY character, non-greedy

	runes := []rune(word)
	for i, char := range runes {
		baseChar := unicode.ToLower(char)

		variants, ok := CharVariants[baseChar]
		if !ok {
			variants = []string{string(char)}
		}

		// Character or variants group
		sb.WriteString("(?:")
		for j, v := range variants {
			if j > 0 {
				sb.WriteString("|")
			}
			sb.WriteString(regexp.QuoteMeta(v))
		}
		sb.WriteString(")")

		// Add noise after the character if not the last
		if i < len(runes)-1 && i > 0 {
			sb.WriteString(noise)
		}

		sb.WriteString("[^\\s]*")
	}

	fmt.Println("Generated regex for word:", word, "is", sb.String())

	return sb.String()
}

// SetBannedWords sets the list of banned words with their responses.
// It creates a regex for each word that allows for obfuscation and variations.
// The response defines how the word should be handled (censored, filtered, or banned).
func SetBannedWords(words []string, response Rule) {

	// Sort words by length in descending order to ensure longer words are matched first
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j]) // longest to shortest
	})

	BannedWords = make([]BannedWord, len(words))
	for i, word := range words {
		regex := ObfuscatedRegex(word)
		BannedWords[i] = BannedWord{
			Word:     word,
			Regex:    regexp.MustCompile(regex),
			Response: response,
		}
	}
}
