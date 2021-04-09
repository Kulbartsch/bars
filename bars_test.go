package main

import "testing"

func TestNumberCharsLength(t *testing.T) {
	type args struct {
		text      string
		numChars  string
		fromRight bool
	}
	tests := []struct {
		name      string
		args      args
		wantStart int
		wantEnd   int
		wantLeng   int
	}{
		{name: "number text", args: {text: "123abc", numChars: "123", fromRight: false}, wantEnd: 3, wantLeng: 3},
		{"number", {"123.4","1234.",false}, 0, 5, 4},
		{"text number right", {"abc 123,4","1234,.",true}, 4, 8, 8},
		{"text number left", {"abc 1234","1234,.",false}, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, gotLen := NumberCharsLength(tt.args.text, tt.args.numChars, tt.args.fromRight)
			if gotStart != tt.wantStart {
				t.Errorf("NumberCharsLength() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("NumberCharsLength() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
			if gotLen != tt.wantLeng {
				t.Errorf("NumberCharsLength() gotLen = %v, want %v", gotLen, tt.wantLeng)
			}
		})
	}
}
