package leetcode

import (
	"fmt"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	result := fullJustify(words, 16)
	for _, line := range result {
		fmt.Println(line)
	}
}