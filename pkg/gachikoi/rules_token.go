package gachikoi

import "regexp"

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
