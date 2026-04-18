package gachikoi

import (
	"regexp"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type replaceRule struct {
	pattern *regexp.Regexp
	to      string
}

var tokenRules = []replaceRule{
	{regexp.MustCompile(`んちゃ`), `んた`},
	{regexp.MustCompile(`ちゃい`), `さい`},
	{regexp.MustCompile(`ちゃち`), `さし`},
	{regexp.MustCompile(`っちぇ`), `って`},
	{regexp.MustCompile(`ちぇ`), `せ`},
	{regexp.MustCompile(`ちゅ`), `す`},
	{regexp.MustCompile(`ちょ`), `と`},
	{regexp.MustCompile(`ちゃ`), `た`},
	{regexp.MustCompile(`じゃん`), `ざん`},
}

var postRules = []replaceRule{
	{regexp.MustCompile(`にゃい`), `ない`},
	{regexp.MustCompile(`にゃん`), `なん`},
	{regexp.MustCompile(`にゃ([ぁ-ん])`), `な$1`},
	{regexp.MustCompile(`にぇ([ぁ-ん])`), `ね$1`},
	{regexp.MustCompile(`にぇ([。！？!?、…\s]|$)`), `ね$1`},
	{regexp.MustCompile(`にゃ([。！？!?、…\s]|$)`), `ね$1`},
	{regexp.MustCompile(`きゅ([ち〜ー!！\?？。、…])`), `く$1`},
	{regexp.MustCompile(`きゅ$`), `く`},
}

var lexicalRules = strings.NewReplacer(
	"はっぴぃ", "ハッピー",
	"ばえんたいん", "バレンタイン",
	"ちぇぶん", "セブン",
	"かえー", "カレー",
	"めいく", "メイク",
	"にちーく", "にチーク",
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
