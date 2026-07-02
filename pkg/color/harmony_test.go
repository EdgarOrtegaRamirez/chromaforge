package color

import (
	"math"
	"testing"
)

func TestComplementary(t *testing.T) {
	tests := []struct {
		input RGB
		want  RGB
	}{
		{RGB{255, 0, 0}, RGB{0, 255, 255}},   // Red -> Cyan
		{RGB{0, 255, 0}, RGB{255, 0, 255}},   // Green -> Magenta
		{RGB{0, 0, 255}, RGB{255, 255, 0}},   // Blue -> Yellow
	}

	for _, tt := range tests {
		got := tt.input.Complementary()
		// Allow small rounding differences
		if math.Abs(float64(got.R)-float64(tt.want.R)) > 1 ||
			math.Abs(float64(got.G)-float64(tt.want.G)) > 1 ||
			math.Abs(float64(got.B)-float64(tt.want.B)) > 1 {
			t.Errorf("RGB(%d,%d,%d).Complementary() = %v, want %v",
				tt.input.R, tt.input.G, tt.input.B, got, tt.want)
		}
	}
}

func TestTriadic(t *testing.T) {
	red := RGB{255, 0, 0}
	triadic := red.Triadic()
	if len(triadic) != 3 {
		t.Errorf("Triadic() returned %d colors, want 3", len(triadic))
	}
	// First color should be the original
	if triadic[0] != red {
		t.Errorf("Triadic()[0] = %v, want %v", triadic[0], red)
	}
}

func TestAnalogous(t *testing.T) {
	red := RGB{255, 0, 0}
	analogous := red.Analogous()
	if len(analogous) != 3 {
		t.Errorf("Analogous() returned %d colors, want 3", len(analogous))
	}
	// Middle color should be the original
	if analogous[1] != red {
		t.Errorf("Analogous()[1] = %v, want %v", analogous[1], red)
	}
}

func TestSplitComplementary(t *testing.T) {
	red := RGB{255, 0, 0}
	split := red.SplitComplementary()
	if len(split) != 3 {
		t.Errorf("SplitComplementary() returned %d colors, want 3", len(split))
	}
	// First color should be the original
	if split[0] != red {
		t.Errorf("SplitComplementary()[0] = %v, want %v", split[0], red)
	}
}

func TestTetradic(t *testing.T) {
	red := RGB{255, 0, 0}
	tetradic := red.Tetradic()
	if len(tetradic) != 4 {
		t.Errorf("Tetradic() returned %d colors, want 4", len(tetradic))
	}
}

func TestSquare(t *testing.T) {
	red := RGB{255, 0, 0}
	square := red.Square()
	if len(square) != 4 {
		t.Errorf("Square() returned %d colors, want 4", len(square))
	}
}

func TestMonochromatic(t *testing.T) {
	red := RGB{255, 0, 0}
	mono := red.Monochromatic()
	if len(mono) != 5 {
		t.Errorf("Monochromatic() returned %d colors, want 5", len(mono))
	}
	// All colors should have the same hue
	hue := mono[0].ToHSL().H
	for i, c := range mono {
		h := c.ToHSL().H
		if math.Abs(h-hue) > 1 {
			t.Errorf("Monochromatic()[%d] hue = %.1f, want %.1f", i, h, hue)
		}
	}
}

func TestGenerateTints(t *testing.T) {
	red := RGB{255, 0, 0}
	tints := red.GenerateTints(5)
	if len(tints) != 5 {
		t.Errorf("GenerateTints(5) returned %d colors, want 5", len(tints))
	}
	// Tints should get lighter
	for i := 1; i < len(tints); i++ {
		if tints[i].Brightness() < tints[i-1].Brightness() {
			t.Errorf("Tint %d (%.1f) should be brighter than tint %d (%.1f)",
				i, tints[i].Brightness(), i-1, tints[i-1].Brightness())
		}
	}
}

func TestGenerateShades(t *testing.T) {
	red := RGB{255, 0, 0}
	shades := red.GenerateShades(5)
	if len(shades) != 5 {
		t.Errorf("GenerateShades(5) returned %d colors, want 5", len(shades))
	}
	// Shades should get darker
	for i := 1; i < len(shades); i++ {
		if shades[i].Brightness() > shades[i-1].Brightness() {
			t.Errorf("Shade %d (%.1f) should be darker than shade %d (%.1f)",
				i, shades[i].Brightness(), i-1, shades[i-1].Brightness())
		}
	}
}

func TestGeneratePalette(t *testing.T) {
	red := RGB{255, 0, 0}
	harmonies := []string{"complementary", "analogous", "triadic", "split-complementary", "tetradic", "square", "monochromatic"}

	for _, h := range harmonies {
		palette := red.GeneratePalette(h)
		if len(palette.Colors) == 0 {
			t.Errorf("GeneratePalette(%q) returned empty palette", h)
		}
		if palette.Seed != red {
			t.Errorf("GeneratePalette(%q) seed = %v, want %v", h, palette.Seed, red)
		}
	}
}
