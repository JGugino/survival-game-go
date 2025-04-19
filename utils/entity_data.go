package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type FacingDirection int32

const (
	UP    FacingDirection = 0
	DOWN  FacingDirection = 1
	LEFT  FacingDirection = 2
	RIGHT FacingDirection = 3
)

type EntityData struct {
	Position  rl.Vector2
	Health    int
	Speed     float32
	Width     int
	Height    int
	Direction FacingDirection
	Moving    bool
}
