package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ParticleRenderer struct {
	particleSystem *FireParticleSystem
	dimensions     Vector2
	isOn           bool
	lastUpdate     time.Time
	density        [][]int // Reusable density grid
}

const HIGH_PARTICLES string = "\x1b[97m@\x1b[0m"        // bright white — hottest core
const MEDIUM_HIGH_PARTICLES string = "\x1b[93m#\x1b[0m" // bright yellow — hot flame
const MEDIUM_PARTICLES string = "\x1b[33m+\x1b[0m"      // yellow — mid heat
const LOW_PARTICLES string = "\x1b[31m*\x1b[0m"         // red — cooler edge
const SPARSE_PARTICLES string = " "                     // unused; 0 density maps to space

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func InitializeParticles(particleSystem *FireParticleSystem, dimensions Vector2) ParticleRenderer {
	width := int(dimensions.x)
	height := int(dimensions.y)

	// Pre-allocate density grid once
	density := make([][]int, height)
	for i := range density {
		density[i] = make([]int, width)
	}

	return ParticleRenderer{
		particleSystem: particleSystem,
		dimensions:     dimensions,
		isOn:           true,
		lastUpdate:     time.Now(),
		density:        density,
	}
}

func (m ParticleRenderer) Init() tea.Cmd {
	return tickCmd()
}

func (m ParticleRenderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ":
			m.isOn = !m.isOn
			if m.isOn {
				m.particleSystem.StartEmmiting()
			} else {
				m.particleSystem.StopEmmiting()
			}
		}

	case tickMsg:
		now := time.Time(msg)
		deltaTime := float32(now.Sub(m.lastUpdate).Seconds())
		m.lastUpdate = now
		m.particleSystem.Update(deltaTime)
		return m, tickCmd()
	}

	return m, nil
}

func (m ParticleRenderer) View() string {
	width := int(m.dimensions.x)
	height := int(m.dimensions.y)

	// Clear the density grid (reuse existing allocation)
	for i := range m.density {
		for j := range m.density[i] {
			m.density[i][j] = 0
		}
	}

	// Count particles at each grid position
	for _, particle := range m.particleSystem.particles {
		if !particle.IsEnabled() {
			continue
		}

		pos := particle.GetPosition()
		x := int(pos.x)
		y := int(pos.y)

		// Check bounds and increment density
		if x >= 0 && x < width && y >= 0 && y < height {
			m.density[y][x]++
		}
	}

	// Convert density grid to string with appropriate characters
	var sb strings.Builder
	for _, row := range m.density {
		for _, count := range row {
			var char string
			if count >= 8 {
				char = HIGH_PARTICLES
			} else if count >= 3 {
				char = MEDIUM_HIGH_PARTICLES
			} else if count == 2 {
				char = MEDIUM_PARTICLES
			} else if count == 1 {
				char = LOW_PARTICLES
			} else {
				char = " "
			}
			sb.WriteString(char)
		}
		sb.WriteString("\n")
	}

	status := "OFF"
	if m.isOn {
		status = "ON"
	}
	sb.WriteString("\nPress SPACE to toggle | Press Q to quit | Status: ")
	sb.WriteString(status)

	return sb.String()
}
