package wordwall

import (
	"strings"

	"github.com/AmorphousShape/wordwall/pkg/internal"
)

// SetCensoredWords sets the list of censored words for the WordWall.
// Censored words will be replaced with asterisks in messages.
func SetCensoredWords(words []string) {
	internal.SetBannedWords(words, internal.RuleCensor)
}

// SetFilteredWords sets the list of filtered words for the WordWall.
// Filtered words will cause messages to be filtered out if they contain any of the words.
func SetFilteredWords(words []string) {
	internal.SetBannedWords(words, internal.RuleFilterMessage)
}

// SetZeroToleranceWords sets the list of zero-tolerance words for the WordWall.
// Zero-tolerance words will result in an instant ban if they are found in a message.
func SetZeroToleranceWords(words []string) {
	internal.SetBannedWords(words, internal.RuleInstantBan)
}

// FilterString processes a message and returns a list of filtered words.
func FilterString(message string) (newMessage string, hitCensor bool, hitFilter bool, hitZeroTolerance bool) {
	newMessage = message

	for _, bw := range internal.BannedWords {
		if bw.Regex.MatchString(message) {
			switch bw.Response {
			case internal.RuleCensor:
				hitCensor = true
				// Replace banned word with asterisks or a marker
				newMessage = bw.Regex.ReplaceAllStringFunc(newMessage, func(matched string) string {
					return strings.Repeat("*", len([]rune(matched)))
				})

			// Filtered words cause the message to be empty
			case internal.RuleFilterMessage:
				newMessage = ""
				hitFilter = true

			// Zero-tolerance words also result in an empty message but are handled in a separate case for clarity
			case internal.RuleInstantBan:
				newMessage = ""
				hitZeroTolerance = true
			}
		}
	}

	return
}
