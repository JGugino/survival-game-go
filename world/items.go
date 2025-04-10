package world

import (
	"errors"
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
}

var itemMap map[string]*Item = make(map[string]*Item, 0)

func InitItemMap() {
	itemMap["rock"] = &Item{Id: I_ROCK, Name: "Rock", Type: ITEM, MaxStack: 100, MineLevel: ML_LOW, MineDamage: 1}
	itemMap["coal"] = &Item{Id: I_COAL, Name: "Coal", Type: ITEM, MaxStack: 100, MineLevel: ML_LOW, MineDamage: 1}
	itemMap["seeds"] = &Item{Id: I_SEEDS, Name: "Seeds", Type: ITEM, MaxStack: 50, MineLevel: ML_LOW, MineDamage: 1}
	itemMap["torch"] = &Item{Id: I_TORCH, Name: "Torch", Type: ITEM, MaxStack: 50, MineLevel: ML_LOW, MineDamage: 1}
	itemMap["pickaxe"] = &Item{Id: I_PICKAXE, Name: "Pickaxe", Type: WEAPON, MaxStack: 1, MineLevel: ML_MED, MineDamage: 10}
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
