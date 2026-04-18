package gachikoi

import "regexp"

// normalRules contains regular regex replacements applied per token.
// It runs after specialRules to normalize in-token sounds and spellings.
var normalRules = []replaceRule{
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
