package no10

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{s: "aa", p: "a"}, want: false},
		{name: "", args: args{s: "aa", p: "a*"}, want: true},
		{name: "", args: args{s: "ab", p: ".*"}, want: true},
		{name: "", args: args{s: "aab", p: "c*a*b"}, want: true},
		{name: "", args: args{s: "mississippi", p: "mis*is*p*."}, want: false},
		{name: "", args: args{s: "ab", p: ".*c"}, want: false},
		{name: "", args: args{s: "aaa", p: "ab*a*c*a"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
