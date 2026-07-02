package export

import (
	"strings"
	"testing"

	"github.com/EdgarOrtegaRamirez/chromaforge/pkg/color"
)

func TestExportCSS(t *testing.T) {
	palette := color.Palette{
		Name:   "test",
		Colors: []color.RGB{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}},
		Seed:   color.RGB{R: 255, G: 0, B: 0},
	}

	css := ExportCSS(palette)

	if !strings.Contains(css, ":root {") {
		t.Error("CSS should contain :root {")
	}
	if !strings.Contains(css, "--color-0: #FF0000;") {
		t.Error("CSS should contain --color-0: #FF0000;")
	}
	if !strings.Contains(css, "--color-1: #00FF00;") {
		t.Error("CSS should contain --color-1: #00FF00;")
	}
}

func TestExportTailwind(t *testing.T) {
	palette := color.Palette{
		Name:   "brand",
		Colors: []color.RGB{{R: 255, G: 0, B: 0}},
		Seed:   color.RGB{R: 255, G: 0, B: 0},
	}

	tw := ExportTailwind(palette)

	if !strings.Contains(tw, "module.exports") {
		t.Error("Tailwind should contain module.exports")
	}
	if !strings.Contains(tw, "brand-0") {
		t.Error("Tailwind should contain brand-0")
	}
}

func TestExportJSON(t *testing.T) {
	palette := color.Palette{
		Name:    "test",
		Harmony: "triadic",
		Colors:  []color.RGB{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}},
		Seed:    color.RGB{R: 255, G: 0, B: 0},
	}

	json := ExportJSON(palette)

	if !strings.Contains(json, `"name": "test"`) {
		t.Error("JSON should contain name")
	}
	if !strings.Contains(json, `"harmony": "triadic"`) {
		t.Error("JSON should contain harmony")
	}
	if !strings.Contains(json, `"hex": "#FF0000"`) {
		t.Error("JSON should contain hex")
	}
	if !strings.Contains(json, `"rgb": [255, 0, 0]`) {
		t.Error("JSON should contain rgb")
	}
}

func TestExportSVG(t *testing.T) {
	palette := color.Palette{
		Name:   "test",
		Colors: []color.RGB{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}},
		Seed:   color.RGB{R: 255, G: 0, B: 0},
	}

	svg := ExportSVG(palette)

	if !strings.Contains(svg, "<svg") {
		t.Error("SVG should contain <svg")
	}
	if !strings.Contains(svg, `fill="#FF0000"`) {
		t.Error("SVG should contain fill for red")
	}
	if !strings.Contains(svg, `fill="#00FF00"`) {
		t.Error("SVG should contain fill for green")
	}
}

func TestExportSCSS(t *testing.T) {
	palette := color.Palette{
		Name:   "brand",
		Colors: []color.RGB{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}},
		Seed:   color.RGB{R: 255, G: 0, B: 0},
	}

	scss := ExportSCSS(palette)

	if !strings.Contains(scss, "$color-brand-0: #FF0000;") {
		t.Error("SCSS should contain $color-brand-0: #FF0000;")
	}
	if !strings.Contains(scss, "$brand-palette:") {
		t.Error("SCSS should contain $brand-palette:")
	}
}
