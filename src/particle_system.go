package main

import "math/rand"

type IParticleSystem interface {
	AddParticle(particle IParticle)
	Update(deltaTime float32)
	StartEmmiting()
	StopEmmiting()
}

type FireParticleSystem struct {
	particles     []FireParticle
	origin        Vector2
	lifeTimeRange Vector2
	isEmmiting    bool
}

func GetFireParticleSystem(particleCount int, origin Vector2, lifeTimeRange Vector2, wind float32, particle_velocity_x Vector2, particle_velocity_y float32) FireParticleSystem {
	particles := make([]FireParticle, particleCount)

	for i := range particles {
		particles[i] = GetFireParticle()
		particles[i].SetPosition(origin)
		particles[i].SetLifeTime(lifeTimeRange.GetRandomFloat32())
		// Set initial velocities for fire effect
		particles[i].Velocity = Vector2{
			x: particle_velocity_x.GetRandomFloat32() + wind, // Random horizontal drift
			y: -rand.Float32() * particle_velocity_y,         // Upward movement
		}
	}

	return FireParticleSystem{particles: particles, origin: origin, lifeTimeRange: lifeTimeRange, isEmmiting: false}
}

func (system *FireParticleSystem) AddParticle(particle FireParticle) {
	system.particles = append(system.particles, particle)
}

func (system *FireParticleSystem) Update(deltaTime float32) {
	for i := range system.particles {
		particle := &system.particles[i]

		if particle.IsEnabled() == false {
			continue
		}

		particle.Lifetime -= deltaTime

		if particle.Lifetime <= 0 {
			particle.SetPosition(system.origin)
			if system.isEmmiting {
				particle.Lifetime = system.lifeTimeRange.GetRandomFloat32()

			} else {
				particle.SetEnabled(false)
			}
			continue
		}

		velocity := particle.Velocity
		position := particle.GetPosition()
		position.x += velocity.x * deltaTime
		position.y += velocity.y * deltaTime
		particle.SetPosition(position)
	}
}

func (system *FireParticleSystem) StartEmmiting() {
	system.isEmmiting = true

	for i := range system.particles {
		system.particles[i].SetEnabled(true)
	}
}

func (system *FireParticleSystem) StopEmmiting() {
	system.isEmmiting = false
}
