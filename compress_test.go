package incenc

import "testing"

func checkCompress(t *testing.T, a []string, exp string) {
	s := Compress(a)
	if s != exp {
		t.Errorf("Compress failed: expected=%#v actually=%#v", exp, s)
	}
}

func TestCompressEmpty(t *testing.T) {
	checkCompress(t, []string{}, "")
}

func TestCompressNone(t *testing.T) {
	checkCompress(t, []string{"abc", "def"}, "abc\x1edef\x1e")
}

func TestCompressSimple(t *testing.T) {
	checkCompress(t, []string{"foo", "foobar"},
		"foo\x1e3\x1fbar\x1e")
	checkCompress(t, []string{"ourwork", "ourdestiny"},
		"ourwork\x1e3\x1fdestiny\x1e")
}

func TestCompressShort(t *testing.T) {
	checkCompress(t, []string{"foobar", "foo"},
		"foobar\x1e3\x1f\x1e")
}

func TestCompressCancel(t *testing.T) {
	checkCompress(t, []string{"mywork", "myoffice"},
		"mywork\x1emyoffice\x1e")
	checkCompress(t, []string{"do", "dog"}, "do\x1edog\x1e")
	checkCompress(t, []string{"a", "ab", "abc"},
		"a\x1eab\x1eabc\x1e")
}

func TestCompressMany(t *testing.T) {
	checkCompress(t, []string{"abc", "abcdef", "abcdefghi"},
		"abc\x1e3\x1fdef\x1e6\x1fghi\x1e")
}
