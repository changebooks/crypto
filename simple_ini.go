package crypto

import (
	"bytes"
	"strings"
	"unicode"
)

const EOL = "\n"

func SimpleIniParse(s string) map[string]string {
	r := make(map[string]string)

	lines := strings.Split(s, EOL)
	for _, l := range lines {
		name, value := iniParseLine(l)
		if name != "" {
			r[name] = value
		}
	}

	return r
}

func iniParseLine(s string) (string, string) {
	if s = strings.TrimFunc(s, iniIsLineBreak); s == "" {
		return "", ""
	}

	// 注释
	if s[0] == '#' {
		return "", ""
	}

	// 名称=值
	if strings.IndexRune(s, '=') > 0 {
		parts := strings.SplitN(s, "=", 2)
		if len(parts) != 2 {
			return "", ""
		}

		return iniSanitiseName(parts[0]), iniSanitiseValue(parts[1])
	}

	return "", ""
}

func iniIsLineBreak(r rune) bool {
	return r == '\r' || r == '\n'
}

// 清除：'\''、'"'、' '、'\\'
// 阻断：'#', '\t', '\n', '\v', '\f', '\r', 0x85, 0xA0
func iniSanitiseName(s string) string {
	var bucket bytes.Buffer

	for _, c := range s {
		if c == '\'' || c == '"' || c == ' ' || c == '\\' {
			continue
		}

		if c == '#' || unicode.IsSpace(c) {
			break
		}

		bucket.WriteRune(c)
	}

	return bucket.String()
}

// 阻断：'#', '\r', '\n'
func iniSanitiseValue(s string) string {
	var bucket bytes.Buffer

	for _, c := range s {
		if c == '#' || c == '\r' || c == '\n' {
			break
		}

		bucket.WriteRune(c)
	}

	return bucket.String()
}
