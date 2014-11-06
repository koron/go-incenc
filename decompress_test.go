package incenc

import "testing"

func checkDecompress(t *testing.T, s string, exp []string) {
	a, err := Decompress(s)
	if err != nil {
		t.Errorf("Decompress failed: %s", err)
	}
	// check length.
	l, m := len(a), len(exp)
	if l != m {
		t.Errorf("Decompress failed: length %d mismatches expected %d", l, m)
	}
	if m < l {
		l = m
	}
	// check elements.
	for i := 0; i < l; i++ {
		if a[i] != exp[i] {
			t.Errorf("Decompress failed: #%d item %#v mismatches expected %#v", a[i], exp[i])
		}
	}
}

func TestDecompressEmpty(t *testing.T) {
	checkDecompress(t, "", []string{})
}

func TestDecompressNone(t *testing.T) {
	checkDecompress(t, "abc\x1edef\x1e", []string{"abc", "def"})
}

func TestDecompressSimple(t *testing.T) {
	checkDecompress(t, "foo\x1e3\x1fbar\x1e", []string{"foo", "foobar"})
	checkDecompress(t, "ourwork\x1e3\x1fdestiny\x1e",
		[]string{"ourwork", "ourdestiny"})
}

func TestDecompressShort(t *testing.T) {
	checkDecompress(t, "foobar\x1e3\x1f\x1e", []string{"foobar", "foo"})
}

func TestDecompressCancel(t *testing.T) {
	checkDecompress(t, "mywork\x1emyoffice\x1e",
		[]string{"mywork", "myoffice"})
	checkDecompress(t, "do\x1edog\x1e", []string{"do", "dog"})
	checkDecompress(t, "a\x1eab\x1eabc\x1e", []string{"a", "ab", "abc"})
}

func TestDecompressMany(t *testing.T) {
	checkDecompress(t, "abc\x1e3\x1fdef\x1e6\x1fghi\x1e",
		[]string{"abc", "abcdef", "abcdefghi"})
}
