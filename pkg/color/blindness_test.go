package color

import (
	"testing"
)

func TestSimulateColorBlindness(t *testing.T) {
	red := RGB{255, 0, 0}
	green := RGB{0, 255, 0}
	blue := RGB{0, 0, 255}

	// Test all types don't panic
	for cbType := range ColorBlindnessNames {
		_ = red.SimulateColorBlindness(cbType)
		_ = green.SimulateColorBlindness(cbType)
		_ = blue.SimulateColorBlindness(cbType)
	}
}

func TestProtanopia(t *testing.T) {
	red := RGB{255, 0, 0}
	simulated := red.SimulateColorBlindness(Protanopia)

	// Red should appear different to a protanope
	if simulated == red {
		t.Error("Protanopia simulation should change red")
	}
}

func TestDeuteranopia(t *testing.T) {
	green := RGB{0, 255, 0}
	simulated := green.SimulateColorBlindness(Deuteranopia)

	// Green should appear different to a deuteranope
	if simulated == green {
		t.Error("Deuteranopia simulation should change green")
	}
}

func TestTritanopia(t *testing.T) {
	blue := RGB{0, 0, 255}
	simulated := blue.SimulateColorBlindness(Tritanopia)

	// Blue should appear different to a tritanope
	if simulated == blue {
		t.Error("Tritanopia simulation should change blue")
	}
}

func TestAchromatopsia(t *testing.T) {
	red := RGB{255, 0, 0}
	simulated := red.SimulateColorBlindness(Achromatopsia)

	// Grayscale should have equal R, G, B
	if simulated.R != simulated.G || simulated.G != simulated.B {
		t.Errorf("Achromatopsia should produce grayscale, got %v", simulated)
	}
}

func TestSimulateAll(t *testing.T) {
	red := RGB{255, 0, 0}
	all := red.SimulateAll()

	if len(all) != len(ColorBlindnessNames) {
		t.Errorf("SimulateAll() returned %d results, want %d", len(all), len(ColorBlindnessNames))
	}

	for name, color := range all {
		if color.R == 0 && color.G == 0 && color.B == 0 {
			t.Errorf("SimulateAll() color for %s is black", name)
		}
	}
}

func TestColorBlindnessNames(t *testing.T) {
	if len(ColorBlindnessNames) != 7 {
		t.Errorf("ColorBlindnessNames has %d entries, want 7", len(ColorBlindnessNames))
	}

	expectedNames := []string{
		"Protanopia (red-blind)",
		"Deuteranopia (green-blind)",
		"Tritanopia (blue-blind)",
		"Achromatopsia (total color blindness)",
		"Protanomaly (red-weak)",
		"Deuteranomaly (green-weak)",
		"Tritanomaly (blue-weak)",
	}

	for _, name := range expectedNames {
		found := false
		for _, v := range ColorBlindnessNames {
			if v == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ColorBlindnessNames missing %q", name)
		}
	}
}
