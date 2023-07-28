package checks

import "testing"

func TestUsername(t *testing.T) {

}

func Test_f(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"qwertyuuio", true},
		{"123456789", true},
		{"asdqwe5451232", true},
		{"1", false},
		{"", false},
		{"1564165....", true},
		{"!!!!....", true},
		{"@@@@@@aaa...", true},
		{"###%%%%...", true},
		{"+-()==", false},
		{"]][[", false},
		{"}}{{", false},
		{"....", true},
		{"++++", true},
		{"----", true},
		{"~~~~", true},
		{"~~~", false},
		{":::'''\"\"|||\\\\///", false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := f(tt.args); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
