package main

import (
	"regexp"
)

func validRepoName(s string) bool {
	ok, err := regexp.MatchString(`^[a-z][a-z0-9-]+$`, s)
	return err == nil && ok
}
