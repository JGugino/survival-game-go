package main

import (
	"fmt"
	"math"

	"github.com/JGugino/survival-game-go/entities"
	"github.com/JGugino/survival-game-go/handlers"
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Generator       *handlers.WorldGenerator
	Camera          *rl.Camera2D
	CurrentPlayer   *entities.Player
	PlayerInventory *handlers.Inventory
	DebugPanel      *handlers.Debug
}

func (g *Game) Init() {
	//Load Textures
	utils.LoadNewTextureMap("world-tiles", "assets/world_tiles.png", 16)
	utils.LoadNewTextureMap("world-objects", "assets/world_objects.png", 16)
	utils.LoadNewTextureMap("items", "assets/items.png", 16)
	utils.LoadNewTextureMap("player", "assets/player.png", 16)

	//Initialize all game items
	utils.InitItemMap()

	//Generate a new world and spawn point
	g.Generator.GenerateNewWorldNoiseImage(int(rl.GetRandomValue(0, 20)), int(rl.GetRandomValue(0, 20)), 2)
	g.Generator.InitWorldGrid()
	g.Generator.SelectValidSpawnPoint()

	//Initalizes the inventory and hotbar to an empty state, also inits crafting recipes
	g.PlayerInventory.InitInventory()

	//Moves the player to the world's spawn point
	g.CurrentPlayer.MoveToWorldPosition(g.Generator.SpawnPoint)
}

func (g *Game) CleanUp() {
	utils.UnloadTextureMaps()
}

func (g *Game) HandleInput() {
	g.CurrentPlayer.HandleInput()
	g.PlayerInventory.InputHandler()
	g.DebugPanel.HandleInput()

	if rl.IsKeyPressed(rl.KeyP) {
		err := g.PlayerInventory.CraftItem("pickaxe")

		if err != nil {
			return
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		worldPos := rl.GetScreenToWorld2D(rl.Vector2{X: float32(math.Round(float64(rl.GetMouseX()))), Y: float32(float64(rl.GetMouseY()))}, (*g.Camera))

		_, obj, err := g.Generator.ObjectManager.GetObjectAtWorldPosition(worldPos)

		if err != nil {
			return
		}

		activeHotbarSlotItem, _ := utils.GetItemByItemId(g.PlayerInventory.Hotbar[g.PlayerInventory.SelectedHotbarSlot])

		if obj.Mineable {
			if g.PlayerInventory.Hotbar[g.PlayerInventory.SelectedHotbarSlot] == 0 {
				if obj.Level == utils.ML_LOW {
					err = g.Generator.ObjectManager.DamageObject(obj.Id, 1)

					if err != nil {
						fmt.Println(err)
						return
					}
				}
			} else if activeHotbarSlotItem.MineLevel >= obj.Level {
				err = g.Generator.ObjectManager.DamageObject(obj.Id, activeHotbarSlotItem.MineDamage)

				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}

func (g *Game) Update() {

	g.CurrentPlayer.Update(g.Generator.MapWidth*g.Generator.CellSize, g.Generator.MapHeight*g.Generator.CellSize)
	g.DebugPanel.Update()

	g.Camera.Target = g.CurrentPlayer.Position

	for _, i := range utils.WorldItems {

		colliding, _ := i.HandleItemCollision(rl.Rectangle{X: g.CurrentPlayer.Position.X, Y: g.CurrentPlayer.Position.Y, Width: float32(g.CurrentPlayer.Width), Height: float32(g.CurrentPlayer.Height)})

		if colliding {
			err := g.PlayerInventory.AddItemToInventory(i.Item.Id)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = utils.RemoveWorldItem(i.Id)

			if err != nil {
				return
			}
		}
	}

	g.limitCamera()

	if g.DebugPanel.WireMode {
		rl.EnableWireMode()
	} else if !g.DebugPanel.WireMode {
		rl.DisableWireMode()
	}
}

func (g *Game) Draw() {

	rl.BeginMode2D(*g.Camera)

	g.Generator.DrawWorld(g.Camera, &g.PlayerInventory.Visible)
	g.CurrentPlayer.Draw()

	g.PlayerInventory.DrawSelectedHotbarItem(g.CurrentPlayer.Position)

	utils.DrawWorldItems()

	rl.EndMode2D()

	g.PlayerInventory.DrawHotbar()

	if g.PlayerInventory.Visible {
		g.PlayerInventory.DrawInventory()
	}

	g.PlayerInventory.DrawHoldingItem()

	if g.DebugPanel.DebugOpen {
		g.DebugPanel.Draw()
	}
}

func (g *Game) limitCamera() {
	// Gets the total island width and height
	worldWidth := g.Generator.MapWidth * g.Generator.CellSize
	worldHeight := g.Generator.MapHeight * g.Generator.CellSize

	halfWindowWidth := rl.GetScreenWidth() / 2
	halfWindowHeight := rl.GetScreenHeight() / 2

	targetX := g.Camera.Target.X
	targetY := g.Camera.Target.Y

	// log.Info(std::to_string(targetY));
	// log.Info(std::to_string(0 + (halfWindowWidth - player.GetWidth()) - 66));

	// TODO: Make sure camera goes fully to edge of the map.
	//  Ensures the camera doesn't go outside the map
	if targetX <= float32((halfWindowWidth-g.CurrentPlayer.Width)-66) {
		// log.Warn("Camera Outside - Left");
		g.Camera.Target.X = float32((halfWindowWidth - g.CurrentPlayer.Width) - 66)
	}
	if targetX+float32(halfWindowWidth) > float32(worldWidth) {
		// log.Warn("Camera Outside - Right");
		g.Camera.Target.X = float32(worldWidth - halfWindowWidth)
	}
	if targetY-float32(halfWindowHeight) < 0 {
		// log.Warn("Camera Outside - Up");
		g.Camera.Target.Y = float32(halfWindowHeight)
	}
	if targetY+float32(halfWindowHeight) > float32(worldHeight) {
		// log.Warn("Camera Outside - Down");
		g.Camera.Target.Y = float32(worldHeight - halfWindowHeight)
	}
}
