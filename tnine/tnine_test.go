package tnine

import (
	"reflect"
	"testing"
)

type Case struct {
	digits   string
	expected []string
}

var cases = []Case{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"", []string{}},
	{"2", []string{"a", "b", "c"}},
	{"84", []string{"tg", "th", "ti", "ug", "uh", "ui", "vg", "vh", "vi"}},
}

func TestLetterCombinations(t *testing.T) {
	for _, c := range cases {
		got := letterCombinations(c.digits)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"letterCombinations(%s) == %v, want %v",
				c.digits,
				got,
				c.expected,
			)
		}
	}
}
