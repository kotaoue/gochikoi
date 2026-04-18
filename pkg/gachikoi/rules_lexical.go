package gachikoi

import "strings"

// specialRules contains fixed string replacements for known special patterns.
// It is applied first to each token after tokenization.
var specialRules = strings.NewReplacer(
	"はっぴぃ", "ハッピー",
	"ばえんたいん", "バレンタイン",
	"ちぇぶん", "セブン",
	"かえー", "カレー",
	"めいく", "メイク",
	"にちーく", "にチーク",
)
