package templating

import (
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

// Functions returns the set of custom functions available to templates.
func Functions() template.FuncMap {
	return template.FuncMap{
		// Built-in string manipulation functions
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"trim":  strings.TrimSpace,
		"join":  strings.Join,
		"quote": strconv.Quote,

		// Wrapped so pipeline argument order reads naturally
		"contains":  func(substr, s string) bool { return strings.Contains(s, substr) },
		"hasPrefix": func(prefix, s string) bool { return strings.HasPrefix(s, prefix) },
		"hasSuffix": func(suffix, s string) bool { return strings.HasSuffix(s, suffix) },
		"replace":   func(old, new, s string) string { return strings.ReplaceAll(s, old, new) },
		"split":     func(sep, s string) []string { return strings.Split(s, sep) },

		// Custom string manipulation functions
		"pascal": pascalCase,
	}
}

// pascalCase converts a string to PascalCase. Words are determined by spaces, hyphens, and underscores.
// Example: "hello_world" -> "HelloWorld"
func pascalCase(s string) string {
	var b strings.Builder
	upperNext := true

	runes := []rune(s)
	for i, r := range runes {
		switch {
		case r == ' ' || r == '-' || r == '_':
			upperNext = true
		case upperNext:
			b.WriteRune(unicode.ToUpper(r))
			upperNext = false
		case unicode.IsUpper(r) && i > 0 && unicode.IsLower(runes[i-1]):
			b.WriteRune(r)
		default:
			b.WriteRune(unicode.ToLower(r))
		}
	}

	return b.String()
}
