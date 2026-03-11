package core

import "regexp"

var (
	nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9\s]+`)
	multipleHyphensRegex = regexp.MustCompile(`-+`)
)
