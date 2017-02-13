package kit

import (
	"strings"
)

func ParseName(s string) (string, string) {
	l := strings.Split(s, "/")
	if len(l) == 0 {
		return "", ""
	}
	if len(l) == 1 {
		return l[0], ""
	}
	return l[0], l[1]
}
