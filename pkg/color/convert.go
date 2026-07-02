package color

import "math"

// ToHSL converts RGB to HSL.
func (c RGB) ToHSL() HSL {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l := (max + min) / 2.0

	if max == min {
		return HSL{H: 0, S: 0, L: l * 100}
	}

	d := max - min
	var s float64
	if l > 0.5 {
		s = d / (2.0 - max - min)
	} else {
		s = d / (max + min)
	}

	var h float64
	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}
	h /= 6.0

	return HSL{H: h * 360, S: s * 100, L: l * 100}
}

// ToRGB converts HSL to RGB.
func (c HSL) ToRGB() RGB {
	h := c.H / 360.0
	s := c.S / 100.0
	l := c.L / 100.0

	if s == 0 {
		v := uint8(math.Round(l * 255))
		return RGB{v, v, v}
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r := hueToRGB(p, q, h+1.0/3.0)
	g := hueToRGB(p, q, h)
	b := hueToRGB(p, q, h-1.0/3.0)

	return RGB{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
	}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// ToHSV converts RGB to HSV.
func (c RGB) ToHSV() HSV {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	d := max - min

	v := max
	var s float64
	if max != 0 {
		s = d / max
	}

	var h float64
	if d == 0 {
		h = 0
	} else {
		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6.0
	}

	return HSV{H: h * 360, S: s * 100, V: v * 100}
}

// ToRGB converts HSV to RGB.
func (c HSV) ToRGB() RGB {
	h := c.H / 360.0
	s := c.S / 100.0
	v := c.V / 100.0

	if s == 0 {
		val := uint8(math.Round(v * 255))
		return RGB{val, val, val}
	}

	i := math.Floor(h * 6)
	f := h*6 - i
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	var r, g, b float64
	switch int(i) % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return RGB{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
	}
}

// ToCMYK converts RGB to CMYK.
func (c RGB) ToCMYK() CMYK {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	k := 1.0 - math.Max(r, math.Max(g, b))
	if k == 1 {
		return CMYK{K: 100}
	}

	cyan := (1 - r - k) / (1 - k)
	mag := (1 - g - k) / (1 - k)
	yel := (1 - b - k) / (1 - k)

	return CMYK{
		C: RoundTo(cyan*100, 2),
		M: RoundTo(mag*100, 2),
		Y: RoundTo(yel*100, 2),
		K: RoundTo(k*100, 2),
	}
}

// ToRGB converts CMYK to RGB.
func (c CMYK) ToRGB() RGB {
	cyan := c.C / 100.0
	mag := c.M / 100.0
	yel := c.Y / 100.0
	k := c.K / 100.0

	r := 255 * (1 - cyan) * (1 - k)
	g := 255 * (1 - mag) * (1 - k)
	b := 255 * (1 - yel) * (1 - k)

	return RGB{
		R: uint8(math.Round(r)),
		G: uint8(math.Round(g)),
		B: uint8(math.Round(b)),
	}
}

// ToXYZ converts RGB to CIE XYZ (D65 illuminant).
func (c RGB) ToXYZ() XYZ {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	// Apply gamma correction
	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r /= 12.92
	}
	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g /= 12.92
	}
	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b /= 12.92
	}

	// sRGB to XYZ (D65)
	x := r*0.4124564 + g*0.3575761 + b*0.1804375
	y := r*0.2126729 + g*0.7151522 + b*0.0721750
	z := r*0.0193339 + g*0.1191920 + b*0.9503041

	return XYZ{X: x, Y: y, Z: z}
}

// ToRGB converts CIE XYZ to RGB.
func (c XYZ) ToRGB() RGB {
	// XYZ to linear RGB
	r := c.X*3.2404542 + c.Y*-1.5371385 + c.Z*-0.4985314
	g := c.X*-0.9692660 + c.Y*1.8760108 + c.Z*0.0415560
	b := c.X*0.0556434 + c.Y*-0.2040259 + c.Z*1.0572252

	// Apply inverse gamma correction
	if r > 0.0031308 {
		r = 1.055*math.Pow(r, 1.0/2.4) - 0.055
	} else {
		r *= 12.92
	}
	if g > 0.0031308 {
		g = 1.055*math.Pow(g, 1.0/2.4) - 0.055
	} else {
		g *= 12.92
	}
	if b > 0.0031308 {
		b = 1.055*math.Pow(b, 1.0/2.4) - 0.055
	} else {
		b *= 12.92
	}

	return RGB{
		R: uint8(math.Round(math.Max(0, math.Min(1, r)) * 255)),
		G: uint8(math.Round(math.Max(0, math.Min(1, g)) * 255)),
		B: uint8(math.Round(math.Max(0, math.Min(1, b)) * 255)),
	}
}

// ToLab converts RGB to CIELAB via XYZ.
func (c RGB) ToLab() Lab {
	xyz := c.ToXYZ()
	return xyz.ToLab()
}

// ToLab converts CIE XYZ to CIELAB.
func (c XYZ) ToLab() Lab {
	// D65 reference white point
	xn := 0.95047
	yn := 1.0
	zn := 1.08883

	f := func(t float64) float64 {
		d := 6.0 / 29.0
		if t > d*d*d {
			return math.Cbrt(t)
		}
		return t/(3*d*d) + 4.0/29.0
	}

	fx := f(c.X / xn)
	fy := f(c.Y / yn)
	fz := f(c.Z / zn)

	l := 116*fy - 16
	a := 500 * (fx - fy)
	b := 200 * (fy - fz)

	return Lab{L: l, A: a, B: b}
}

// ToRGB converts CIELAB to RGB via XYZ.
func (c Lab) ToRGB() RGB {
	xyz := c.ToXYZ()
	return xyz.ToRGB()
}

// ToXYZ converts CIELAB to CIE XYZ.
func (c Lab) ToXYZ() XYZ {
	// D65 reference white point
	xn := 0.95047
	yn := 1.0
	zn := 1.08883

	fy := (c.L + 16) / 116
	fx := c.A/500 + fy
	fz := fy - c.B/200

	finv := func(t float64) float64 {
		d := 6.0 / 29.0
		if t > d {
			return t * t * t
		}
		return 3 * d * d * (t - 4.0/29.0)
	}

	return XYZ{
		X: xn * finv(fx),
		Y: yn * finv(fy),
		Z: zn * finv(fz),
	}
}

// CIEDE2000 calculates the color difference between two Lab colors
// using the CIEDE2000 formula.
func (c1 Lab) CIEDE2000(c2 Lab) float64 {
	// Simplified CIEDE2000 implementation
	l1, a1, b1 := c1.L, c1.A, c1.B
	l2, a2, b2 := c2.L, c2.A, c2.B

	// Calculate Cab and Hab
	c1ab := math.Sqrt(a1*a1 + b1*b1)
	c2ab := math.Sqrt(a2*a2 + b2*b2)
	cabAvg := (c1ab + c2ab) / 2
	cabAvg7 := math.Pow(cabAvg, 7)
	g := 0.5 * (1 - math.Sqrt(cabAvg7/(cabAvg7+math.Pow(25, 7))))

	a1p := a1 * (1 + g)
	a2p := a2 * (1 + g)
	c1p := math.Sqrt(a1p*a1p + b1*b1)
	c2p := math.Sqrt(a2p*a2p + b2*b2)

	var h1p, h2p float64
	if b1 == 0 && a1p == 0 {
		h1p = 0
	} else {
		h1p = math.Atan2(b1, a1p) * 180 / math.Pi
		if h1p < 0 {
			h1p += 360
		}
	}
	if b2 == 0 && a2p == 0 {
		h2p = 0
	} else {
		h2p = math.Atan2(b2, a2p) * 180 / math.Pi
		if h2p < 0 {
			h2p += 360
		}
	}

	dlp := l2 - l1
	dcp := c2p - c1p

	var dhp float64
	if c1p*c2p == 0 {
		dhp = 0
	} else if math.Abs(h2p-h1p) <= 180 {
		dhp = h2p - h1p
	} else if h2p-h1p > 180 {
		dhp = h2p - h1p - 360
	} else {
		dhp = h2p - h1p + 360
	}
	dhp = 2 * math.Sqrt(c1p*c2p) * math.Sin(dhp*math.Pi/360)

	lpAvg := (l1 + l2) / 2
	cpAvg := (c1p + c2p) / 2

	var hpAvg float64
	if c1p*c2p == 0 {
		hpAvg = h1p + h2p
	} else if math.Abs(h1p-h2p) <= 180 {
		hpAvg = (h1p + h2p) / 2
	} else if h1p+h2p < 360 {
		hpAvg = (h1p + h2p + 360) / 2
	} else {
		hpAvg = (h1p + h2p - 360) / 2
	}

	t := 1 - 0.17*math.Cos((hpAvg-30)*math.Pi/180) + 0.24*math.Cos(2*hpAvg*math.Pi/180) +
		0.32*math.Cos((3*hpAvg+6)*math.Pi/180) - 0.20*math.Cos((4*hpAvg-63)*math.Pi/180)

	sl := 1 + 0.015*math.Pow(lpAvg-50, 2)/math.Sqrt(20+math.Pow(lpAvg-50, 2))
	sc := 1 + 0.045*cpAvg
	sh := 1 + 0.015*cpAvg*t

	lt := math.Pow(lpAvg-50, 2) / (20 + math.Pow(lpAvg-50, 2))
	rc := 2 * math.Sqrt(math.Pow(cpAvg, 7)/(math.Pow(cpAvg, 7)+math.Pow(25, 7)))
	_ = -math.Sin(2*lt*rc*math.Pi/180) * math.Sqrt(math.Abs(rc))

	kl := 1.0
	kc := 1.0
	kh := 1.0

	dl := dlp / (kl * sl)
	dc := dcp / (kc * sc)
	dh := dhp / (kh * sh)

	return math.Sqrt(dl*dl + dc*dc + dh*dh)
}
