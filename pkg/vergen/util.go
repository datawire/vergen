package vergen

import (
	"regexp"
	"strings"
)

func normalizeBranchName(name string) string {
	r, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return strings.Trim(r.ReplaceAllString(name, "-"), "-")
}
