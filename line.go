package goline

import (
	"regexp"
	"strings"
)

const (
	windowsLineSeparator    = "\r\n"
	nonWindowsLineSeparator = "\n"
)

type Line string

func FromMultilineString(content string) []Line {
	normalized := normalizeLineSeparator(content)
	lines := strings.Split(normalized, nonWindowsLineSeparator)
	wrappedLines := []Line{}
	for _, line := range lines {
		wrappedLines = append(wrappedLines, Line(line))
	}
	return wrappedLines
}

func isWindowContent(content string) bool {
	return strings.Contains(content, windowsLineSeparator)
}

func normalizeLineSeparator(content string) string {
	if isWindowContent(content) {
		content = strings.ReplaceAll(content, windowsLineSeparator, nonWindowsLineSeparator)
	}
	return content
}

func (l Line) Trim() Line {
	return Line(strings.TrimSpace(string(l)))
}

func (l Line) Capture(pattern string, index int) (string, bool) {
	captures, ok := l.CaptureAll(pattern)
	if ok {
		if len(captures)-1 < index {
			return "", false
		}
		return captures[index], true
	}
	return "", false
}

func (l Line) CaptureAll(pattern string) ([]string, bool) {
	pat, err := regexp.Compile(pattern)
	if err != nil {
		return nil, false
	}

	matches := pat.FindStringSubmatch(string(l))
	if len(matches) <= 1 {
		return nil, false
	}
	return matches, true
}

func (l Line) LookLike(pattern string) bool {
	pat, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return pat.MatchString(string(l))
}

func (l Line) Value() string {
	return string(l)
}
