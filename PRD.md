# **Product Requirements Document: Terminal-Based Fire Particle Effect**

## **1. Overview**
A terminal-based interactive particle effect simulator that renders an animated fire effect using ASCII characters. The application uses the Bubbletea TUI framework to create a real-time visual display of particles with physics-based movement.

## **2. Core Functionality**

### **2.1 Particle System**
- **Particle Creation**: Initialize 2000 fire particles at a defined origin point (40, 40)
- **Particle Properties**:
  - Position (x, y coordinates)
  - Velocity (horizontal and vertical movement)
  - Lifetime (0.5 to 2.0 seconds, randomized per particle)
  - Enabled state (active/inactive)

### **2.2 Fire Effect Behavior**
- **Movement Pattern**:
  - Particles move upward with negative vertical velocity (0 to -30 units/sec, randomized)
  - Random horizontal drift (-10 to +10 units/sec) for natural fire appearance
  - Physics-based position updates using velocity × deltaTime
  
- **Lifecycle Management**:
  - Particles spawn at origin point
  - Move according to velocity over their lifetime
  - Reset to origin when lifetime expires (if emitting is active)
  - Disable when lifetime expires (if emitting is stopped)

### **2.3 Visual Rendering**
- **Display Area**: 80 columns × 40 rows terminal grid
- **Particle Density Visualization**:
  - Very high density (4+ particles): `@` character
  - High density (3 particles): `#` character
  - Medium density (2 particles): `+` character
  - Low density (1 particle): `*` character
  - Sparse (unused): `·` character (currently not rendered)
  - Empty space: ` ` (blank)
  
- **Frame Rate**: 16ms per frame (~60 FPS)

### **2.4 User Interaction**
- **Controls**:
  - `SPACE`: Toggle particle emission on/off
  - `Q` or `Ctrl+C`: Quit application
  
- **Status Display**: Show current emission state (ON/OFF) at bottom of screen

## **3. Technical Architecture**

### **3.1 Core Components**

**Particle Interface** (`particles.go`):
- `IParticle` interface defining particle behavior contract
- `FireParticle` struct implementing position, velocity, lifetime, and enabled state

**Particle System** (`particle_system.go`):
- `IParticleSystem` interface for system-level operations
- `FireParticleSystem` managing collection of particles
- Handles emission state and particle updates

**Data Structures** (`structs.go`):
- `Vector2` for 2D coordinates and ranges
- Random number generation for particle variation

**UI Model** (`model.go`):
- Bubbletea model following Elm architecture (Model/Update/View)
- Density grid for efficient rendering
- Delta time calculation for frame-rate independent physics

**Entry Point** (`main.go`):
- Initialize particle system with fire configuration
- Set up Bubbletea program and start event loop

### **3.2 Performance Optimizations**
- Pre-allocated density grid (reused each frame)
- Bounds checking to prevent out-of-range rendering
- Efficient grid clearing without reallocation
- Frame-rate independent physics using delta time calculation

## **4. Expected User Experience**

1. **Launch**: User runs the application with `go run ./src` or compiled binary
2. **Initial State**: Fire effect immediately begins animating at position (40, 40) on screen
3. **Observation**: 2000 particles flow upward with realistic fire-like movement and drift
4. **Interaction**: User can pause/resume emission with spacebar
5. **Exit**: User quits cleanly with Q key or Ctrl+C

## **5. Visual Effect Goals**
- Create a convincing ASCII fire animation with 2000 particles
- Smooth, continuous particle flow when emitting (~60 FPS)
- Natural particle dissipation when emission stops
- Responsive controls with immediate visual feedback
- Variable density visualization for realistic fire appearance

## **6. Current Implementation Status**
✅ All core features implemented
✅ Physics-based particle movement with configurable velocity ranges
✅ Interactive controls functional (Space to toggle, Q/Ctrl+C to quit)
✅ Performance-optimized rendering with reusable density grid
✅ 2000 particle system with 0.5-2.0s lifetime range
✅ Multi-level density visualization (@, #, +, *, space)

## **7. Known Optimization Opportunities**
🔄 Active particle tracking to skip disabled particle checks
🔄 Spatial hashing for efficient density grid population
🔄 String builder capacity pre-calculation
🔄 Proper random number seeding (currently unseeded)
🔄 Minor typo fix: "Emmiting" → "Emitting" throughout codebase

## **8. Build & Run Commands**
- **Build**: `go build -o particle_effect ./src`
- **Run**: `go run ./src`
- **Test**: `go test ./...`
- **Format**: `go fmt ./...`
- **Vet**: `go vet ./...`

---

This application demonstrates real-time particle simulation in a terminal environment with proper physics, lifecycle management, and user interactivity using Go and the Bubbletea framework. The system efficiently handles 2000 particles at 60 FPS with frame-rate independent physics.
