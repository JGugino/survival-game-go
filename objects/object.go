package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MineLevel int16

const (
	ML_LOW  MineLevel = 0
	ML_MED  MineLevel = 1
	ML_HIGH MineLevel = 2
)

type Object struct {
	Id          string
	ObjectId    int
	Position    rl.Vector2
	Health      int
	Color       rl.Color
	Movable     bool
	Mineable    bool
	Level       MineLevel
	DroppedItem int
}
