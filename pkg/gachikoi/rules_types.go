package gachikoi

import "regexp"

// replaceRule represents a regex-based replacement rule.
// It is shared by normalRules and finalRules.
type replaceRule struct {
	pattern *regexp.Regexp
	to      string
}
