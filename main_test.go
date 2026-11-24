package main

import "testing"

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s, t, want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""}, // multiplicity matters
		{"ab", "b", "b"},
		{"", "a", ""},
		{"abc", "", ""},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
	}
	for _, c := range cases {
		if got := MinWindow(c.s, c.t); got != c.want {
			t.Errorf("MinWindow(%q,%q) = %q, want %q", c.s, c.t, got, c.want)
		}
	}
}
