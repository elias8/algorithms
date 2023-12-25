package string

import "slices"

type textStack struct {
	data []byte
}

func (s *textStack) push(b byte) {
	if b == '#' {
		if len(s.data) > 0 {
			s.data = s.data[:len(s.data)-1]
		}
	} else if b != '#' {
		s.data = append(s.data, b)
	}
}

// 844. Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
//
// O(max(n, m)) time , O(n+m) space
func backspaceCompare(s string, t string) bool {
	var i, j int
	var ss, ts textStack
	for i < len(s) || j < len(t) {
		if i < len(s) {
			ss.push(s[i])
			i++
		}
		if j < len(t) {
			ts.push(t[j])
			j++
		}
	}
	return slices.Equal(ss.data, ts.data)
}
