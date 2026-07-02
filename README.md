# ChromaForge

A color palette and scheme management CLI tool and Go library.

## Features

- **Color Format Conversion** — Convert between RGB, HSL, HSV, CMYK, Lab, and hex
- **Color Harmony Generation** — Complementary, triadic, analogous, split-complementary, tetradic, square, and monochromatic palettes
- **WCAG Contrast Checking** — Calculate contrast ratios and check accessibility compliance (AA, AAA)
- **Color Blindness Simulation** — Simulate 7 types of color vision deficiency
- **Tint & Shade Generation** — Generate lighter tints and darker shades of any color
- **Multi-Format Export** — Export palettes as CSS variables, Tailwind config, JSON, SVG, and SCSS
- **CIEDE2000 Color Difference** — Calculate perceptual color difference using the CIEDE2000 formula

## Quick Start

```bash
# Install
go install github.com/EdgarOrtegaRamirez/chromaforge/cmd/chromaforge@latest

# Or build from source
git clone https://github.com/EdgarOrtegaRamirez/chromaforge
cd chromaforge
go build ./cmd/chromaforge/
```

## Usage

### Convert Colors

```bash
# Hex to HSL
chromaforge convert "#3498db" hsl
# Output: hsl(204.1, 69.9%, 53.1%)

# RGB to hex
chromaforge convert "rgb(255, 100, 0)" hex
# Output: #FF6400

# Any format to Lab
chromaforge convert "#3498db" lab
# Output: lab(60.16, -6.10, -42.23)
```

### Generate Color Harmonies

```bash
# Triadic palette from a seed color
chromaforge harmony "#e74c3c" triadic
# Output:
# Palette: triadic
# ========================================
#   0: #E74C3C  rgb(231,76,60)  hsl(6,78%,57%)
#   1: #3CE74C  rgb(60,231,76)  hsl(126,78%,57%)
#   2: #4C3CE7  rgb(76,60,231)  hsl(246,78%,57%)

# Available: complementary, analogous, triadic, split-complementary, tetradic, square, monochromatic
```

### Check Contrast Ratios

```bash
# Check WCAG contrast
chromaforge contrast "#000000" "#ffffff"
# Output:
# Contrast Ratio: 21.00:1
# WCAG Level:     AAA
# Normal text:    ✓ Pass
# Large text:     ✓ Pass
```

### Generate Tints and Shades

```bash
# Lighter tints
chromaforge tints "#3498db" 5
# Output:
# Tints of #3498DB:
#   0: #3498DB
#   1: #5DADE2
#   2: #85C1E9
#   3: #AED6F1
#   4: #D6EAF8

# Darker shades
chromaforge shades "#3498db" 5
```

### Simulate Color Blindness

```bash
# Simulate protanopia (red-blind)
chromaforge blindness "#ff6600" protanopia
# Output: Protanopia (red-blind): #BDBB19

# Simulate all types
chromaforge blindness "#ff6600"
```

### Get Color Information

```bash
chromaforge info "#3498db"
# Output:
# Color Information: #3498DB
# ==================================================
#   RGB:       rgb(52, 152, 219)
#   HSL:       hsl(204.1, 69.9%, 53.1%)
#   HSV:       hsv(204.1, 76.3%, 85.9%)
#   CMYK:      cmyk(76.3%, 30.6%, 0.0%, 14.1%)
#   Lab:       lab(60.16, -6.10, -42.23)
#   Brightness: 50.9%
#   Light:     ✓ Pass
#
#   Contrast with white: 3.15:1 (AALarge)
#   Contrast with black: 6.66:1 (AA)
#   Suggested text:      #000000 (black)
```

## Library API

```go
import "github.com/EdgarOrtegaRamirez/chromaforge/pkg/color"
import "github.com/EdgarOrtegaRamirez/chromaforge/pkg/export"

// Parse a color
c, _ := color.ParseColor("#3498db")

// Convert between formats
hsl := c.ToHSL()
lab := c.ToLab()

// Generate harmonies
triadic := c.Triadic()
analogous := c.Analogous()
palette := c.GeneratePalette("complementary")

// Check accessibility
ratio := color.ContrastRatio(c, color.RGB{255, 255, 255})
level := color.WCAGLevel(ratio)

// Simulate color blindness
simulated := c.SimulateColorBlindness(color.Protanopia)

// Calculate color difference (CIEDE2000)
c1, _ := color.ParseColor("#FF0000")
c2, _ := color.ParseColor("#FF0100")
diff := c1.ToLab().CIEDE2000(c2.ToLab())

// Export palette
css := export.ExportCSS(palette)
tailwind := export.ExportTailwind(palette)
json := export.ExportJSON(palette)
svg := export.ExportSVG(palette)
```

## Architecture

```
chromaforge/
├── cmd/chromaforge/     # CLI entry point
├── pkg/
│   ├── color/           # Core color library
│   │   ├── models.go    # RGB, HSL, HSV, CMYK, Lab types
│   │   ├── convert.go   # Color space conversions
│   │   ├── harmony.go   # Color harmony algorithms
│   │   ├── contrast.go  # WCAG contrast calculation
│   │   └── blindness.go # Color blindness simulation
│   └── export/          # Multi-format export
│       └── export.go    # CSS, Tailwind, JSON, SVG, SCSS
├── tests/               # Test files
└── .github/workflows/   # CI/CD
```

## Algorithms

- **Color Space Conversions** — Standard RGB↔HSL↔HSV↔CMYK↔Lab↔XYZ conversions with proper gamma correction
- **Color Harmony** — Mathematical color wheel relationships (complementary, triadic, analogous, etc.)
- **WCAG Contrast** — Per WCAG 2.1 specification with relative luminance calculation
- **CIEDE2000** — Perceptual color difference formula
- **Color Blindness** — Simulation matrices based on Brettel, Viénot, and Mollon research

## License

MIT
