package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/EdgarOrtegaRamirez/chromaforge/pkg/color"
	"github.com/EdgarOrtegaRamirez/chromaforge/pkg/export"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "convert":
		cmdConvert(args)
	case "harmony":
		cmdHarmony(args)
	case "contrast":
		cmdContrast(args)
	case "tints":
		cmdTints(args)
	case "shades":
		cmdShades(args)
	case "blindness":
		cmdBlindness(args)
	case "info":
		cmdInfo(args)
	case "version":
		fmt.Printf("chromaforge %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`ChromaForge - Color Palette & Scheme Management CLI

Usage:
  chromaforge <command> [arguments]

Commands:
  convert <color> <format>    Convert a color to different formats
  harmony <color> <type>      Generate a color harmony palette
  contrast <color1> <color2>  Calculate WCAG contrast ratio
  tints <color> [steps]       Generate lighter tints
  shades <color> [steps]      Generate darker shades
  blindness <color> [type]    Simulate color blindness
  info <color>                Show detailed color information
  version                     Show version
  help                        Show this help message

Color Formats:
  Hex:    #FF0000, FF0000, #F00
  RGB:    rgb(255, 0, 0)
  HSL:    hsl(0, 100%, 50%)

Harmony Types:
  complementary     Opposite on color wheel
  analogous         Adjacent colors
  triadic           Three evenly spaced
  split-complementary  Base + two adjacent to complement
  tetradic          Four colors in rectangle
  square            Four colors 90° apart
  monochromatic     Variations of one color

Examples:
  chromaforge convert "#3498db" hsl
  chromaforge harmony "#e74c3c" triadic
  chromaforge contrast "#000000" "#ffffff"
  chromaforge tints "#3498db" 5
  chromaforge blindness "#ff6600" protanopia
  chromaforge info "#3498db"`)
}

func cmdConvert(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge convert <color> <format>\n")
		fmt.Fprintf(os.Stderr, "Formats: hex, rgb, hsl, hsv, cmyk, lab\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	format := strings.ToLower(args[1])
	switch format {
	case "hex":
		fmt.Println(c.String())
	case "rgb":
		fmt.Printf("rgb(%d, %d, %d)\n", c.R, c.G, c.B)
	case "hsl":
		hsl := c.ToHSL()
		fmt.Printf("hsl(%.1f, %.1f%%, %.1f%%)\n", hsl.H, hsl.S, hsl.L)
	case "hsv":
		hsv := c.ToHSV()
		fmt.Printf("hsv(%.1f, %.1f%%, %.1f%%)\n", hsv.H, hsv.S, hsv.V)
	case "cmyk":
		cmyk := c.ToCMYK()
		fmt.Printf("cmyk(%.1f%%, %.1f%%, %.1f%%, %.1f%%)\n", cmyk.C, cmyk.M, cmyk.Y, cmyk.K)
	case "lab":
		lab := c.ToLab()
		fmt.Printf("lab(%.2f, %.2f, %.2f)\n", lab.L, lab.A, lab.B)
	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s\n", format)
		os.Exit(1)
	}
}

func cmdHarmony(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge harmony <color> <type>\n")
		fmt.Fprintf(os.Stderr, "Types: complementary, analogous, triadic, split-complementary, tetradic, square, monochromatic\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	harmony := strings.ToLower(args[1])
	palette := c.GeneratePalette(harmony)

	fmt.Printf("Palette: %s\n", palette.Name)
	fmt.Println(strings.Repeat("=", 40))
	for i, clr := range palette.Colors {
		hsl := clr.ToHSL()
		fmt.Printf("  %d: %s  rgb(%d,%d,%d)  hsl(%.0f,%.0f%%,%.0f%%)\n",
			i, clr.String(), clr.R, clr.G, clr.B, hsl.H, hsl.S, hsl.L)
	}

	// Also output CSS
	fmt.Println("\nCSS:")
	fmt.Println(export.ExportCSS(palette))
}

func cmdContrast(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge contrast <color1> <color2>\n")
		os.Exit(1)
	}

	c1, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing color 1: %v\n", err)
		os.Exit(1)
	}

	c2, err := color.ParseColor(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing color 2: %v\n", err)
		os.Exit(1)
	}

	ratio := color.ContrastRatio(c1, c2)
	level := color.WCAGLevel(ratio)

	fmt.Printf("Contrast Ratio: %.2f:1\n", ratio)
	fmt.Printf("WCAG Level:     %s\n", level)
	fmt.Printf("Normal text:    %s\n", boolStr(color.IsAccessible(c1, c2)))
	fmt.Printf("Large text:     %s\n", boolStr(color.IsAccessibleLarge(c1, c2)))
}

func cmdTints(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge tints <color> [steps]\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	steps := 5
	if len(args) > 1 {
		fmt.Sscanf(args[1], "%d", &steps)
	}

	tints := c.GenerateTints(steps)
	fmt.Printf("Tints of %s:\n", c.String())
	for i, t := range tints {
		fmt.Printf("  %d: %s\n", i, t.String())
	}
}

func cmdShades(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge shades <color> [steps]\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	steps := 5
	if len(args) > 1 {
		fmt.Sscanf(args[1], "%d", &steps)
	}

	shades := c.GenerateShades(steps)
	fmt.Printf("Shades of %s:\n", c.String())
	for i, s := range shades {
		fmt.Printf("  %d: %s\n", i, s.String())
	}
}

func cmdBlindness(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge blindness <color> [type]\n")
		fmt.Fprintf(os.Stderr, "Types: protanopia, deuteranopia, tritanopia, achromatopsia, protanomaly, deuteranomaly, tritanomaly\n")
		fmt.Fprintf(os.Stderr, "If no type specified, simulates all types.\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Original: %s\n\n", c.String())

	if len(args) > 1 {
		typeName := strings.ToLower(args[1])
		cbType := parseBlindnessType(typeName)
		simulated := c.SimulateColorBlindness(cbType)
		fmt.Printf("  %s: %s\n", color.ColorBlindnessNames[cbType], simulated.String())
	} else {
		for cbType, name := range color.ColorBlindnessNames {
			simulated := c.SimulateColorBlindness(cbType)
			fmt.Printf("  %s: %s\n", name, simulated.String())
		}
	}
}

func cmdInfo(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: chromaforge info <color>\n")
		os.Exit(1)
	}

	c, err := color.ParseColor(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	hsl := c.ToHSL()
	hsv := c.ToHSV()
	cmyk := c.ToCMYK()
	lab := c.ToLab()
	white := color.RGB{R: 255, G: 255, B: 255}
	black := color.RGB{R: 0, G: 0, B: 0}
	textColor := color.SuggestTextColor(c)

	fmt.Printf("Color Information: %s\n", c.String())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("  RGB:       rgb(%d, %d, %d)\n", c.R, c.G, c.B)
	fmt.Printf("  HSL:       hsl(%.1f, %.1f%%, %.1f%%)\n", hsl.H, hsl.S, hsl.L)
	fmt.Printf("  HSV:       hsv(%.1f, %.1f%%, %.1f%%)\n", hsv.H, hsv.S, hsv.V)
	fmt.Printf("  CMYK:      cmyk(%.1f%%, %.1f%%, %.1f%%, %.1f%%)\n", cmyk.C, cmyk.M, cmyk.Y, cmyk.K)
	fmt.Printf("  Lab:       lab(%.2f, %.2f, %.2f)\n", lab.L, lab.A, lab.B)
	fmt.Printf("  Brightness: %.1f%%\n", c.Brightness()/255*100)
	fmt.Printf("  Light:     %s\n", boolStr(c.IsLight()))
	fmt.Println()
	fmt.Printf("  Contrast with white: %.2f:1 (%s)\n", color.ContrastRatio(c, white), color.WCAGLevel(color.ContrastRatio(c, white)))
	fmt.Printf("  Contrast with black: %.2f:1 (%s)\n", color.ContrastRatio(c, black), color.WCAGLevel(color.ContrastRatio(c, black)))
	fmt.Printf("  Suggested text:      %s (%s)\n", textColor.String(), textColorName(textColor))
}

func parseBlindnessType(s string) color.ColorBlindnessType {
	switch strings.ToLower(s) {
	case "protanopia":
		return color.Protanopia
	case "deuteranopia":
		return color.Deuteranopia
	case "tritanopia":
		return color.Tritanopia
	case "achromatopsia":
		return color.Achromatopsia
	case "protanomaly":
		return color.Protanomaly
	case "deuteranomaly":
		return color.Deuteranomaly
	case "tritanomaly":
		return color.Tritanomaly
	default:
		return color.Achromatopsia
	}
}

func boolStr(b bool) string {
	if b {
		return "✓ Pass"
	}
	return "✗ Fail"
}

func textColorName(c color.RGB) string {
	if c.R == 0 && c.G == 0 && c.B == 0 {
		return "black"
	}
	return "white"
}
