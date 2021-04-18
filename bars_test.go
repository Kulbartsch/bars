package main

import "testing"

/*func TestNumberCharsLength(t *testing.T) {
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
		wantLeng  int
	}{
		{name: "number+text", args: args{text: "123abc", numChars: "123", fromRight: false}, wantEnd: 3, wantLeng: 3},
		{"number with dot only", args{"123.4", "1234.", false}, 0, 5, 5},
		{"text number right", args{"abc 123,4", "1234,.", true}, 4, 9, 5},
		{"text number left", args{"abc 1234", "1234,.", false}, 0, 0, 0},
		{"number only", args{"1234", "1234567890.", true}, 0, 4, 4},
		{"empty", args{"", "1234567890.", false}, 0, 0, 0},
		{"only text", args{"only text", "1234567890.", false}, 0, 0, 0},
		{"number+text with umlaute", args{"1234 färß", "1234567890.", false}, 0, 4, 4},
		{"text with umlaute + number", args{"färß1234", "1234567890.", true}, 6, 9, 4},
		{"text + number with €", args{"Möt 12,34€", "1234567890.,€", true}, 5, 12, 6},
		{"number with € + text with akzent", args{"12,34€ Avéng", "1234567890.,€", false}, 0, 8, 8},

	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, gotLen := NumberCharsLength(tt.args.text, tt.args.numChars, tt.args.fromRight)
			if gotStart != tt.wantStart {
				t.Errorf("%v - NumberCharsLength() gotStart = %v, want %v", i, gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("%v - NumberCharsLength() gotEnd = %v, want %v", i, gotEnd, tt.wantEnd)
			}
			if gotLen != tt.wantLeng {
				t.Errorf("%v - NumberCharsLength() gotLen = %v, want %v", i, gotLen, tt.wantLeng)
			}
		})
	}
}*/

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
		{ "remove one invalid", args{"1_234€", "12345"}, "1234"},
		{ "remove all invalid", args{"Hugo", "12345"}, ""}, // test of test - result should be ""
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInvalidChars(tt.args.text, tt.args.valid); got != tt.want {
				t.Errorf("RemoveInvalidChars() = %v, want %v", got, tt.want)
			}
		})
	}
}