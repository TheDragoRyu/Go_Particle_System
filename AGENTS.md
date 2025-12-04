# Agent Guidelines

## Project Overview
A terminal-based fire particle effect simulator (2000 particles @ 60 FPS) using Bubbletea TUI framework. Core behavior: particles spawn at origin (40,40), move upward with random horizontal drift, reset on lifetime expiration. Controls: SPACE to toggle emission, Q to quit.

## Build/Run/Test Commands
- **Build**: `go build -o particle_effect ./src`
- **Run**: `go run ./src`
- **Test all**: `go test ./...`
- **Test single package**: `go test ./src`
- **Lint**: `go vet ./...` or `golangci-lint run` (if available)

## Code Style
- **Package**: All source code in `src/` directory, package `main`
- **Imports**: Standard library first, then external packages (grouped and sorted alphabetically)
- **Types**: Use PascalCase for exported types, camelCase for unexported
- **Interfaces**: Prefix with `I` (e.g., `IParticle`, `IParticleSystem`)
- **Structs**: Use lowercase field names for unexported fields (e.g., `x`, `y` in `Vector2`)
- **Formatting**: Run `go fmt ./...` before committing
- **Go version**: 1.25.3 (as specified in go.mod)
- **Framework**: Uses Bubbletea for TUI - follow Elm architecture (Model/Update/View)
- **Error handling**: Return errors explicitly, avoid panics in library code
- **Comments**: Add package-level comments and document exported functions/types

## Architecture Patterns
- **Particle lifecycle**: Enable → Update position → Lifetime expires → Reset to origin (if emitting) or Disable
- **Physics**: Frame-rate independent using deltaTime calculation (`position += velocity * deltaTime`)
- **Rendering**: Density-based visualization (≥4:@ 3:# 2:+ 1:* 0:space) using pre-allocated reusable grid
- **Performance**: Pre-allocated slices, bounds checking, avoid reallocations per frame

## Known Issues to Fix
- Typo: "Emmiting" → "Emitting" throughout codebase (StartEmmiting, StopEmmiting, isEmmiting)
- Random number generation: Need proper seeding (currently unseeded)
- Particle density mapping inconsistency: PRD says 2:+, code implements 2:* (verify intended behavior)
