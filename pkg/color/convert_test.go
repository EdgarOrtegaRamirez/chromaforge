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

func TestHSVToRGB(t *testing.T) {
	tests := []struct {
		input HSV
		want  RGB
	}{
		{HSV{0, 100, 100}, RGB{255, 0, 0}},
		{HSV{120, 100, 100}, RGB{0, 255, 0}},
		{HSV{240, 100, 100}, RGB{0, 0, 255}},
		{HSV{0, 0, 0}, RGB{0, 0, 0}},
		{HSV{0, 0, 100}, RGB{255, 255, 255}},
		{HSV{60, 100, 100}, RGB{255, 255, 0}},  // yellow
		{HSV{180, 100, 100}, RGB{0, 255, 255}}, // cyan
		{HSV{300, 100, 100}, RGB{255, 0, 255}}, // magenta
	}

	for _, tt := range tests {
		got := tt.input.ToRGB()
		if got != tt.want {
			t.Errorf("HSV(%.0f,%.0f,%.0f).ToRGB() = %v, want %v",
				tt.input.H, tt.input.S, tt.input.V, got, tt.want)
		}
	}
}

func TestCMYKToRGB(t *testing.T) {
	tests := []struct {
		input CMYK
		want  RGB
	}{
		{CMYK{0, 100, 100, 0}, RGB{255, 0, 0}},
		{CMYK{100, 0, 100, 0}, RGB{0, 255, 0}},
		{CMYK{100, 100, 0, 0}, RGB{0, 0, 255}},
		{CMYK{0, 0, 0, 100}, RGB{0, 0, 0}},
		{CMYK{0, 0, 0, 0}, RGB{255, 255, 255}},
	}

	for _, tt := range tests {
		got := tt.input.ToRGB()
		if got != tt.want {
			t.Errorf("CMYK(%.0f,%.0f,%.0f,%.0f).ToRGB() = %v, want %v",
				tt.input.C, tt.input.M, tt.input.Y, tt.input.K, got, tt.want)
		}
	}
}

func TestXYZToRGB(t *testing.T) {
	// D65 white point
	white := XYZ{X: 0.95047, Y: 1.0, Z: 1.08883}
	got := white.ToRGB()
	if got != (RGB{255, 255, 255}) {
		t.Errorf("D65 white XYZ -> RGB = %v, want {255,255,255}", got)
	}
}

func TestCIEDE2000(t *testing.T) {
	white := RGB{255, 255, 255}.ToLab()
	black := RGB{0, 0, 0}.ToLab()

	// White and black should have maximum difference
	diff := white.CIEDE2000(black)
	if diff < 99 || diff > 100.01 {
		t.Errorf("White/Black CIEDE2000 = %.2f, want ~100", diff)
	}

	// Same color should have 0 difference
	same := RGB{128, 64, 192}.ToLab()
	diff = same.CIEDE2000(same)
	if diff > 0.01 {
		t.Errorf("Same color CIEDE2000 = %.4f, want 0", diff)
	}

	// Red and green should be different
	red := RGB{255, 0, 0}.ToLab()
	green := RGB{0, 255, 0}.ToLab()
	diff = red.CIEDE2000(green)
	if diff < 50 {
		t.Errorf("Red/Green CIEDE2000 = %.2f, want >50", diff)
	}
}

func TestClampf(t *testing.T) {
	tests := []struct {
		input  float64
		expect float64
	}{
		{-1.0, 0.0},
		{-0.5, 0.0},
		{0.0, 0.0},
		{0.5, 0.5},
		{1.0, 1.0},
		{1.5, 1.0},
		{100.0, 1.0},
		{-100.0, 0.0},
	}

	for _, tt := range tests {
		got := Clampf(tt.input)
		if math.Abs(got-tt.expect) > 0.001 {
			t.Errorf("Clampf(%.2f) = %.2f, want %.2f", tt.input, got, tt.expect)
		}
	}
}
