package incenc

import (
	"fmt"
	"strconv"
	"strings"
)

func countMatchedPrefix(s, prev string) int {
	// Calc length of matched prefix string.
	l := len(s)
	if n := len(prev); n < l {
		l = n
	}
	for i := 0; i < l; i++ {
		if s[i] != prev[i] {
			l = i
			break
		}
	}
	// Not compress too short prefix.
	if l < 3 {
		return 0
	}
	return l
}

// Compress returns a string which encoded a string array.
func Compress(a []string) (string, error) {
	if len(a) == 0 {
		return "", nil
	}
	b := &buffer{}
	prev := ""
	for i, s := range a {
		if strings.IndexAny(s, asciiRS+asciiUS) >= 0 {
			return "", fmt.Errorf("item %d contains inhibited chars", i)
		}
		if l := countMatchedPrefix(s, prev); l > 0 {
			b.WriteString(strconv.Itoa(l))
			b.WriteString(asciiUS)
			b.WriteString(s[l:])
			b.WriteString(asciiRS)
		} else {
			b.WriteString(s)
			b.WriteString(asciiRS)
		}
		prev = s
	}
	return string(*b), nil
}
