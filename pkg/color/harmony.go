package color

import "math"

// Complementary returns the complementary color (180° apart on the color wheel).
func (c RGB) Complementary() RGB {
	hsl := c.ToHSL()
	hsl.H = math.Mod(hsl.H+180, 360)
	return hsl.ToRGB()
}

// Analogous returns two analogous colors (±30° on the color wheel).
func (c RGB) Analogous() []RGB {
	hsl := c.ToHSL()
	return []RGB{
		HSL{H: math.Mod(hsl.H+330, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		c,
		HSL{H: math.Mod(hsl.H+30, 360), S: hsl.S, L: hsl.L}.ToRGB(),
	}
}

// Triadic returns two triadic colors (120° apart on the color wheel).
func (c RGB) Triadic() []RGB {
	hsl := c.ToHSL()
	return []RGB{
		c,
		HSL{H: math.Mod(hsl.H+120, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+240, 360), S: hsl.S, L: hsl.L}.ToRGB(),
	}
}

// SplitComplementary returns two split-complementary colors.
func (c RGB) SplitComplementary() []RGB {
	hsl := c.ToHSL()
	return []RGB{
		c,
		HSL{H: math.Mod(hsl.H+150, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+210, 360), S: hsl.S, L: hsl.L}.ToRGB(),
	}
}

// Tetradic (Rectangular) returns three tetradic colors.
func (c RGB) Tetradic() []RGB {
	hsl := c.ToHSL()
	return []RGB{
		c,
		HSL{H: math.Mod(hsl.H+60, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+180, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+240, 360), S: hsl.S, L: hsl.L}.ToRGB(),
	}
}

// Square returns three square colors (90° apart on the color wheel).
func (c RGB) Square() []RGB {
	hsl := c.ToHSL()
	return []RGB{
		c,
		HSL{H: math.Mod(hsl.H+90, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+180, 360), S: hsl.S, L: hsl.L}.ToRGB(),
		HSL{H: math.Mod(hsl.H+270, 360), S: hsl.S, L: hsl.L}.ToRGB(),
	}
}

// Monochromatic returns a monochromatic palette with 5 shades.
func (c RGB) Monochromatic() []RGB {
	hsl := c.ToHSL()
	var palette []RGB
	steps := 5
	for i := 0; i < steps; i++ {
		l := float64(i) / float64(steps-1) // 0 to 1
		l = 10 + l*80                       // 10% to 90%
		palette = append(palette, HSL{H: hsl.H, S: hsl.S, L: l}.ToRGB())
	}
	return palette
}

// GeneratePalette generates a color palette based on the harmony type.
func (c RGB) GeneratePalette(harmony string) Palette {
	var colors []RGB

	switch harmony {
	case "complementary":
		colors = []RGB{c, c.Complementary()}
	case "analogous":
		colors = c.Analogous()
	case "triadic":
		colors = c.Triadic()
	case "split-complementary", "split":
		colors = c.SplitComplementary()
	case "tetradic", "rectangular":
		colors = c.Tetradic()
	case "square":
		colors = c.Square()
	case "monochromatic", "mono":
		colors = c.Monochromatic()
	default:
		colors = []RGB{c}
	}

	return Palette{
		Name:    harmony,
		Colors:  colors,
		Seed:    c,
		Harmony: harmony,
	}
}

// GenerateTints generates lighter tints of a color.
func (c RGB) GenerateTints(steps int) []RGB {
	if steps < 1 {
		steps = 5
	}
	hsl := c.ToHSL()
	var tints []RGB
	for i := 0; i < steps; i++ {
		factor := float64(i) / float64(steps)
		l := hsl.L + (100-hsl.L)*factor
		tints = append(tints, HSL{H: hsl.H, S: hsl.S, L: l}.ToRGB())
	}
	return tints
}

// GenerateShades generates darker shades of a color.
func (c RGB) GenerateShades(steps int) []RGB {
	if steps < 1 {
		steps = 5
	}
	hsl := c.ToHSL()
	var shades []RGB
	for i := 0; i < steps; i++ {
		factor := float64(i) / float64(steps)
		l := hsl.L * (1 - factor)
		shades = append(shades, HSL{H: hsl.H, S: hsl.S, L: l}.ToRGB())
	}
	return shades
}
