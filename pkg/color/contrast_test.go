package color

import (
	"math"
	"testing"
)

func TestRelativeLuminance(t *testing.T) {
	tests := []struct {
		color RGB
		want  float64
	}{
		{RGB{0, 0, 0}, 0.0},
		{RGB{255, 255, 255}, 1.0},
	}

	for _, tt := range tests {
		got := tt.color.RelativeLuminance()
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf("RGB(%d,%d,%d).RelativeLuminance() = %.4f, want %.4f",
				tt.color.R, tt.color.G, tt.color.B, got, tt.want)
		}
	}
}

func TestContrastRatio(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}

	ratio := ContrastRatio(black, white)
	if math.Abs(ratio-21.0) > 0.01 {
		t.Errorf("ContrastRatio(black, white) = %.2f, want 21.00", ratio)
	}

	// Same color should have ratio 1:1
	ratio = ContrastRatio(white, white)
	if math.Abs(ratio-1.0) > 0.01 {
		t.Errorf("ContrastRatio(white, white) = %.2f, want 1.00", ratio)
	}
}

func TestWCAGLevel(t *testing.T) {
	tests := []struct {
		ratio float64
		want  string
	}{
		{21.0, "AAA"},
		{7.0, "AAA"},
		{5.0, "AA"},
		{4.5, "AA"},
		{3.5, "AALarge"},
		{3.0, "AALarge"},
		{2.5, "Fail"},
	}

	for _, tt := range tests {
		got := WCAGLevel(tt.ratio)
		if got != tt.want {
			t.Errorf("WCAGLevel(%.1f) = %q, want %q", tt.ratio, got, tt.want)
		}
	}
}

func TestIsAccessible(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}

	if !IsAccessible(black, white) {
		t.Error("Black/white should be accessible")
	}

	if IsAccessible(white, white) {
		t.Error("White/white should not be accessible")
	}
}

func TestSuggestTextColor(t *testing.T) {
	// Light background should suggest black text
	lightBg := RGB{240, 240, 240}
	text := SuggestTextColor(lightBg)
	if text != (RGB{0, 0, 0}) {
		t.Errorf("SuggestTextColor(light) = %v, want black", text)
	}

	// Dark background should suggest white text
	darkBg := RGB{30, 30, 30}
	text = SuggestTextColor(darkBg)
	if text != (RGB{255, 255, 255}) {
		t.Errorf("SuggestTextColor(dark) = %v, want white", text)
	}
}

func TestBrightness(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}
	red := RGB{255, 0, 0}

	if black.Brightness() != 0 {
		t.Errorf("Black brightness = %.1f, want 0", black.Brightness())
	}
	if white.Brightness() != 255 {
		t.Errorf("White brightness = %.1f, want 255", white.Brightness())
	}
	if red.Brightness() <= 0 || red.Brightness() >= 255 {
		t.Errorf("Red brightness = %.1f, want between 0 and 255", red.Brightness())
	}
}

func TestIsLight(t *testing.T) {
	white := RGB{255, 255, 255}
	black := RGB{0, 0, 0}
	lightGray := RGB{200, 200, 200}

	if !white.IsLight() {
		t.Error("White should be light")
	}
	if black.IsLight() {
		t.Error("Black should not be light")
	}
	if !lightGray.IsLight() {
		t.Error("Light gray should be light")
	}
}

func TestIsAccessibleLarge(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}
	gray := RGB{128, 128, 128}

	// Black/white should pass large text
	if !IsAccessibleLarge(black, white) {
		t.Error("Black/white should be accessible for large text")
	}
	// Same color should fail
	if IsAccessibleLarge(gray, gray) {
		t.Error("Same color should not be accessible")
	}
	// Low contrast should fail
	lowContrast := RGB{200, 200, 200}
	if IsAccessibleLarge(gray, lowContrast) {
		t.Error("Low contrast should not be accessible for large text")
	}
}

func TestLuminance(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}
	red := RGB{255, 0, 0}

	if black.Luminance() != 0 {
		t.Errorf("Black luminance = %.4f, want 0", black.Luminance())
	}
	if math.Abs(white.Luminance()-1.0) > 0.001 {
		t.Errorf("White luminance = %.4f, want 1.0", white.Luminance())
	}
	// Red: 0.299*255/255 = 0.299
	wantRed := 0.299
	if math.Abs(red.Luminance()-wantRed) > 0.01 {
		t.Errorf("Red luminance = %.4f, want %.4f", red.Luminance(), wantRed)
	}
}
