package handlers

import (
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type ObjectId int

const (
	//Object Id's
	EMPTY ObjectId = 0
	ROCK  ObjectId = 1
	TREE  ObjectId = 2
)

type Object struct {
	Id          uuid.UUID
	ObjectId    int
	Position    rl.Vector2
	Health      int
	Color       rl.Color
	Movable     bool
	Mineable    bool
	Level       utils.MineLevel
	DroppedItem utils.Item
}
