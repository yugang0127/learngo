package main

import "testing"

func TestNonRepeatingStr(t *testing.T) {
	tests := []struct{s string
	answer int}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _,tt := range tests {
		actual := lengthOfNonRepeatingStr(tt.s)
		if actual != tt.answer {
			t.Errorf("get %d for input %s, expected %d", actual, tt.s, tt.answer)
		}
	}
}

func BenchmarkNonRepeatingStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	answer := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingStr(s)
		if actual != answer {
			b.Errorf("get %d for input %s, expected %d", actual, s, answer)
		}
	}
}