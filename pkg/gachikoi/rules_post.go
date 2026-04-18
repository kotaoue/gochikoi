package gachikoi

import "regexp"

// finalRules contains final post-processing rules applied after tokens are joined.
// It handles boundary-crossing patterns and ending-like variations.
var finalRules = []replaceRule{
	{regexp.MustCompile(`にゃい`), `ない`},
	{regexp.MustCompile(`にゃん`), `なん`},
	{regexp.MustCompile(`にゃ([ぁ-ん])`), `な$1`},
	{regexp.MustCompile(`にぇ([ぁ-ん])`), `ね$1`},
	{regexp.MustCompile(`にぇ([。！？!?、…\s]|$)`), `ね$1`},
	{regexp.MustCompile(`にゃ([。！？!?、…\s]|$)`), `ね$1`},
	{regexp.MustCompile(`きゅ([ち〜ー!！\?？。、…])`), `く$1`},
	{regexp.MustCompile(`きゅ$`), `く`},
}
