package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create a fire particle system with 100 particles
	origin := Vector2{x: 40, y: 40}
	lifeTimeRange := Vector2{x: 0.5, y: 2.0}
	particleSystem := GetFireParticleSystem(2000, origin, lifeTimeRange, 0, Vector2{-10, 10}, 30)

	// Start emitting particles
	particleSystem.StartEmmiting()

	// Initialize model with dimensions
	dimensions := Vector2{x: 80, y: 40}
	m := InitializeParticles(&particleSystem, dimensions)

	// Run the Bubbletea program
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
