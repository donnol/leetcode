package no8

import "testing"

func Test_myAtoi(t *testing.T) {
	a := 2<<31 - 1
	b := -2 << 30
	t.Log(a, b)

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{s: "42"}, 42},
		{"", args{s: "-42"}, -42},
		{"", args{s: " -42"}, -42},
		{"", args{s: "  -42"}, -42},
		{"", args{s: "   "}, 0},
		{"", args{s: "4193 with words"}, 4193},
		{"", args{s: "-91283472332"}, -2147483648},
		{"", args{s: "21474836460"}, 2147483647},
		{"", args{s: "2147483648"}, 2147483647},
		{"", args{s: "9223372036854775808"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
