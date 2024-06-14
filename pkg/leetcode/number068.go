package leetcode

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var curWords []string
	curLen := 0

	for _, word := range words {
		if curLen + len(word) + len(curWords) > maxWidth {
			result = append(result, justifyLine(curWords, maxWidth, curLen))
			curWords = []string{}
			curLen = 0
		}
		curWords = append(curWords, word)
		curLen += len(word)
	}

	if len(curWords) > 0 {
		lastLine := strings.Join(curWords, " ")
		lastLine += strings.Repeat(" ", maxWidth - len(lastLine))
		result = append(result, lastLine)
	}

	return result
}

func justifyLine(words []string, maxWidth int, curLen int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth - curLen)
	}

	totalSpaces := maxWidth - curLen
	gaps := len(words) - 1
	spacesPerGap := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var result strings.Builder
	for i := 0; i < len(words) - 1; i++ {
		result.WriteString(words[i])
		spaceCount := spacesPerGap
		if i < extraSpaces {
			spaceCount++
		}
		result.WriteString(strings.Repeat(" ", spaceCount))
	}
	result.WriteString(words[len(words) - 1])

	return result.String()
}
