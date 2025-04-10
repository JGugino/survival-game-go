package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type ObjectId int
type MineLevel int16

const (
	//Object Id's
	EMPTY ObjectId = 0
	ROCK  ObjectId = 1
	TREE  ObjectId = 2

	//Mining Levels
	ML_LOW  MineLevel = 0
	ML_MED  MineLevel = 1
	ML_HIGH MineLevel = 2
)

type Object struct {
	Id          uuid.UUID
	ObjectId    int
	Position    rl.Vector2
	Health      int
	Color       rl.Color
	Movable     bool
	Mineable    bool
	Level       MineLevel
	DroppedItem Item
}
