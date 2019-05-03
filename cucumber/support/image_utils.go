package support

import (
	"strings"
)

// ImageExtention returns the extension of a provided url.
func ImageExtention(imageURL string) string {
	s := strings.Split(imageURL, ".")
	return s[len(s)-1]
}
