package gachikoi

import "regexp"

// replaceRule represents a regex-based replacement rule.
// It is shared by tokenRules and postRules.
type replaceRule struct {
	pattern *regexp.Regexp
	to      string
}
