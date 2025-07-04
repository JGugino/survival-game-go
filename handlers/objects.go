package handlers

import (
	"errors"
	"fmt"

	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type Objects struct {
	Objs       map[uuid.UUID]*Object
	Width      int
	Height     int
	CellSize   int
	ObjectGrid [64][64]int
}

func (o *Objects) InitObjectGrid(mapWidth int, mapHeight int, cellSize int) {
	o.Width = mapWidth
	o.Height = mapHeight
	o.CellSize = cellSize

	for y := range o.Height {
		for x := range o.Width {
			o.ObjectGrid[x][y] = int(EMPTY)
		}
	}
}

func (o *Objects) CreateNewObject(object ObjectId, position rl.Vector2, health int, color rl.Color, movable bool, mineable bool, mineLevel utils.MineLevel, droppedItem utils.Item) uuid.UUID {

	objId := uuid.New()

	obj := Object{
		Id:          objId,
		ObjectId:    int(object),
		Position:    position,
		Health:      health,
		MaxHealth:   health,
		Color:       color,
		Movable:     movable,
		Mineable:    mineable,
		Level:       mineLevel,
		DroppedItem: droppedItem,
	}

	o.ObjectGrid[int(position.X)][int(position.Y)] = int(object)

	o.Objs[objId] = &obj

	return objId
}

func (o *Objects) RemoveObject(objId uuid.UUID) {
	obj := o.Objs[objId]
	o.ObjectGrid[int(obj.Position.X)][int(obj.Position.Y)] = 0
	delete(o.Objs, objId)
}

func (o *Objects) DamageObject(objId uuid.UUID, damage int) error {
	obj, ok := o.Objs[objId]
	if !ok {
		return errors.New("no-object")
	}

	obj.Health -= damage

	if obj.Health <= 0 {
		_, _, err := utils.NewWorldItem(obj.DroppedItem.Id, rl.Vector2{X: (obj.Position.X * float32(o.CellSize)), Y: (obj.Position.Y * float32(o.CellSize))})

		if err != nil {

			return err
		}

		o.RemoveObject(objId)
		return nil
	}

	return nil
}

func (o *Objects) GetObjectAtWorldPosition(position rl.Vector2) (uuid.UUID, *Object, error) {
	//Clicked position on the object grid
	gridX := int(position.X) / o.CellSize
	gridY := int(position.Y) / o.CellSize

	//Goes through all of the objects to
	for _, x := range o.Objs {
		if int(x.Position.X) == gridX && int(x.Position.Y) == gridY {
			return x.Id, x, nil
		}
	}

	return uuid.UUID{}, &Object{}, errors.New("no-object")
}

func (o *Objects) DrawObjects(x int, y int, camera *rl.Camera2D) {
	tile := o.ObjectGrid[x][y]

	tMap, err := utils.GetTextureMap("world-objects")

	objectLocation := rl.Vector2{X: float32(x * o.CellSize), Y: float32(y * o.CellSize)}

	if err != nil {
		return
	}

	if tile != 0 {

		_, obj, err := o.GetObjectAtWorldPosition(objectLocation)

		if err != nil {
			return
		}

		if obj.Health < obj.MaxHealth {
			rl.DrawRectangle(int32(objectLocation.X), int32(objectLocation.Y), int32(o.CellSize), 5, rl.White)
			rl.DrawRectangle(int32(objectLocation.X), int32(objectLocation.Y), int32(o.CellSize*obj.Health/obj.MaxHealth), 5, rl.Red)
		}

		switch tile {
		case 1: //Rock
			tMap.DrawTextureAtPositionWithScaling(rl.Vector2Zero(), objectLocation, o.CellSize)
			break
		case 2: //Bush
			tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 1, Y: 0}, objectLocation, o.CellSize)
			break
		}
	}

}

func (o *Objects) DrawObjectGridToConsole() {
	fmt.Println("###OBJECTS###")
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
