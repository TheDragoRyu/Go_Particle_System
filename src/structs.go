package main

import "math/rand"

type Vector2 struct {
	x float32
	y float32
}

func (vectorRange Vector2) GetRandomFloat32() float32 {
	return rand.Float32()*(vectorRange.y-vectorRange.x) + vectorRange.x
}
