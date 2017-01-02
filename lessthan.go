package main

import (
	"regexp"
	"strings"

	"github.com/hashicorp/go-version"
)

func lessThan(v1 string, v2 string) (bool, error) {
	re := regexp.MustCompile("[.0-9]+")
	v1 = strings.Join(re.FindAllString(strings.Fields(v1)[0], -1), "")
	v2 = strings.Join(re.FindAllString(strings.Fields(v2)[0], -1), "")
	foo, err := version.NewVersion(v1)
	bar, err := version.NewVersion(v2)
	if foo.LessThan(bar) || foo.Equal(bar) {
		return true, nil
	}
	return false, err
}
