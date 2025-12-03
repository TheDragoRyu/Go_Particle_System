package main

type IParticle interface {
	GetPosition() Vector2
	SetPosition(position Vector2)
	SetEnabled(bool)
	IsEnabled() bool
}

type FireParticle struct {
	position Vector2
	enabled  bool
	Velocity Vector2
	Lifetime float32
}

func GetFireParticle() FireParticle {
	return FireParticle{}
}

func (particle FireParticle) GetPosition() Vector2 {
	return particle.position
}
func (particle *FireParticle) SetPosition(position Vector2) {
	particle.position = position
}
func (particle *FireParticle) SetLifeTime(lifeTime float32) {
	particle.Lifetime = lifeTime
}
func (particle *FireParticle) IsEnabled() bool {
	return particle.enabled
}
func (particle *FireParticle) SetEnabled(enabledState bool) {
	particle.enabled = enabledState
}
