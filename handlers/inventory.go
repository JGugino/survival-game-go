package handlers

import (
	"errors"
	"fmt"

	"github.com/JGugino/survival-game-go/utils"
	"github.com/JGugino/survival-game-go/world"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type InventoryPosition struct {
	InventoryXPadding    int
	InventoryYPadding    int
	InventoryYOffset     int
	InventoryCellXOffset int
	InventoryCellYOffset int
	HotbarXOffset        int
	HotbarYOffset        int
	HotbarXPadding       int
	HotbarYPadding       int
	HotbarCellXOffset    int
	HotbarCellYOffset    int
}

type Inventory struct {
	HotbarSize            int
	MainInventory         utils.Container
	Hotbar                [9]world.ItemId
	CellSize              int
	SelectedHotbarSlot    int
	Visible               bool
	Positioning           InventoryPosition
	HoldingItem           bool
	SelectedInventoryItem struct {
		itemId   world.ItemId
		stackId  uuid.UUID
		startPos rl.Vector2
		section  world.StackLocation
	}
	Hover struct {
		InventoryHovering  bool
		InventoryHoverSlot rl.Vector2
		HotbarHovering     bool
		HotbarHoverSlot    rl.Vector2
	}
}

var transparentWhite rl.Color = rl.Color{R: 255, G: 255, B: 255, A: 230}

func (i *Inventory) DrawInventory() {
	offsetX := (rl.GetScreenWidth() - i.MainInventory.ScreenWidth) / 2
	offsetY := (rl.GetScreenHeight() - i.MainInventory.ScreenHeight) / 2

	// Inventory background
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(offsetX), Y: float32(offsetY - i.Positioning.InventoryYOffset), Width: float32(i.MainInventory.ScreenWidth + i.Positioning.InventoryXPadding), Height: float32(i.MainInventory.ScreenHeight + i.Positioning.InventoryYPadding)}, 0.1, 1, transparentWhite)

	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			slotXOffset := int32(x*i.CellSize + offsetX + i.Positioning.InventoryCellXOffset)
			slotYOffset := int32(y*i.CellSize + offsetY + i.Positioning.InventoryCellYOffset)
			// Inventory slots
			rl.DrawRectangleLines(slotXOffset, slotYOffset, int32(i.CellSize), int32(i.CellSize), rl.Black)

			if i.Hover.InventoryHovering {
				if int(i.Hover.InventoryHoverSlot.X) == x && int(i.Hover.InventoryHoverSlot.Y) == y {
					rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(slotXOffset), Y: float32(slotYOffset), Width: float32(i.CellSize), Height: float32(i.CellSize)}, 2, rl.Black)
				}
			}

			//Inventory slot items
			if i.MainInventory.ItemGrid[x][y] != 0 {
				slotItem, err := world.GetItemByItemId(i.MainInventory.ItemGrid[x][y])
				if err != nil {
					return
				}

				_, stack, err := world.GetItemStackAtInventorySlot(rl.Vector2{X: float32(x), Y: float32(y)}, world.SL_INVENTORY)

				if err != nil {
					return
				}

				rl.DrawCircle(slotXOffset+int32(i.CellSize)/2, slotYOffset+int32(i.CellSize)/2, 12, slotItem.Color)
				//Item stack quantity text
				if stack.StackSize > 1 {
					rl.DrawText(fmt.Sprintf("%d", stack.StackSize), slotXOffset+int32(i.CellSize)-10, slotYOffset+int32(i.CellSize)-10, 16, rl.Black)
				}
			}
		}
	}
}

func (i *Inventory) DrawHotbar() {
	// Hotbar background
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(i.Positioning.HotbarXOffset), Y: float32(i.Positioning.HotbarYOffset), Width: float32(i.HotbarSize)*float32(i.CellSize) + float32(i.Positioning.HotbarXPadding), Height: float32(i.CellSize) + float32(i.Positioning.HotbarYPadding)}, 0.1, 1, transparentWhite)

	for x := range i.HotbarSize {
		slotXOffset := int32(x*i.CellSize + i.Positioning.HotbarXOffset + i.Positioning.HotbarCellXOffset)
		slotYOffset := int32(i.Positioning.HotbarYOffset) + int32(i.Positioning.HotbarCellYOffset)

		// Hotbar slots
		rl.DrawRectangleLines(slotXOffset, slotYOffset, int32(i.CellSize), int32(i.CellSize), rl.Black)
		if i.SelectedHotbarSlot == x {
			rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(slotXOffset), Y: float32(slotYOffset), Width: float32(i.CellSize), Height: float32(i.CellSize)}, 3, rl.Black)
		}

		//Slot item
		if i.Hotbar[x] != 0 {
			slotItem, err := world.GetItemByItemId(i.Hotbar[x])
			if err != nil {
				return
			}

			rl.DrawCircle(slotXOffset+int32(i.CellSize)/2, slotYOffset+int32(i.CellSize)/2, 12, slotItem.Color)
		}
	}
}

func (i *Inventory) DrawHoldingItem() {
	if i.HoldingItem {
		item, err := world.GetItemByItemId(i.SelectedInventoryItem.itemId)
		if err != nil {
			fmt.Println(err)
			return
		}

		rl.DrawCircle(rl.GetMouseX(), rl.GetMouseY(), 12, item.Color)
	}
}

func (i *Inventory) AddItemToInventory(itemId world.ItemId) error {
	//Check if item already in inventory
	exists, gridPos := i.ItemExistsInsideInventory(itemId, world.SL_INVENTORY)

	item, err := world.GetItemByItemId(itemId)

	if err != nil {
		return err
	}

	//If it is add to the existing items ItemStack if not over max stack
	if exists {
		_, stack, err := world.GetItemStackAtInventorySlot(gridPos, world.SL_INVENTORY)

		if err != nil {
			return err
		}

		if stack.StackSize+1 > item.MaxStack {
			i.AddItemToOpenSlot(itemId)
			return nil
		}

		stack.StackSize++
		return nil
	}

	//Add item to available slot if not in inventory or if new stack
	i.AddItemToOpenSlot(itemId)

	return nil
}

func (i *Inventory) AddItemToOpenSlot(itemId world.ItemId) error {
	slot, err := i.FindOpenInventorySlot()

	if err != nil {
		return err
	}

	world.CreateNewItemStack(slot, itemId, world.SL_INVENTORY)
	i.MainInventory.ItemGrid[int(slot.X)][int(slot.Y)] = itemId
	return nil
}

func (i *Inventory) ItemExistsInsideInventory(itemId world.ItemId, section world.StackLocation) (itemExists bool, gridPosition rl.Vector2) {
	if section == world.SL_INVENTORY {
		for y := range i.MainInventory.Height {
			for x := range i.MainInventory.Width {
				if i.MainInventory.ItemGrid[x][y] == itemId {
					_, stack, _ := world.GetItemStackAtInventorySlot(rl.Vector2{X: float32(x), Y: float32(y)}, section)
					item, _ := world.GetItemByItemId(itemId)

					if stack.StackSize < item.MaxStack {
						return true, rl.Vector2{X: float32(x), Y: float32(y)}
					}
				}
			}
		}
	} else if section == world.SL_HOTBAR {
		for x := range i.HotbarSize {
			if i.Hotbar[x] == itemId {
				_, stack, _ := world.GetItemStackAtInventorySlot(rl.Vector2{X: float32(x), Y: float32(0)}, section)
				item, _ := world.GetItemByItemId(itemId)

				if stack.StackSize < item.MaxStack {
					return true, rl.Vector2{X: float32(x), Y: float32(0)}
				}
			}
		}
	}

	return false, rl.Vector2{}
}

func (i *Inventory) FindOpenInventorySlot() (rl.Vector2, error) {
	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			if i.MainInventory.ItemGrid[x][y] == 0 {
				return rl.Vector2{X: float32(x), Y: float32(y)}, nil
			}
		}
	}

	return rl.Vector2{}, errors.New("no-open-slot")
}

func (i *Inventory) AddItemToHotbar(slot int, itemId world.ItemId) {
	world.CreateNewItemStack(rl.Vector2{X: float32(slot), Y: 0}, itemId, world.SL_HOTBAR)
	i.Hotbar[slot] = itemId
}

func (i *Inventory) ResetHoldingItem() {
	i.SelectedInventoryItem.startPos = rl.Vector2Zero()
	i.SelectedInventoryItem.stackId = uuid.UUID{}
	i.SelectedInventoryItem.itemId = 0
	i.HoldingItem = false
}

func (i *Inventory) InputHandler() {
	//Toggle Inventory
	if rl.IsKeyPressed(rl.KeyTab) {
		i.Visible = !i.Visible
	}

	if i.Visible {
		i.InventoryInputHandler()
	}

	i.HotbarInputHandler()
}

func (i *Inventory) MoveItemToInventoryGridPosition(gridPos rl.Vector2, itemId world.ItemId, section world.StackLocation) {
	x := int(gridPos.X)
	y := int(gridPos.Y)

	if section == world.SL_INVENTORY {
		i.MainInventory.ItemGrid[x][y] = itemId
	} else if section == world.SL_HOTBAR {
		i.Hotbar[x] = itemId
	}
}

func (i *Inventory) InventoryInputHandler() {

	if !i.Hover.InventoryHovering {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && i.HoldingItem {
			i.MainInventory.ItemGrid[int(i.SelectedInventoryItem.startPos.X)][int(i.SelectedInventoryItem.startPos.Y)] = i.SelectedInventoryItem.itemId
			i.HoldingItem = false
			i.SelectedInventoryItem.startPos = rl.Vector2Zero()
			i.SelectedInventoryItem.itemId = 0
		}
	}

	offsetX := (rl.GetScreenWidth() - i.MainInventory.ScreenWidth) / 2
	offsetY := (rl.GetScreenHeight() - i.MainInventory.ScreenHeight) / 2

	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			slotX := x*i.CellSize + offsetX + i.Positioning.InventoryCellXOffset
			slotY := y*i.CellSize + offsetY + i.Positioning.InventoryCellYOffset

			if rl.GetMouseX() >= int32(slotX) && rl.GetMouseX() <= int32(slotX)+int32(i.CellSize) {
				if rl.GetMouseY() >= int32(slotY) && rl.GetMouseY() <= int32(slotY)+int32(i.CellSize) {
					i.Hover.InventoryHovering = true
					i.Hover.InventoryHoverSlot.X = float32(x)
					i.Hover.InventoryHoverSlot.Y = float32(y)

					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						//Pickup or swap items if the slot isn't empty
						if i.MainInventory.ItemGrid[x][y] != 0 {
							//Pick up item if not currently holding something
							if !i.HoldingItem {
								i.PickupItemFromInventory(rl.Vector2{X: float32(x), Y: float32(y)}, world.SL_INVENTORY)
								return
							} else {
								//Current item stack at the selected main inventory slot
								currentStackId, currentItemStack, err := world.GetItemStackAtInventorySlot(rl.Vector2{X: float32(x), Y: float32(y)}, world.SL_INVENTORY)

								if err != nil {
									fmt.Println(err)
									return
								}

								//Move item stack from one inventory spot to another
								if i.SelectedInventoryItem.section == world.SL_INVENTORY {
									//Move stack and and add id to starting position of the currently held item
									world.MoveItemStack(currentStackId, i.SelectedInventoryItem.startPos, world.SL_INVENTORY)
									i.MoveItemToInventoryGridPosition(i.SelectedInventoryItem.startPos, currentItemStack.ItemId, world.SL_INVENTORY)

								} else if i.SelectedInventoryItem.section == world.SL_HOTBAR { //Move an item stack from the hotbar to the inventory
									//Move stack and and add id to starting position of the currently held item
									world.MoveItemStack(currentStackId, i.SelectedInventoryItem.startPos, world.SL_HOTBAR)
									i.MoveItemToInventoryGridPosition(i.SelectedInventoryItem.startPos, currentItemStack.ItemId, world.SL_HOTBAR)
								}

								//Move currently held item to selected slot
								i.MoveItemToInventoryGridPosition(rl.Vector2{X: float32(x), Y: float32(y)}, i.SelectedInventoryItem.itemId, world.SL_INVENTORY)
								world.MoveItemStack(i.SelectedInventoryItem.stackId, rl.Vector2{X: float32(x), Y: float32(y)}, world.SL_INVENTORY)

								i.ResetHoldingItem()
							}
						} else {
							if i.HoldingItem {
								stackId, _, err := world.GetItemStackAtInventorySlot(i.SelectedInventoryItem.startPos, i.SelectedInventoryItem.section)

								if err != nil {
									fmt.Println(err)
									return
								}

								newLocation := rl.Vector2{X: float32(x), Y: float32(y)}

								world.MoveItemStack(stackId, newLocation, world.SL_INVENTORY)
								i.MoveItemToInventoryGridPosition(rl.Vector2{X: float32(x), Y: float32(y)}, i.SelectedInventoryItem.itemId, world.SL_INVENTORY)

								i.ResetHoldingItem()
							}
						}
					}
				}
			}
		}
	}

}

func (i *Inventory) HotbarInputHandler() {
	//Hotbar scrollwheel input
	scrollwheel := rl.GetMouseWheelMove()

	if scrollwheel > 0 {
		i.SelectedHotbarSlot++
		if i.SelectedHotbarSlot > 8 {
			i.SelectedHotbarSlot = 8
		}
	} else if scrollwheel < 0 {
		i.SelectedHotbarSlot--
		if i.SelectedHotbarSlot < 0 {
			i.SelectedHotbarSlot = 0
		}
	}

	//Hotbar shortcut buttons
	if !i.Visible {
		//Hotbar Shortcuts
		if rl.IsKeyPressed(rl.KeyOne) {
			i.SelectedHotbarSlot = 0
		} else if rl.IsKeyPressed(rl.KeyTwo) {
			i.SelectedHotbarSlot = 1
		} else if rl.IsKeyPressed(rl.KeyThree) {
			i.SelectedHotbarSlot = 2
		} else if rl.IsKeyPressed(rl.KeyFour) {
			i.SelectedHotbarSlot = 3
		} else if rl.IsKeyPressed(rl.KeyFive) {
			i.SelectedHotbarSlot = 4
		} else if rl.IsKeyPressed(rl.KeySix) {
			i.SelectedHotbarSlot = 5
		} else if rl.IsKeyPressed(rl.KeySeven) {
			i.SelectedHotbarSlot = 6
		} else if rl.IsKeyPressed(rl.KeyEight) {
			i.SelectedHotbarSlot = 7
		} else if rl.IsKeyPressed(rl.KeyNine) {
			i.SelectedHotbarSlot = 8
		}
	}

	//Hotbar intersection detection
	for x := range i.HotbarSize {
		slotX := int32(x*i.CellSize + i.Positioning.HotbarXOffset + i.Positioning.HotbarCellXOffset)
		slotY := int32(i.Positioning.HotbarYOffset) + int32(i.Positioning.HotbarCellYOffset)

		if rl.GetMouseX() >= int32(slotX) && rl.GetMouseX() <= int32(slotX)+int32(i.CellSize) {
			if rl.GetMouseY() >= int32(slotY) && rl.GetMouseY() <= int32(slotY)+int32(i.CellSize) {
				i.Hover.HotbarHovering = true
				i.Hover.HotbarHoverSlot.X = float32(x)
				i.Hover.HotbarHoverSlot.Y = 0

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if i.Hotbar[x] != 0 {
						//Pick up item if not currently holding something
						if !i.HoldingItem {
							i.PickupItemFromInventory(rl.Vector2{X: float32(x), Y: float32(0)}, world.SL_HOTBAR)
							return
						} else {
							//Current item stack at the selected hotbar slot
							currentStackId, currentItemStack, err := world.GetItemStackAtInventorySlot(rl.Vector2{X: float32(x), Y: float32(0)}, world.SL_HOTBAR)

							if err != nil {
								fmt.Println(err)
								return
							}

							//Move item stack from an inventory slot to the hotbar
							if i.SelectedInventoryItem.section == world.SL_INVENTORY {
								//Move stack and and add id to starting position of the currently held item
								world.MoveItemStack(currentStackId, i.SelectedInventoryItem.startPos, world.SL_INVENTORY)
								i.MoveItemToInventoryGridPosition(i.SelectedInventoryItem.startPos, currentItemStack.ItemId, world.SL_INVENTORY)

							} else if i.SelectedInventoryItem.section == world.SL_HOTBAR { //Move an item stack from one hotbar slot to another
								//Move stack and and add id to starting position of the currently held item
								world.MoveItemStack(currentStackId, i.SelectedInventoryItem.startPos, world.SL_HOTBAR)
								i.MoveItemToInventoryGridPosition(i.SelectedInventoryItem.startPos, currentItemStack.ItemId, world.SL_HOTBAR)
							}

							//Move currently held item to selected slot
							i.MoveItemToInventoryGridPosition(rl.Vector2{X: float32(x), Y: float32(0)}, i.SelectedInventoryItem.itemId, world.SL_HOTBAR)
							world.MoveItemStack(i.SelectedInventoryItem.stackId, rl.Vector2{X: float32(x), Y: float32(0)}, world.SL_HOTBAR)

							i.ResetHoldingItem()
						}
					} else {
						if i.HoldingItem {
							stackId, _, err := world.GetItemStackAtInventorySlot(i.SelectedInventoryItem.startPos, i.SelectedInventoryItem.section)

							if err != nil {
								fmt.Println(err)
								return
							}

							newLocation := rl.Vector2{X: float32(x), Y: float32(0)}

							world.MoveItemStack(stackId, newLocation, world.SL_HOTBAR)
							i.MoveItemToInventoryGridPosition(rl.Vector2{X: float32(x), Y: float32(0)}, i.SelectedInventoryItem.itemId, world.SL_HOTBAR)

							i.ResetHoldingItem()
						}
					}
				}
			}
		}
	}

}

func (i *Inventory) PickupItemFromInventory(position rl.Vector2, section world.StackLocation) {
	stackId, stack, err := world.GetItemStackAtInventorySlot(position, section)

	if err != nil {
		fmt.Println(err)
		return
	}

	i.SelectedInventoryItem.stackId = stackId
	i.SelectedInventoryItem.itemId = stack.ItemId
	i.SelectedInventoryItem.section = section
	i.SelectedInventoryItem.startPos = position

	if section == world.SL_INVENTORY {
		i.MainInventory.ItemGrid[int(position.X)][int(position.Y)] = 0
	} else if section == world.SL_HOTBAR {
		i.Hotbar[int(position.X)] = 0
	}

	i.HoldingItem = true
}

func (i *Inventory) DrawInventoryToConsole() {
	fmt.Println("###INVENTORY###")
	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			fmt.Print(i.MainInventory.ItemGrid[x][y])
		}
		fmt.Print("\n")
	}
}

func (i *Inventory) DrawHotbarToConsole() {
	fmt.Println("###HOTBAR###")
	for x := range i.HotbarSize {
		fmt.Print(i.Hotbar[x])
	}
	fmt.Print("\n")
}

func (i *Inventory) ClearInventory() {
	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			i.MainInventory.ItemGrid[x][y] = 0
		}
	}
}
func (i *Inventory) ClearHotbar() {
	for x := range i.HotbarSize {
		i.Hotbar[x] = 0
	}
}
