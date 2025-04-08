package objects

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectId int

const (
	EMPTY ObjectId = 0
	ROCK  ObjectId = 1
	TREE  ObjectId = 2
)

type Objects struct {
	Objs       map[string]Object
	Width      int
	Height     int
	CellSize   int
	ObjectGrid [64][64]int
}

func (o *Objects) InitObjectGrid() {
	for y := range o.Height {
		for x := range o.Width {
			o.ObjectGrid[x][y] = int(EMPTY)
		}
	}
}

func (o *Objects) CreateNewObject(object ObjectId, position rl.Vector2, health int, color rl.Color, movable bool, mineable bool, mineLevel MineLevel, droppedItem int) string {

	objId := string(rl.GetRandomValue(0, 2147483647))

	obj := Object{
		Id:          objId,
		ObjectId:    int(object),
		Position:    position,
		Health:      health,
		Color:       color,
		Movable:     movable,
		Mineable:    mineable,
		Level:       mineLevel,
		DroppedItem: droppedItem,
	}

	o.ObjectGrid[int(position.X)][int(position.Y)] = int(object)

	o.Objs[objId] = obj

	return objId
}

func (o *Objects) RemoveObject(objId string) {
	obj := o.Objs[objId]
	o.ObjectGrid[int(obj.Position.X)][int(obj.Position.Y)] = 0
	delete(o.Objs, objId)
}

func (o *Objects) GetObjectAtWorldPosition(position rl.Vector2) (string, Object) {
	gridX := int(position.X) / o.CellSize
	gridY := int(position.Y) / o.CellSize

	for _, x := range o.Objs {
		fmt.Println("---Object Start---")
		fmt.Println(x.Position.X)
		fmt.Println(x.Position.Y)
		fmt.Println(gridX)
		fmt.Println(gridY)
		fmt.Println("---Object End---")
		if int(x.Position.X) == gridX && int(x.Position.Y) == gridY {

			return x.Id, x
		}
	}

	return "", Object{}
}

func (o *Objects) DrawObjects() {
	for y := range o.Height {
		for x := range o.Width {
			tile := o.ObjectGrid[x][y]

			if tile == 1 { //Rock
				rl.DrawRectangle(int32(x*o.CellSize), int32(y*o.CellSize), int32(o.CellSize), int32(o.CellSize), rl.DarkGray)
			}

		}
	}
}

func (o *Objects) DrawObjectGridToConsole() {
	for y := range o.Height {
		for x := range o.Width {
			fmt.Print(o.ObjectGrid[x][y])
		}
		fmt.Print("\n")
	}
}

func (o *Objects) DrawObjectsToConsole() {
	for _, x := range o.Objs {
		fmt.Print(x)
	}
}
