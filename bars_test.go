package main

import "testing"

func TestRemoveInvalidChars(t *testing.T) {
	type args struct {
		text  string
		valid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty text", args{"", "12345"}, ""},
		{"empty valid", args{"123", ""}, ""},
		{"all valid", args{"123", "12345"}, "123"},
		{"remove one invalid", args{"1_234€", "12345"}, "1234"},
		{"remove all invalid", args{"Hugo", "12345"}, ""}, // test of test - result should be ""
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInvalidChars(tt.args.text, tt.args.valid); got != tt.want {
				t.Errorf("RemoveInvalidChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
