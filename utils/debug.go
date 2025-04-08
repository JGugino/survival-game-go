package utils

import (
	"fmt"

	"github.com/JGugino/survival-game-go/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Debug struct {
	Generator        *WorldGenerator
	CurrentPlayer    *entities.Player
	DebugOpen        bool
	DebugFontSize    float32
	DebugTextSpacing float32
	OffsetX          float32
	OffsetY          float32
	ContentOffsetX   float32
	ContentOffsetY   float32
	StandingTile     string
}

func (d *Debug) Update() {
	switch d.Generator.GetTileTypeAtWorldPosition(int(d.CurrentPlayer.Position.X+(float32(d.CurrentPlayer.Width)/2)), int(d.CurrentPlayer.Position.Y+(float32(d.CurrentPlayer.Height)/2))) {
	case 0:
		d.StandingTile = "Water"
		break
	case 1:
		d.StandingTile = "Sand"
		break
	case 2:
		d.StandingTile = "Grass"
		break
	case 3:
		d.StandingTile = "Stone"
		break
	case 4:
		d.StandingTile = "Snow"
		break

	default:
		break
	}
}

func (d *Debug) Draw() {
	rl.DrawRectangleRounded(rl.Rectangle{X: d.OffsetX, Y: d.OffsetY, Width: 400, Height: 400}, 0.1, 1, rl.Color{R: 255, G: 255, B: 255, A: 200})
	// FPS
	rl.DrawFPS(int32(d.ContentOffsetX), int32(d.ContentOffsetY))

	// World Spawn
	rl.DrawText(fmt.Sprintf("Spawn - X: %.0f / Y: %.0f", d.Generator.SpawnPoint.X, d.Generator.SpawnPoint.Y), int32(d.ContentOffsetX), int32(d.ContentOffsetY+d.DebugTextSpacing), int32(d.DebugFontSize), rl.Black)

	// Player Position
	rl.DrawText(fmt.Sprintf("Player - X: %.0f / Y: %.0f", d.CurrentPlayer.Position.X, d.CurrentPlayer.Position.Y), int32(d.ContentOffsetX), int32(d.ContentOffsetY)+int32(d.DebugTextSpacing)*2, int32(d.DebugFontSize), rl.Black)

	// Mouse Position
	rl.DrawText(fmt.Sprintf("Mouse - X: %d / Y: %d", rl.GetMouseX(), rl.GetMouseY()), int32(d.ContentOffsetX), int32(d.ContentOffsetY)+int32(d.DebugTextSpacing)*3, int32(d.DebugFontSize), rl.Black)

	//Standing Tile
	rl.DrawText(fmt.Sprintf("Tile - %s", d.StandingTile), int32(d.ContentOffsetX), int32(d.ContentOffsetY)+int32(d.DebugTextSpacing)*4, int32(d.DebugFontSize), rl.Black)
}

func (d *Debug) HandleInput() {
	//Toggle debug window
	if rl.IsKeyPressed(rl.KeyF3) {
		d.DebugOpen = !d.DebugOpen
	}
}
