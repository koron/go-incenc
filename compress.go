package incenc

import "strconv"

func prefix_match(s, prev string) int {
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

func Compress(a []string) string {
	if len(a) == 0 {
		return ""
	}
	b := &buffer{}
	prev := ""
	for _, s := range a {
		if l := prefix_match(s, prev); l > 0 {
			b.WriteString(strconv.Itoa(l))
			b.WriteString(ASCII_US)
			b.WriteString(s[l:])
			b.WriteString(ASCII_RS)
		} else {
			b.WriteString(s)
			b.WriteString(ASCII_RS)
		}
		prev = s
	}
	return string(*b)
}
