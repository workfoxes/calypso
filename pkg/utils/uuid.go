package utils

import (
	"github.com/ventu-io/go-shortid"
	"strings"
)

func ShortUUID(prefix string) string {
	uuidV4, _ := shortid.Generate()
	return strings.ToLower(prefix + uuidV4)
}
