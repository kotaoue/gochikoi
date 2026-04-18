package gachikoi

import "regexp"

// postRules contains post-processing rules applied after tokens are joined.
// It finalizes normalization for boundary-crossing and ending variations.
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
