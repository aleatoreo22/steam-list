package model

import (
	"regexp"
	"strconv"
	"strings"
)

func GetStringLen(tag string) int {
	re := regexp.MustCompile(`strlen\s+(\d+)`)
	matches := re.FindStringSubmatch(tag)
	if len(matches) < 2 {
		return 0
	}
	strlen, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}
	return strlen
}

func IsPrimaryKey(tag string) bool {
	return strings.Contains(tag, "key")
}
