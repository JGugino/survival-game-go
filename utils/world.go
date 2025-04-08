package utils

import (
	"fmt"

	"github.com/JGugino/survival-game-go/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileType int32

const (
	WATER TileType = 0
	SAND  TileType = 1
	GRASS TileType = 2
	STONE TileType = 3
	SNOW  TileType = 4
)

type WorldGenerator struct {
	CellSize        int
	MapWidth        int
	MapHeight       int
	WorldNoiseImage rl.Image
	WorldGrid       [64][64]int
	SpawnPoint      rl.Vector2
	ObjectManager   objects.Objects
}

func (g *WorldGenerator) InitWorldGrid() {
	g.ObjectManager.InitObjectGrid()
	for y := range g.MapHeight {
		for x := range g.MapWidth {
			pixelColor := rl.GetImageColor(g.WorldNoiseImage, int32(x), int32(y))
			tileType := g.GetPixelType(pixelColor)

			if tileType == STONE {
				test := rl.GetRandomValue(0, 6)
				if test >= 5 {
					g.ObjectManager.CreateNewObject(objects.ROCK, rl.Vector2{X: float32(x), Y: float32(y)}, 100, rl.DarkGray, false, true, objects.ML_MED, 1)
				}
			}

			g.WorldGrid[x][y] = int(tileType)

		}
	}
}

func (g *WorldGenerator) DrawWorld() {
	for y := range g.MapHeight {
		for x := range g.MapWidth {
			pixelColor := rl.GetImageColor(g.WorldNoiseImage, int32(x), int32(y))
			rl.DrawRectangle(int32(x*g.CellSize), int32(y*g.CellSize), int32(g.CellSize), int32(g.CellSize), g.GetColorForTileType(g.GetPixelType(pixelColor)))
		}
	}

	g.ObjectManager.DrawObjects()
}

func (g *WorldGenerator) DrawWorldToConsole() {
	for y := range g.MapHeight {
		for x := range g.MapWidth {
			tile := g.WorldGrid[x][y]

			fmt.Print(tile)
		}
		fmt.Print("\n")
	}
}

func (g *WorldGenerator) HandleInput() {

}

func (g *WorldGenerator) SelectValidSpawnPoint() {
	selectedX := rl.GetRandomValue(int32(g.CellSize), int32(g.MapWidth)-int32(g.CellSize))
	selectedY := rl.GetRandomValue(int32(g.CellSize), int32(g.MapHeight)-int32(g.CellSize))

	if g.IsWalkableTile(int(selectedX), int(selectedY)) {
		g.SpawnPoint = rl.Vector2{X: float32(selectedX * int32(g.CellSize)), Y: float32(selectedY * int32(g.CellSize))}
		return
	}

	g.SelectValidSpawnPoint()
}

func (g *WorldGenerator) IsWalkableTile(x int, y int) bool {
	if g.WorldGrid[x][y] == 0 {
		return false
	}

	return true
}

func (g *WorldGenerator) GetTileTypeAtWorldPosition(x int, y int) int {
	gridX := x / g.CellSize
	gridY := y / g.CellSize

	return g.WorldGrid[gridX][gridY]
}

func (g *WorldGenerator) GenerateNewWorldNoiseImage(offsetX int, offsetY int, scale float32) {
	g.WorldNoiseImage = *rl.GenImagePerlinNoise(g.MapWidth, g.MapHeight, offsetX, offsetY, scale)
}

func (g *WorldGenerator) GetPixelType(pixelColor rl.Color) TileType {
	start := 0
	increment := 51
	if pixelColor.R > uint8(start) && pixelColor.R <= uint8(start)+(uint8(increment)*2) { //WATER
		return WATER
	} else if pixelColor.R > uint8(start)+(uint8(increment)*2) && pixelColor.R <= uint8(start)+(uint8(increment)*2)+10 { //SAND
		return SAND
	} else if pixelColor.R > uint8(start)+(uint8(increment)*2)+10 && pixelColor.R <= uint8(start)+(uint8(increment)*3) { //GRASS
		return GRASS
	} else if pixelColor.R > uint8(start)+(uint8(increment)*3) && pixelColor.R <= uint8(start)+(uint8(increment)*4) { //STONE
		return STONE
	} else if pixelColor.R > uint8(start)+(uint8(increment)*4) && pixelColor.R <= uint8(start)+(uint8(increment)*5) { //SNOW
		return SNOW
	}

	return WATER
}

func (g *WorldGenerator) GetColorForTileType(tileType TileType) rl.Color {
	switch tileType {
	case WATER:
		return rl.Color{R: 75, G: 77, B: 242, A: 255}
	case SAND:
		return rl.Color{R: 242, G: 212, B: 135, A: 255}
	case GRASS:
		return rl.Color{R: 80, G: 170, B: 15, A: 255}
	case STONE:
		return rl.Color{R: 164, G: 165, B: 164, A: 255}
	case SNOW:
		return rl.Color{R: 239, G: 242, B: 237, A: 255}
	default:
		return rl.Color{R: 75, G: 77, B: 242, A: 255}
	}
}
