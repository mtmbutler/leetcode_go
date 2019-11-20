package genparens

import (
	"reflect"
	"sort"
	"testing"
)

type Case struct {
	n        int
	expected []string
}

var cases = []Case{
	{0, []string{}},
	{1, []string{"()"}},
	{2, []string{"(())", "()()"}},
	{3, []string{"((()))", "()(())", "(())()", "(()())", "()()()"}},
	{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, c := range cases {
		got := generateParenthesis(c.n)
		sort.Strings(got)
		sort.Strings(c.expected)
		bothEmpty := len(got) == 0 && len(c.expected) == 0
		if !reflect.DeepEqual(got, c.expected) && !bothEmpty {
			t.Errorf(
				"generateParenthesis(%d) == %v, want %v",
				c.n,
				got,
				c.expected,
			)
		}
	}
}
