package defangipaddr

import "testing"

type Case struct {
	address  string
	expected string
}

var cases = []Case{
	{"", ""},
	{".", "[.]"},
	{"1.1.1.1", "1[.]1[.]1[.]1"},
	{"255.100.50.0", "255[.]100[.]50[.]0"},
}

func TestDefangIPaddr(t *testing.T) {
	for _, c := range cases {
		got := defangIPaddr(c.address)
		if got != c.expected {
			t.Errorf(
				"defangIpAddr(%s) == %s, want %s",
				c.address,
				got,
				c.expected,
			)
		}
	}
}
