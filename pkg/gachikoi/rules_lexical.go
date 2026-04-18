package gachikoi

import "strings"

var lexicalRules = strings.NewReplacer(
	"はっぴぃ", "ハッピー",
	"ばえんたいん", "バレンタイン",
	"ちぇぶん", "セブン",
	"かえー", "カレー",
	"めいく", "メイク",
	"にちーく", "にチーク",
)
