package gachikoi

import "regexp"

type replaceRule struct {
	pattern *regexp.Regexp
	to      string
}
