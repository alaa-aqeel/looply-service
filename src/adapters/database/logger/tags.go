package logger

import "strings"

const (
	TAG_HTTP    = "HTTP"
	TAG_SQL     = "SQL"
	TAG_SERVICE = "SERVICE"
	TAG_REPO    = "REPOSITORY"
)

func MakeTag(parts ...string) string {
	return strings.Join(parts, ":")
}
