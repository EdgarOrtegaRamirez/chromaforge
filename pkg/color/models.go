// Package color provides color models, conversions, and palette generation.
package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RGB represents a color in the RGB color space.
// Values are in range [0, 255].
type RGB struct {
	R, G, B uint8
}

// HSL represents a color in the HSL color space.
// H is in [0, 360), S and L are in [0, 100].
type HSL struct {
	H, S, L float64
}

// HSV represents a color in the HSV (HSB) color space.
// H is in [0, 360), S and V are in [0, 100].
type HSV struct {
	H, S, V float64
}

// CMYK represents a color in the CMYK color space.
// All values are in [0, 100].
type CMYK struct {
	C, M, Y, K float64
}

// Lab represents a color in the CIELAB color space.
// L is in [0, 100], a and b are in approximately [-128, 127].
type Lab struct {
	L, A, B float64
}

// XYZ represents a color in the CIE XYZ color space.
type XYZ struct {
	X, Y, Z float64
}

// Palette represents a collection of colors with a name.
type Palette struct {
	Name    string
	Colors  []RGB
	Seed    RGB
	Harmony string
}

// String returns the hex representation of an RGB color.
func (c RGB) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// ParseHex parses a hex color string like "#FF0000" or "FF0000".
func ParseHex(s string) (RGB, error) {
	s = strings.TrimPrefix(s, "#")
	s = strings.TrimPrefix(s, "0x")

	// Support 3-character hex (e.g., "F00" -> "FF0000")
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}

	if len(s) != 6 {
		return RGB{}, fmt.Errorf("invalid hex color: %s", s)
	}

	r, err := strconv.ParseUint(s[0:2], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid hex color: %s", s)
	}
	g, err := strconv.ParseUint(s[2:4], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid hex color: %s", s)
	}
	b, err := strconv.ParseUint(s[4:6], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid hex color: %s", s)
	}

	return RGB{uint8(r), uint8(g), uint8(b)}, nil
}

// ParseColor parses a color string in various formats:
// - Hex: "#FF0000", "FF0000", "#F00"
// - RGB: "rgb(255, 0, 0)"
// - HSL: "hsl(0, 100%, 50%)"
func ParseColor(s string) (RGB, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	// Hex format
	if strings.HasPrefix(s, "#") || strings.HasPrefix(s, "0x") || len(s) == 6 {
		return ParseHex(s)
	}

	// RGB format
	if strings.HasPrefix(s, "rgb(") && strings.HasSuffix(s, ")") {
		inner := s[4 : len(s)-1]
		parts := strings.Split(inner, ",")
		if len(parts) != 3 {
			return RGB{}, fmt.Errorf("invalid rgb format: %s", s)
		}
		r, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 8)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid rgb format: %s", s)
		}
		g, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 8)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid rgb format: %s", s)
		}
		b, err := strconv.ParseUint(strings.TrimSpace(parts[2]), 10, 8)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid rgb format: %s", s)
		}
		return RGB{uint8(r), uint8(g), uint8(b)}, nil
	}

	// HSL format
	if strings.HasPrefix(s, "hsl(") && strings.HasSuffix(s, ")") {
		inner := s[4 : len(s)-1]
		parts := strings.Split(inner, ",")
		if len(parts) != 3 {
			return RGB{}, fmt.Errorf("invalid hsl format: %s", s)
		}
		h, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(parts[0], "deg")), 64)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid hsl format: %s", s)
		}
		sVal, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(parts[1], "%")), 64)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid hsl format: %s", s)
		}
		l, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(parts[2], "%")), 64)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid hsl format: %s", s)
		}
		return HSL{H: h, S: sVal, L: l}.ToRGB(), nil
	}

	return RGB{}, fmt.Errorf("unsupported color format: %s", s)
}

// Clampf clamps a float64 to [0, 1].
func Clampf(v float64) float64 {
	return math.Max(0, math.Min(1, v))
}

// RoundTo rounds a float64 to the given number of decimal places.
func RoundTo(v float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(v*factor) / factor
}
