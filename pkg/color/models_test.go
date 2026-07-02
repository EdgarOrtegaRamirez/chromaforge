package color

import (
	"testing"
)

func TestParseHex(t *testing.T) {
	tests := []struct {
		input string
		want  RGB
		err   bool
	}{
		{"#FF0000", RGB{255, 0, 0}, false},
		{"FF0000", RGB{255, 0, 0}, false},
		{"#F00", RGB{255, 0, 0}, false},
		{"#00FF00", RGB{0, 255, 0}, false},
		{"#0000FF", RGB{0, 0, 255}, false},
		{"#3498DB", RGB{52, 152, 219}, false},
		{"0x1A2B3C", RGB{26, 43, 60}, false},
		{"#GGG", RGB{}, true},
		{"#12", RGB{}, true},
		{"", RGB{}, true},
	}

	for _, tt := range tests {
		got, err := ParseHex(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("ParseHex(%q) error = %v, wantErr %v", tt.input, err, tt.err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseHex(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseColor(t *testing.T) {
	tests := []struct {
		input string
		want  RGB
		err   bool
	}{
		{"#FF0000", RGB{255, 0, 0}, false},
		{"rgb(255, 0, 0)", RGB{255, 0, 0}, false},
		{"hsl(0, 100%, 50%)", RGB{255, 0, 0}, false},
		{"hsl(120, 100%, 50%)", RGB{0, 255, 0}, false},
		{"hsl(240, 100%, 50%)", RGB{0, 0, 255}, false},
		{"invalid", RGB{}, true},
		{"rgb(256, 0, 0)", RGB{}, true},
	}

	for _, tt := range tests {
		got, err := ParseColor(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("ParseColor(%q) error = %v, wantErr %v", tt.input, err, tt.err)
			continue
		}
		if !tt.err && got != tt.want {
			t.Errorf("ParseColor(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestRGBString(t *testing.T) {
	tests := []struct {
		color RGB
		want  string
	}{
		{RGB{255, 0, 0}, "#FF0000"},
		{RGB{0, 255, 0}, "#00FF00"},
		{RGB{0, 0, 255}, "#0000FF"},
		{RGB{0, 0, 0}, "#000000"},
		{RGB{255, 255, 255}, "#FFFFFF"},
	}

	for _, tt := range tests {
		got := tt.color.String()
		if got != tt.want {
			t.Errorf("RGB.String() = %q, want %q", got, tt.want)
		}
	}
}
