package main

import (
	"github.com/JGugino/survival-game-go/entities"
	"github.com/JGugino/survival-game-go/handlers"
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

const (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 800
	WINDOW_TITLE  = "Survival Game - Raylib Go"
	TARGET_FPS    = 60

	DEFAULT_CELL_SIZE  = 40
	DEFAULT_MAP_WIDTH  = 64
	DEFAULT_MAP_HEIGHT = 64
)

//TODO: Add window close requesting and saving of data on close
// var windowCloseRequested bool = false
// var windowClose bool = false

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, WINDOW_TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(TARGET_FPS)

	//TODO: Work on resizable window
	//rl.SetWindowState(rl.FlagWindowResizable)

	worldGenerator := handlers.WorldGenerator{
		CellSize:  DEFAULT_CELL_SIZE,
		MapWidth:  DEFAULT_MAP_WIDTH,
		MapHeight: DEFAULT_MAP_HEIGHT,
		ObjectManager: handlers.Objects{
			Objs: make(map[uuid.UUID]*handlers.Object, 0),
		},
	}

	player := entities.Player{
		Position:  rl.Vector2{X: worldGenerator.SpawnPoint.X, Y: worldGenerator.SpawnPoint.Y},
		Health:    100,
		MaxHealth: 100,
		Speed:     2,
		Width:     DEFAULT_CELL_SIZE,
		Height:    DEFAULT_CELL_SIZE,
		Direction: utils.DOWN,
		Moving:    false}

	mainCamera := rl.Camera2D{
		Offset:   rl.Vector2{X: WINDOW_WIDTH / 2, Y: WINDOW_HEIGHT / 2},
		Target:   player.Position,
		Rotation: 0,
		Zoom:     1.4,
	}

	inv := handlers.Inventory{
		CraftingHandler: &handlers.Crafting{
			CraftingRecipes: make(map[string]handlers.CraftingRecipe, 0),
		},
		HotbarSize:         9,
		CellSize:           DEFAULT_CELL_SIZE,
		Visible:            false,
		SelectedHotbarSlot: 0,
		Hover: struct {
			InventoryHovering  bool
			InventoryHoverSlot rl.Vector2
			HotbarHovering     bool
			HotbarHoverSlot    rl.Vector2
		}{
			InventoryHovering:  false,
			InventoryHoverSlot: rl.Vector2Zero(),
			HotbarHovering:     false,
			HotbarHoverSlot:    rl.Vector2Zero(),
		},
		MainInventory: handlers.Container{Width: 9, Height: 4, ScreenWidth: 9 * DEFAULT_CELL_SIZE, ScreenHeight: 4 * DEFAULT_CELL_SIZE},
		Positioning: handlers.InventoryPosition{
			InventoryYOffset:     40,
			InventoryXPadding:    10,
			InventoryYPadding:    200,
			InventoryCellXOffset: 5,
			InventoryCellYOffset: 130,
			HotbarYOffset:        10,
			HotbarXOffset:        10,
			HotbarXPadding:       10,
			HotbarYPadding:       10,
			HotbarCellXOffset:    5,
			HotbarCellYOffset:    5,
		},
	}

	debug := handlers.Debug{
		DebugOpen:        false,
		DebugFontSize:    19,
		DebugTextSpacing: 30,
		OffsetX:          WINDOW_WIDTH - 440,
		OffsetY:          10,
		ContentOffsetX:   WINDOW_WIDTH - 440 + 20,
		ContentOffsetY:   20,
		Generator:        &worldGenerator,
		CurrentPlayer:    &player,
	}

	game := Game{
		Generator:       &worldGenerator,
		Camera:          &mainCamera,
		CurrentPlayer:   &player,
		PlayerInventory: &inv,
		DebugPanel:      &debug,
	}

	game.Init()

	for !rl.WindowShouldClose() {
		game.HandleInput()

		game.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		game.Draw()

		rl.EndDrawing()
	}
	game.CleanUp()
}
