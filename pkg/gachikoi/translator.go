package gachikoi

import (
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var tokenized = mustNewTokenizer()

func mustNewTokenizer() *tokenizer.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	return t
}

func normalizeToken(surface string) string {
	if surface == "" {
		return surface
	}

	normalized := lexicalRules.Replace(surface)
	for _, rule := range tokenRules {
		normalized = rule.pattern.ReplaceAllString(normalized, rule.to)
	}

	return normalized
}

func Translate(input string) string {
	tokens := tokenized.Tokenize(input)
	if len(tokens) == 0 {
		return input
	}

	var builder strings.Builder
	builder.Grow(len(input))

	for _, tok := range tokens {
		surface := tok.Surface
		if surface == "" {
			continue
		}
		builder.WriteString(normalizeToken(surface))
	}

	out := builder.String()
	for _, rule := range postRules {
		out = rule.pattern.ReplaceAllString(out, rule.to)
	}

	return out
}
