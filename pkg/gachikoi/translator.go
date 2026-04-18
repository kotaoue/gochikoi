package gachikoi

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var (
	tok     *tokenizer.Tokenizer
	tokOnce sync.Once
	tokErr  error
)

func getTokenizer() (*tokenizer.Tokenizer, error) {
	tokOnce.Do(func() {
		tok, tokErr = tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	})
	return tok, tokErr
}

func normalizeToken(surface string) string {
	if surface == "" {
		return surface
	}

	normalized := specialRules.Replace(surface)
	for _, rule := range normalRules {
		normalized = rule.pattern.ReplaceAllString(normalized, rule.to)
	}

	return normalized
}

func Translate(input string) (string, error) {
	t, err := getTokenizer()
	if err != nil {
		return "", fmt.Errorf("gachikoi: tokenizer initialization failed: %w", err)
	}

	tokens := t.Tokenize(input)
	if len(tokens) == 0 {
		return input, nil
	}

	var builder strings.Builder
	builder.Grow(len(input))

	for _, token := range tokens {
		surface := token.Surface
		if surface == "" {
			continue
		}
		builder.WriteString(normalizeToken(surface))
	}

	out := builder.String()
	for _, rule := range finalRules {
		out = rule.pattern.ReplaceAllString(out, rule.to)
	}

	return out, nil
}
