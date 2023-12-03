package bprimes

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    uint
		want bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{9, false},
		{11, true},
		{13, true},
		{15, false},
		{17, true},
		{19, true},
		{21, false},
		{23, true},
		{25, false},
		{27, false},
		{29, true},
		{31, true},
		{33, false},
		{35, false},
		{37, true},
		{39, false},
		{41, true},
		{43, true},
		{45, false},
		{47, true},
		{49, false},
		{51, false},
		{53, true},
		{55, false},
		{57, false},
		{59, true},
		{61, true},
		{63, false},
		{65, false},
		{67, true},
		{69, false},
		{71, true},
		{73, true},
		{75, false},
		{77, false},
		{79, true},
		{81, false},
		{83, true},
		{85, false},
		{87, false},
		{89, true},
		{91, false},
		{93, false},
		{95, false},
		{97, true},
		{99, false},
		{100, false},
		{101, true},
		{307, true},
		{503, true},
		{997, true},
		{1001, false},
		{4001, true},
		{4993, true},
		{4999, true},
		{99989, true},
		{99991, true},
		{99999, false},
		{100000, false},
		{100003, true},
		{899981, true},
		{899987, false},
		{899999, false},
		{900001, true},
		{999983, true},
		{999999, false},
		{1000000, false},
	}
	t.Run("fast", func(t *testing.T) {
		for _, tt := range tests {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		}
	})
}
