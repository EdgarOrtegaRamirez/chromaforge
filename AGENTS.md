# AGENTS.md — ChromaForge

## Project Overview
ChromaForge is a color palette and scheme management CLI tool and Go library. It provides color format conversion, harmony generation, WCAG contrast checking, color blindness simulation, and multi-format export.

## Build & Test
```bash
# Build
go build ./cmd/chromaforge/

# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run specific test package
go test ./pkg/color/ -v
go test ./pkg/export/ -v
```

## Project Structure
- `cmd/chromaforge/main.go` — CLI entry point
- `pkg/color/models.go` — Color type definitions (RGB, HSL, HSV, CMYK, Lab, XYZ)
- `pkg/color/convert.go` — Color space conversions with gamma correction
- `pkg/color/harmony.go` — Color harmony algorithms
- `pkg/color/contrast.go` — WCAG contrast ratio calculation
- `pkg/color/blindness.go` — Color blindness simulation
- `pkg/export/export.go` — Multi-format palette export (CSS, Tailwind, JSON, SVG, SCSS)

## Code Style
- Use Go standard formatting (gofmt)
- Keep functions focused and well-documented
- Use table-driven tests
- Handle edge cases (black, white, grays)

## Key Algorithms
- Color space conversions with proper sRGB gamma correction
- WCAG 2.1 relative luminance calculation
- CIEDE2000 perceptual color difference
- Color blindness simulation matrices

## Dependencies
- No external dependencies — pure Go standard library only
