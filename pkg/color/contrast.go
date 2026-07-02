package color

import "math"

// RelativeLuminance calculates the relative luminance of an RGB color
// per WCAG 2.1 specification.
func (c RGB) RelativeLuminance() float64 {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	// Apply sRGB gamma correction
	if r <= 0.03928 {
		r /= 12.92
	} else {
		r = math.Pow((r+0.055)/1.055, 2.4)
	}
	if g <= 0.03928 {
		g /= 12.92
	} else {
		g = math.Pow((g+0.055)/1.055, 2.4)
	}
	if b <= 0.03928 {
		b /= 12.92
	} else {
		b = math.Pow((b+0.055)/1.055, 2.4)
	}

	return 0.2126*r + 0.7152*g + 0.0722*b
}

// ContrastRatio calculates the WCAG contrast ratio between two colors.
// Returns a value between 1 and 21.
func ContrastRatio(c1, c2 RGB) float64 {
	l1 := c1.RelativeLuminance()
	l2 := c2.RelativeLuminance()

	if l1 < l2 {
		l1, l2 = l2, l1
	}

	return (l1 + 0.05) / (l2 + 0.05)
}

// WCAGLevel returns the WCAG compliance level for a contrast ratio.
// Returns "AAA", "AA", "AALarge", or "Fail".
func WCAGLevel(ratio float64) string {
	if ratio >= 7.0 {
		return "AAA"
	}
	if ratio >= 4.5 {
		return "AA"
	}
	if ratio >= 3.0 {
		return "AALarge"
	}
	return "Fail"
}

// IsAccessible checks if two colors meet WCAG AA requirements
// for normal text (4.5:1 ratio).
func IsAccessible(c1, c2 RGB) bool {
	return ContrastRatio(c1, c2) >= 4.5
}

// IsAccessibleLarge checks if two colors meet WCAG AA requirements
// for large text (3:1 ratio).
func IsAccessibleLarge(c1, c2 RGB) bool {
	return ContrastRatio(c1, c2) >= 3.0
}

// SuggestTextColor suggests black or white text for optimal contrast
// against the given background color.
func SuggestTextColor(bg RGB) RGB {
	lum := bg.RelativeLuminance()
	if lum > 0.179 {
		return RGB{0, 0, 0} // Black text for light backgrounds
	}
	return RGB{255, 255, 255} // White text for dark backgrounds
}

// Brightness returns the perceived brightness of a color (0-255).
func (c RGB) Brightness() float64 {
	return 0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)
}

// IsLight returns true if the color is considered "light".
func (c RGB) IsLight() bool {
	return c.Brightness() > 128
}

// Luminance is a simpler (non-WCAG) brightness calculation.
func (c RGB) Luminance() float64 {
	return (0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)) / 255.0
}
