package export

import (
	"fmt"
	"strings"

	"github.com/EdgarOrtegaRamirez/chromaforge/pkg/color"
)

// ExportCSS generates CSS custom properties from a palette.
func ExportCSS(p color.Palette) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("/* %s palette */\n", p.Name))
	sb.WriteString(":root {\n")
	for i, c := range p.Colors {
		hsl := c.ToHSL()
		sb.WriteString(fmt.Sprintf("  --color-%d: %s;\n", i, c.String()))
		sb.WriteString(fmt.Sprintf("  --color-%d-rgb: %d, %d, %d;\n", i, c.R, c.G, c.B))
		sb.WriteString(fmt.Sprintf("  --color-%d-hsl: %d, %d%%, %d%%;\n", i, int(hsl.H), int(hsl.S), int(hsl.L)))
	}
	sb.WriteString("}\n")
	return sb.String()
}

// ExportTailwind generates a Tailwind CSS config snippet from a palette.
func ExportTailwind(p color.Palette) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("// %s palette - Tailwind CSS config\n", p.Name))
	sb.WriteString("module.exports = {\n")
	sb.WriteString("  theme: {\n")
	sb.WriteString("    extend: {\n")
	sb.WriteString("      colors: {\n")

	for i, c := range p.Colors {
		hsl := c.ToHSL()
		sb.WriteString(fmt.Sprintf("        '%s-%d': '%s',\n", p.Name, i, c.String()))
		sb.WriteString(fmt.Sprintf("        '%s-%d-hsl': 'hsl(%d, %d%%, %d%%)',\n", p.Name, i, int(hsl.H), int(hsl.S), int(hsl.L)))
	}

	sb.WriteString("      },\n")
	sb.WriteString("    },\n")
	sb.WriteString("  },\n")
	sb.WriteString("}\n")
	return sb.String()
}

// ExportJSON generates a JSON representation of a palette.
func ExportJSON(p color.Palette) string {
	var sb strings.Builder
	sb.WriteString("{\n")
	sb.WriteString(fmt.Sprintf("  \"name\": \"%s\",\n", p.Name))
	sb.WriteString(fmt.Sprintf("  \"harmony\": \"%s\",\n", p.Harmony))
	sb.WriteString("  \"seed\": {\n")
	sb.WriteString(fmt.Sprintf("    \"hex\": \"%s\",\n", p.Seed.String()))
	lab := p.Seed.ToLab()
	sb.WriteString(fmt.Sprintf("    \"rgb\": [%d, %d, %d],\n", p.Seed.R, p.Seed.G, p.Seed.B))
	sb.WriteString(fmt.Sprintf("    \"hsl\": [%.1f, %.1f, %.1f],\n", p.Seed.ToHSL().H, p.Seed.ToHSL().S, p.Seed.ToHSL().L))
	sb.WriteString(fmt.Sprintf("    \"lab\": [%.2f, %.2f, %.2f]\n", lab.L, lab.A, lab.B))
	sb.WriteString("  },\n")
	sb.WriteString("  \"colors\": [\n")
	for i, c := range p.Colors {
		hsl := c.ToHSL()
		lab := c.ToLab()
		sb.WriteString("    {\n")
		sb.WriteString(fmt.Sprintf("      \"hex\": \"%s\",\n", c.String()))
		sb.WriteString(fmt.Sprintf("      \"rgb\": [%d, %d, %d],\n", c.R, c.G, c.B))
		sb.WriteString(fmt.Sprintf("      \"hsl\": [%.1f, %.1f, %.1f],\n", hsl.H, hsl.S, hsl.L))
		sb.WriteString(fmt.Sprintf("      \"lab\": [%.2f, %.2f, %.2f]\n", lab.L, lab.A, lab.B))
		sb.WriteString("    }")
		if i < len(p.Colors)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}
	sb.WriteString("  ]\n")
	sb.WriteString("}\n")
	return sb.String()
}

// ExportSVG generates an SVG palette visualization.
func ExportSVG(p color.Palette) string {
	width := 600
	boxWidth := width / len(p.Colors)
	height := 200

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">`, width, height))
	sb.WriteString("\n")

	for i, c := range p.Colors {
		x := i * boxWidth
		sb.WriteString(fmt.Sprintf(`  <rect x="%d" y="0" width="%d" height="%d" fill="%s"/>`, x, boxWidth, height, c.String()))
		sb.WriteString("\n")

		// Add text label
		textColor := "white"
		if c.IsLight() {
			textColor = "black"
		}
		sb.WriteString(fmt.Sprintf(`  <text x="%d" y="%d" fill="%s" font-family="monospace" font-size="12" text-anchor="middle">%s</text>`,
			x+boxWidth/2, height-10, textColor, c.String()))
		sb.WriteString("\n")
	}

	sb.WriteString("</svg>\n")
	return sb.String()
}

// ExportSCSS generates SCSS variables from a palette.
func ExportSCSS(p color.Palette) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("// %s palette - SCSS variables\n", p.Name))
	for i, c := range p.Colors {
		sb.WriteString(fmt.Sprintf("$color-%s-%d: %s;\n", p.Name, i, c.String()))
	}
	sb.WriteString("\n// Color map\n")
	sb.WriteString(fmt.Sprintf("$%s-palette: (\n", p.Name))
	for i, c := range p.Colors {
		comma := ","
		if i == len(p.Colors)-1 {
			comma = ""
		}
		sb.WriteString(fmt.Sprintf("  '%d': %s%s\n", i, c.String(), comma))
	}
	sb.WriteString(");\n")
	return sb.String()
}
