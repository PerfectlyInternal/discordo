package markdown

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	boldRe          = regexp.MustCompile(`(?ms)\*\*(.*?)\*\*`)
	italicRe        = regexp.MustCompile(`(?ms)\*(.*?)\*`)
	underlineRe     = regexp.MustCompile(`(?ms)__(.*?)__`)
	strikethroughRe = regexp.MustCompile(`(?ms)~~(.*?)~~`)
	spoilersRe	= regexp.MustCompile(`(?ms)\|\|(.*?)\|\|`)
	codeblockRe     = regexp.MustCompile("(?ms)`" + `([^` + "`" + `\n]+)` + "`")
)

func Parse(input string, spoilers bool) string {
	input = boldRe.ReplaceAllString(input, "[::b]$1[::-]")
	input = italicRe.ReplaceAllString(input, "[::i]$1[::-]")
	input = underlineRe.ReplaceAllString(input, "[::u]$1[::-]")
	input = strikethroughRe.ReplaceAllString(input, "[::s]$1[::-]")
	if (spoilers) { 
		input = spoilersRe.ReplaceAllStringFunc(input, GenSpoiler) 
	} else {
		input = spoilersRe.ReplaceAllString(input, "$1") 
	}
	input = codeblockRe.ReplaceAllString(input, "[::r]$1[::-]")
	return input
}

func GenSpoiler(input string) string {
	runeCount := utf8.RuneCountInString(input) - 4 // remove the spoiler marks
	return strings.Repeat("█", runeCount)
}
