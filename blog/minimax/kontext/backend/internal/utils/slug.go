package utils

import (
	"fmt"
	"strings"
	"time"
)

func GenerateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = fmt.Sprintf("post-%d", time.Now().UnixNano()%1000000)
	}
	slug = fmt.Sprintf("%s-%d", slug, time.Now().UnixNano()%10000)
	return slug
}
