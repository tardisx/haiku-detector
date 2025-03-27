// Package haiku finds haiku within English sentences.
package haiku

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/mtso/syllables"
)

type Haiku struct {
	lines []string
}

// Lines returns the Haiku as multiple lines.
func (h Haiku) Lines() []string {
	stripped := []string{}
	for _, l := range h.lines {
		stripped = append(stripped, strings.TrimSpace(l))
	}
	return stripped
}

// String returns the Haiku as a single string, with newlines between
// each line.
func (h Haiku) String() string {
	stripped := []string{}
	for _, l := range h.lines {
		stripped = append(stripped, strings.TrimSpace(l))
	}
	return strings.Join(stripped, "\n")
}

// Sentences
func sentencesFromText(text string) []string {
	// split into sentences
	re := regexp.MustCompile(`(\w[.!\?+])\s+`)

	// Split the text using the regex
	matches := re.FindAllStringIndex(text, -1)
	var sentences []string
	lastIndex := 0

	for _, match := range matches {
		end := match[1] // Include punctuation
		sentences = append(sentences, strings.TrimSpace(text[lastIndex:end]))
		lastIndex = end
	}

	if lastIndex < len(text) {
		sentences = append(sentences, strings.TrimSpace(text[lastIndex:]))
	}

	return sentences
}

func wordsInSentence(s string) []string {
	s = strings.ToLower(s) // aesthetic :-)
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	if s == "" {
		return nil
	}
	return strings.Split(s, " ")
}

func haikuFromSentence(s string) (Haiku, error) {
	words := wordsInSentence(s)
	if len(words) == 0 {
		return Haiku{}, errors.New("sentence has 0 words")
	}
	line := 0
	counts := []int{5, 7, 5}
	wordIdx := 0
	//	lines := make([]string, len(counts))
	haiku := Haiku{
		lines: make([]string, len(counts)),
	}
	for {
		if line == len(counts) && wordIdx == len(words) {
			// we finished the haiku, at the same time as we ran out of words!
			return haiku, nil
		} else if wordIdx == len(words) {
			// we ran out of words before we filled in the haiku
			return Haiku{}, fmt.Errorf("not a haiku - ran out of words at line: %d, counts: %#v, lines: %#v", line, counts, haiku.lines)
		} else if line == len(counts) {
			return Haiku{}, fmt.Errorf("not a haiku - too many words: %d, counts: %#v, lines: %#v", line, counts, haiku.lines)

		}

		thisWord := words[wordIdx]
		counts[line] -= syllables.In(thisWord)
		haiku.lines[line] += thisWord + " "
		if counts[line] == 0 {
			// we finished a line with the right number of syllables, move to next line
			line++
			wordIdx++
			continue
		} else if counts[line] < 0 {
			// blew past the syllable count
			break
		}
		wordIdx++
	}
	return Haiku{}, errors.New("not a haiku")
}

// Find finds 0 or more haiku in an arbitrary string. The string may contain
// one or more sentences, delimited by normal English punctuation. A haiku
// will only be matched against a complete sentence.
func Find(s string) []Haiku {
	h := []Haiku{}
	sentences := sentencesFromText(s)
	for _, sentence := range sentences {
		println(sentence)
		aHaiku, err := haikuFromSentence(sentence)
		if err == nil {
			h = append(h, aHaiku)
		}
	}
	return h
}
