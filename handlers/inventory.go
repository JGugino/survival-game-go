package handlers

import (
	"fmt"

	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	HotbarSize         int
	MainInventory      utils.Container
	Hotbar             [9]utils.ItemId
	CellSize           int
	SelectedHotbarSlot int
	Visible            bool
	Positioning        InventoryPosition
	Hovering           bool
	HoverSlot          struct {
		X int
		Y int
	}
}

var transparentWhite rl.Color = rl.Color{R: 255, G: 255, B: 255, A: 230}

func (i *Inventory) DrawInventory() {
	offsetX := (rl.GetScreenWidth() - i.MainInventory.ScreenWidth) / 2
	offsetY := (rl.GetScreenHeight() - i.MainInventory.ScreenHeight) / 2

	// Inventory background
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(offsetX), Y: float32(offsetY - i.Positioning.InventoryYOffset), Width: float32(i.MainInventory.ScreenWidth + i.Positioning.InventoryXPadding), Height: float32(i.MainInventory.ScreenHeight + i.Positioning.InventoryYPadding)}, 0.1, 1, transparentWhite)

	// Inventory slots

	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			rl.DrawRectangleLines(int32(x*i.CellSize+offsetX+i.Positioning.InventoryCellXOffset), int32(y*i.CellSize+offsetY+i.Positioning.InventoryCellYOffset), int32(i.CellSize), int32(i.CellSize), rl.Black)
			if i.HoverSlot.X == x && i.HoverSlot.Y == y {
				rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(x*i.CellSize + offsetX + i.Positioning.InventoryCellXOffset), Y: float32(y*i.CellSize + offsetY + i.Positioning.InventoryCellYOffset), Width: float32(i.CellSize), Height: float32(i.CellSize)}, 2, rl.Black)
			}
		}
	}
}

func (i *Inventory) DrawHotbar() {
	// Hotbar background
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(i.Positioning.HotbarXOffset), Y: float32(i.Positioning.HotbarYOffset), Width: float32(i.HotbarSize)*float32(i.CellSize) + float32(i.Positioning.HotbarXPadding), Height: float32(i.CellSize) + float32(i.Positioning.HotbarYPadding)}, 0.1, 1, transparentWhite)

	// Hotbar slots
	for x := range i.HotbarSize {
		rl.DrawRectangleLines(int32(x*i.CellSize+i.Positioning.HotbarXOffset+i.Positioning.HotbarCellXOffset), int32(i.Positioning.HotbarYOffset)+int32(i.Positioning.HotbarCellYOffset), int32(i.CellSize), int32(i.CellSize), rl.Black)
		if i.SelectedHotbarSlot == x {
			rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(x*i.CellSize + i.Positioning.HotbarXOffset + i.Positioning.HotbarCellXOffset), Y: float32(i.Positioning.HotbarYOffset) + float32(i.Positioning.HotbarCellYOffset), Width: float32(i.CellSize), Height: float32(i.CellSize)}, 3, rl.Black)
		}
	}
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

func (i *Inventory) InventoryInputHandler() {
	offsetX := (rl.GetScreenWidth() - i.MainInventory.ScreenWidth) / 2
	offsetY := (rl.GetScreenHeight() - i.MainInventory.ScreenHeight) / 2

	for y := range i.MainInventory.Height {
		for x := range i.MainInventory.Width {
			slotX := x*i.CellSize + offsetX + i.Positioning.InventoryCellXOffset
			slotY := y*i.CellSize + offsetY + i.Positioning.InventoryCellYOffset

			if rl.GetMouseX() >= int32(slotX) && rl.GetMouseX() <= int32(slotX)+int32(i.CellSize) {
				if rl.GetMouseY() >= int32(slotY) && rl.GetMouseY() <= int32(slotY)+int32(i.CellSize) {
					if !i.Hovering {
						i.Hovering = true
					}
					i.HoverSlot.X = x
					i.HoverSlot.Y = y
				}
			}

			if i.Hovering {
				i.Hovering = false
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
			i.MainInventory.ItemGrid[x][y] = utils.EMPTY
		}
	}
}
func (i *Inventory) ClearHotbar() {
	for x := range i.HotbarSize {
		i.Hotbar[x] = utils.EMPTY
	}
}
