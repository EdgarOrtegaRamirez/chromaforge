package color

// ColorBlindnessType represents a type of color vision deficiency.
type ColorBlindnessType int

const (
	Protanopia    ColorBlindnessType = iota // Red-blind
	Deuteranopia                            // Green-blind
	Tritanopia                              // Blue-blind
	Achromatopsia                           // Total color blindness
	Protanomaly                             // Red-weak
	Deuteranomaly                           // Green-weak
	Tritanomaly                             // Blue-weak
)

// ColorBlindnessNames maps types to human-readable names.
var ColorBlindnessNames = map[ColorBlindnessType]string{
	Protanopia:    "Protanopia (red-blind)",
	Deuteranopia:  "Deuteranopia (green-blind)",
	Tritanopia:    "Tritanopia (blue-blind)",
	Achromatopsia: "Achromatopsia (total color blindness)",
	Protanomaly:   "Protanomaly (red-weak)",
	Deuteranomaly: "Deuteranomaly (green-weak)",
	Tritanomaly:   "Tritanomaly (blue-weak)",
}

// SimulateColorBlindness simulates how a color appears to someone
// with the specified type of color vision deficiency.
func (c RGB) SimulateColorBlindness(cbType ColorBlindnessType) RGB {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	var nr, ng, nb float64

	switch cbType {
	case Protanopia:
		// Red-blind transformation matrix
		nr = 0.567*r + 0.433*g + 0.0*b
		ng = 0.558*r + 0.442*g + 0.0*b
		nb = 0.0*r + 0.242*g + 0.758*b
	case Deuteranopia:
		// Green-blind transformation matrix
		nr = 0.625*r + 0.375*g + 0.0*b
		ng = 0.7*r + 0.3*g + 0.0*b
		nb = 0.0*r + 0.3*g + 0.7*b
	case Tritanopia:
		// Blue-blind transformation matrix
		nr = 0.95*r + 0.05*g + 0.0*b
		ng = 0.0*r + 0.433*g + 0.567*b
		nb = 0.0*r + 0.475*g + 0.525*b
	case Achromatopsia:
		// Total color blindness (grayscale)
		lum := 0.2126*r + 0.7152*g + 0.0722*b
		nr = lum
		ng = lum
		nb = lum
	case Protanomaly:
		nr = 0.817*r + 0.183*g + 0.0*b
		ng = 0.333*r + 0.667*g + 0.0*b
		nb = 0.0*r + 0.125*g + 0.875*b
	case Deuteranomaly:
		nr = 0.8*r + 0.2*g + 0.0*b
		ng = 0.258*r + 0.742*g + 0.0*b
		nb = 0.0*r + 0.142*g + 0.858*b
	case Tritanomaly:
		nr = 0.967*r + 0.033*g + 0.0*b
		ng = 0.0*r + 0.733*g + 0.267*b
		nb = 0.0*r + 0.183*g + 0.817*b
	default:
		return c
	}

	return RGB{
		R: clampUint8(nr),
		G: clampUint8(ng),
		B: clampUint8(nb),
	}
}

func clampUint8(v float64) uint8 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 255
	}
	return uint8(v*255 + 0.5)
}

// SimulateAll simulates all color blindness types and returns a map.
func (c RGB) SimulateAll() map[string]RGB {
	result := make(map[string]RGB)
	for cbType, name := range ColorBlindnessNames {
		result[name] = c.SimulateColorBlindness(cbType)
	}
	return result
}
