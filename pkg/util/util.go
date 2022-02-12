package util

import "strings"

func ParseAllowedOrigins(str string) []string {
	origins := strings.Split(strings.ToLower(str), ",")

	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}

	return origins
}