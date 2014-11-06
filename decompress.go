package incenc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Decompress returns an array of string, which decoded from string.
func Decompress(s string) ([]string, error) {
	a := make([]string, 0, strings.Count(s, asciiRS))
	prev := ""
	for s != "" {
		// Split first record.
		x := strings.Index(s, asciiRS)
		if x < 0 {
			return nil, errors.New("not found terminator")
		}
		t := s[:x]
		s = s[x+1:]
		// Try to split "reuse length" unit.
		y := strings.Index(t, asciiUS)
		if y < 0 {
			prev = t
		} else {
			n, err := strconv.ParseInt(t[:y], 10, 32)
			l := int(n)
			if err != nil {
				return nil, fmt.Errorf(
					"failed to parse %#v as reuse length: %s", t[:y], err)
			} else if l > len(prev) {
				return nil, fmt.Errorf("reuse length %d is longer than: %s",
					l, prev)
			}
			prev = prev[:l] + t[y+1:]
		}
		// Emit an element.
		a = append(a, prev)
	}
	return a, nil
}
