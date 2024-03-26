package utils

import (
	"strings"

	"github.com/adshin21/go-fts/constants"
	snowballeng "github.com/kljensen/snowball/english"
)

// return lowercase tokens
func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for id, tok := range tokens {
		r[id] = strings.ToLower(tok)
	}
	return r
}

// removes stopwords from tokens
func stopwordFilter(tokens []string) []string {
	stopwords := constants.StopWords
	r := make([]string, 0, len(tokens))

	for _, tok := range tokens {
		if _, ok := stopwords[tok]; !ok {
			r = append(r, tok)
		}
	}

	return r
}

// stemming -> the process of reducing a word to its word stem that affixes
// to suffixes and prefixes or the roots.
// Eg - waited, waiting, waits ===(to)==> wait
func stemmingFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
