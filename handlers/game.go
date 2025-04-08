package handlers

import (
	"github.com/JGugino/survival-game-go/entities"
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Generator       *utils.WorldGenerator
	Camera          *rl.Camera2D
	CurrentPlayer   *entities.Player
	PlayerInventory *Inventory
	DebugPanel      *utils.Debug
}

func (g *Game) Init() {
	g.Generator.GenerateNewWorldNoiseImage(int(rl.GetRandomValue(0, 20)), int(rl.GetRandomValue(0, 20)), 2)
	g.Generator.InitWorldGrid()
	g.Generator.SelectValidSpawnPoint()

	//Clears the inventory and hotbar grids to all zeros
	g.PlayerInventory.ClearInventory()
	g.PlayerInventory.ClearHotbar()

	//Moves the player to the world's spawn point
	g.CurrentPlayer.MoveToWorldPosition(g.Generator.SpawnPoint)
}

func (g *Game) HandleInput() {
	g.CurrentPlayer.HandleInput()
	g.PlayerInventory.InputHandler()
	g.DebugPanel.HandleInput()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		// worldPos := rl.GetScreenToWorldRay(rl.Vector2{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY())}, (*g.Camera))

		// id, obj := g.Generator.ObjectManager.GetObjectAtWorldPosition(worldPos.Position)

		// fmt.Println(id)
		// fmt.Println(obj)
	}
}

func (g *Game) Update() {

	g.CurrentPlayer.Update(g.Generator.MapWidth*g.Generator.CellSize, g.Generator.MapHeight*g.Generator.CellSize)
	g.DebugPanel.Update()

	g.Camera.Target = g.CurrentPlayer.Position

	g.limitCamera()
}

func (g *Game) Draw() {

	rl.BeginMode2D(*g.Camera)

	g.Generator.DrawWorld()
	g.CurrentPlayer.Draw()

	rl.EndMode2D()

	g.PlayerInventory.DrawHotbar()

	if g.PlayerInventory.Visible {
		g.PlayerInventory.DrawInventory()
	}

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
