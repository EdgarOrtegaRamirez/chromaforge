package color

import (
	"math"
	"testing"
)

func TestRGBToHSL(t *testing.T) {
	tests := []struct {
		input RGB
		want  HSL
	}{
		{RGB{255, 0, 0}, HSL{0, 100, 50}},
		{RGB{0, 255, 0}, HSL{120, 100, 50}},
		{RGB{0, 0, 255}, HSL{240, 100, 50}},
		{RGB{128, 128, 128}, HSL{0, 0, 50.2}},
		{RGB{0, 0, 0}, HSL{0, 0, 0}},
		{RGB{255, 255, 255}, HSL{0, 0, 100}},
	}

	for _, tt := range tests {
		got := tt.input.ToHSL()
		if math.Abs(got.H-tt.want.H) > 0.1 || math.Abs(got.S-tt.want.S) > 0.1 || math.Abs(got.L-tt.want.L) > 0.1 {
			t.Errorf("RGB(%d,%d,%d).ToHSL() = HSL(%.1f,%.1f,%.1f), want HSL(%.1f,%.1f,%.1f)",
				tt.input.R, tt.input.G, tt.input.B, got.H, got.S, got.L, tt.want.H, tt.want.S, tt.want.L)
		}
	}
}

func TestHSLToRGB(t *testing.T) {
	tests := []struct {
		input HSL
		want  RGB
	}{
		{HSL{0, 100, 50}, RGB{255, 0, 0}},
		{HSL{120, 100, 50}, RGB{0, 255, 0}},
		{HSL{240, 100, 50}, RGB{0, 0, 255}},
		{HSL{0, 0, 0}, RGB{0, 0, 0}},
		{HSL{0, 0, 100}, RGB{255, 255, 255}},
	}

	for _, tt := range tests {
		got := tt.input.ToRGB()
		if got != tt.want {
			t.Errorf("HSL(%.0f,%.0f,%.0f).ToRGB() = %v, want %v",
				tt.input.H, tt.input.S, tt.input.L, got, tt.want)
		}
	}
}

func TestRGBToHSLRoundTrip(t *testing.T) {
	colors := []RGB{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{128, 64, 192},
		{200, 150, 100},
	}

	for _, c := range colors {
		hsl := c.ToHSL()
		roundTrip := hsl.ToRGB()
		if roundTrip != c {
			t.Errorf("Round trip failed for %v: HSL(%v) -> %v", c, hsl, roundTrip)
		}
	}
}

func TestRGBToHSV(t *testing.T) {
	tests := []struct {
		input RGB
		want  HSV
	}{
		{RGB{255, 0, 0}, HSV{0, 100, 100}},
		{RGB{0, 255, 0}, HSV{120, 100, 100}},
		{RGB{0, 0, 255}, HSV{240, 100, 100}},
	}

	for _, tt := range tests {
		got := tt.input.ToHSV()
		if math.Abs(got.H-tt.want.H) > 0.1 || math.Abs(got.S-tt.want.S) > 0.1 || math.Abs(got.V-tt.want.V) > 0.1 {
			t.Errorf("RGB(%d,%d,%d).ToHSV() = HSV(%.1f,%.1f,%.1f), want HSV(%.1f,%.1f,%.1f)",
				tt.input.R, tt.input.G, tt.input.B, got.H, got.S, got.V, tt.want.H, tt.want.S, tt.want.V)
		}
	}
}

func TestRGBToCMYK(t *testing.T) {
	tests := []struct {
		input RGB
		want  CMYK
	}{
		{RGB{255, 0, 0}, CMYK{0, 100, 100, 0}},
		{RGB{0, 255, 0}, CMYK{100, 0, 100, 0}},
		{RGB{0, 0, 255}, CMYK{100, 100, 0, 0}},
		{RGB{0, 0, 0}, CMYK{0, 0, 0, 100}},
	}

	for _, tt := range tests {
		got := tt.input.ToCMYK()
		if math.Abs(got.C-tt.want.C) > 0.1 || math.Abs(got.M-tt.want.M) > 0.1 ||
			math.Abs(got.Y-tt.want.Y) > 0.1 || math.Abs(got.K-tt.want.K) > 0.1 {
			t.Errorf("RGB(%d,%d,%d).ToCMYK() = CMYK(%.1f,%.1f,%.1f,%.1f), want CMYK(%.1f,%.1f,%.1f,%.1f)",
				tt.input.R, tt.input.G, tt.input.B, got.C, got.M, got.Y, got.K, tt.want.C, tt.want.M, tt.want.Y, tt.want.K)
		}
	}
}

func TestRGBToLab(t *testing.T) {
	// Test that known colors produce reasonable Lab values
	tests := []struct {
		input RGB
		name  string
	}{
		{RGB{255, 0, 0}, "red"},
		{RGB{0, 255, 0}, "green"},
		{RGB{0, 0, 255}, "blue"},
		{RGB{255, 255, 255}, "white"},
		{RGB{0, 0, 0}, "black"},
	}

	for _, tt := range tests {
		got := tt.input.ToLab()
		// White should have L≈100, a≈0, b≈0
		// Black should have L≈0, a≈0, b≈0
		// Red should have positive a (red-green axis)
		if got.L < 0 || got.L > 100.1 {
			t.Errorf("RGB(%d,%d,%d) (%s) Lab.L = %.2f, want [0, 100]",
				tt.input.R, tt.input.G, tt.input.B, tt.name, got.L)
		}
	}
}

func TestLabRoundTrip(t *testing.T) {
	colors := []RGB{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{128, 64, 192},
	}

	for _, c := range colors {
		lab := c.ToLab()
		roundTrip := lab.ToRGB()
		// Lab round trip may not be exact due to gamma
		if math.Abs(float64(roundTrip.R)-float64(c.R)) > 2 ||
			math.Abs(float64(roundTrip.G)-float64(c.G)) > 2 ||
			math.Abs(float64(roundTrip.B)-float64(c.B)) > 2 {
			t.Errorf("Lab round trip failed for %v: Lab(%v) -> %v", c, lab, roundTrip)
		}
	}
}
