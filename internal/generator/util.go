package generator

import (
	"fmt"
	"strings"
)

func wrapFileName(path string, ext string) string {
	switch ext {
	case ".tmpl":
		return strings.TrimSuffix(path, ext)
	default:
		return path
	}
}

func getTemplateName(name string) string {
	return fmt.Sprintf("%s_tmpl", name)
}
