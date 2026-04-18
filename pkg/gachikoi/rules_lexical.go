package gachikoi

import "strings"

// lexicalRules contains fixed string replacement rules.
// It is applied first to each token after tokenization.
var lexicalRules = strings.NewReplacer(
	"はっぴぃ", "ハッピー",
	"ばえんたいん", "バレンタイン",
	"ちぇぶん", "セブン",
	"かえー", "カレー",
	"めいく", "メイク",
	"にちーく", "にチーク",
)
