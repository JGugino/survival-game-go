package world

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type ItemId int32
type ItemType int32

const (
	//Item Ids
	I_ROCK    ItemId = 1
	I_COAL    ItemId = 2
	I_SEEDS   ItemId = 3
	I_TORCH   ItemId = 4
	I_PICKAXE ItemId = 5

	//Item Type
	ITEM   ItemType = 0
	TOOL   ItemType = 1
	WEAPON ItemType = 2
)

type Item struct {
	Id         ItemId
	Name       string
	MaxStack   int
	Type       ItemType
	MineLevel  MineLevel
	MineDamage int
	Color      rl.Color
}

type WorldItem struct {
	Id       uuid.UUID
	Item     Item
	Position rl.Vector2
}

var itemMap map[string]*Item = make(map[string]*Item, 0)

func InitItemMap() {
	itemMap["rock"] = &Item{Id: I_ROCK, Name: "Rock", Type: ITEM, MaxStack: 10, MineLevel: ML_LOW, MineDamage: 1, Color: rl.DarkGray}
	itemMap["coal"] = &Item{Id: I_COAL, Name: "Coal", Type: ITEM, MaxStack: 100, MineLevel: ML_LOW, MineDamage: 1, Color: rl.Black}
	itemMap["seeds"] = &Item{Id: I_SEEDS, Name: "Seeds", Type: ITEM, MaxStack: 50, MineLevel: ML_LOW, MineDamage: 1, Color: rl.Green}
	itemMap["torch"] = &Item{Id: I_TORCH, Name: "Torch", Type: ITEM, MaxStack: 50, MineLevel: ML_LOW, MineDamage: 1, Color: rl.Orange}
	itemMap["pickaxe"] = &Item{Id: I_PICKAXE, Name: "Pickaxe", Type: WEAPON, MaxStack: 1, MineLevel: ML_MED, MineDamage: 10, Color: rl.Brown}
}

func GetItemByName(name string) (*Item, error) {
	found, ok := itemMap[name]

	if !ok {
		return &Item{}, errors.New("no-item")
	}

	return found, nil
}

func GetItemByItemId(itemId ItemId) (*Item, error) {
	switch itemId {
	case I_ROCK:
		return itemMap["rock"], nil
	case I_COAL:
		return itemMap["coal"], nil
	case I_SEEDS:
		return itemMap["seeds"], nil
	case I_TORCH:
		return itemMap["torch"], nil
	case I_PICKAXE:
		return itemMap["pickaxe"], nil
	default:
		return &Item{}, errors.New("no-item")
	}
}

// World Items
var WorldItems map[uuid.UUID]WorldItem = make(map[uuid.UUID]WorldItem, 0)

func (i *WorldItem) HandleItemCollision(rect rl.Rectangle) (bool, *Item) {
	if rl.CheckCollisionCircleRec(i.Position, 12, rect) {
		return true, &i.Item
	}

	return false, &i.Item
}

func NewWorldItem(itemId ItemId, position rl.Vector2) (uuid.UUID, *WorldItem, error) {
	item, err := GetItemByItemId(itemId)

	if err != nil {
		return uuid.UUID{}, &WorldItem{}, err
	}

	id := uuid.New()
	worldItem := &WorldItem{Id: id, Item: *item, Position: position}

	WorldItems[id] = *worldItem

	return id, worldItem, nil
}

func RemoveWorldItem(id uuid.UUID) error {

	if _, ok := WorldItems[id]; !ok {
		return errors.New("no-world-item")
	}

	delete(WorldItems, id)

	return nil
}

func DrawWorldItems() {
	for _, i := range WorldItems {
		rl.DrawCircle(int32(i.Position.X), int32(i.Position.Y), 12, i.Item.Color)
	}
}
