package main

import (
	"fmt"
	"regexp"
	"strings"
)

func createTokenPattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile("\\A(" + pattern + ")")
}

var tokenTypes = map[string]*regexp.Regexp{
	"def":        createTokenPattern("\\bdef\\b"),
	"end":        createTokenPattern("\\bend\\b"),
	"identifier": createTokenPattern("\\b[a-zA-Z]+\\b"),
	"integer":    createTokenPattern("\\b[0-9]+\\b"),
	"oparen":     createTokenPattern("\\("),
	"cparen":     createTokenPattern("\\)"),
}

type Tokenizer struct {
	Source string
}

type Token struct {
	Symbol string
	Value  string
}

func tokenize(source string) []Token {
	fmt.Println(source)
	tokens := []Token{}
	for token, re := range tokenTypes {
		loc := re.FindStringIndex(source)
		if loc != nil {
			value := source[loc[0]:loc[1]]
			tokens = append(tokens, Token{token, value})
			newSource := strings.TrimSpace(source[loc[1]:len(source)])
			tokens = append(tokens, tokenize(newSource)...)
			return tokens
		}
	}
	return tokens
}
