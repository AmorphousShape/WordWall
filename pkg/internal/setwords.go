package internal

import (
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

	for i, char := range word {
		variants, ok := CharVariants[unicode.ToLower(char)]
		if !ok {
			variants = []string{string(char)}
		}

		// Group of variants (e.g., a|@|4)
		sb.WriteString("(?:")
		for j, v := range variants {
			if j > 0 {
				sb.WriteString("|")
			}
			sb.WriteString(regexp.QuoteMeta(v))
		}
		sb.WriteString(")")

		// Between-character obfuscation only if not last character
		if i < len(word)-1 {
			sb.WriteString("[\\s._\\-~']*")
		}
	}

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
