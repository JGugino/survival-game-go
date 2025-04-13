package world

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type ItemId int32
type ItemType int32

type StackLocation int8

const (
	//Stack Locations
	SL_HOTBAR    StackLocation = 0
	SL_INVENTORY StackLocation = 1
	SL_WORLD     StackLocation = 2

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

type ItemStack struct {
	Id            uuid.UUID
	ItemId        ItemId
	StackSize     int
	InventorySlot rl.Vector2
	Section       StackLocation
}

var itemStacks map[uuid.UUID]*ItemStack = make(map[uuid.UUID]*ItemStack, 0)

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

// Item stacks
func CreateNewItemStack(inventorySlot rl.Vector2, itemId ItemId, section StackLocation) (uuid.UUID, *ItemStack) {
	stackId := uuid.New()
	newStack := &ItemStack{
		Id:            stackId,
		InventorySlot: inventorySlot,
		ItemId:        itemId,
		StackSize:     1,
		Section:       section,
	}

	itemStacks[stackId] = newStack

	return stackId, newStack
}

func MoveItemStack(stackId uuid.UUID, newLocation rl.Vector2, newSection StackLocation) error {
	stack, ok := itemStacks[stackId]

	if !ok {
		return errors.New("no-item-stack")
	}

	stack.InventorySlot = newLocation
	stack.Section = newSection

	return nil
}

func GetItemStackAtInventorySlot(slot rl.Vector2, section StackLocation) (uuid.UUID, *ItemStack, error) {

	for id, stack := range itemStacks {
		if stack.Section == section {
			if stack.InventorySlot == slot {
				return id, stack, nil
			}
		}
	}

	return uuid.UUID{}, &ItemStack{}, errors.New("no-stack-at-slot")
}

func DrawItemStacksToConsole() {
	fmt.Println("###START ITEM STACKS")
	for _, i := range itemStacks {
		if i.Section == SL_HOTBAR {
			fmt.Println("In Hotbar")
		} else if i.Section == SL_INVENTORY {
			fmt.Println("In Inventory")
		}

	}
	fmt.Println("###END ITEM STACKS")
}
