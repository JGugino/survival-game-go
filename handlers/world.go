package handlers

import (
	"fmt"

	"github.com/JGugino/survival-game-go/utils"
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
	ObjectManager   Objects
}

func (g *WorldGenerator) InitWorldGrid() {
	g.ObjectManager.InitObjectGrid(g.MapWidth, g.MapHeight, g.CellSize)

	for y := range g.MapHeight {
		for x := range g.MapWidth {
			pixelColor := rl.GetImageColor(g.WorldNoiseImage, int32(x), int32(y))
			tileType := g.GetPixelType(pixelColor)

			if tileType == STONE {
				test := rl.GetRandomValue(0, 40)
				if test > 5 && test <= 18 {
					rock, err := utils.GetItemByName("rock")

					if err != nil {
						return
					}

					g.ObjectManager.CreateNewObject(ROCK, rl.Vector2{X: float32(x), Y: float32(y)}, 60, rl.DarkGray, false, true, utils.ML_MED, *rock)
				} else if test >= 39 {
					utils.NewWorldItem(utils.I_ROCK, rl.Vector2{X: float32(x * g.CellSize), Y: float32(y * g.CellSize)})
				}
			}

			if tileType == GRASS {
				test := rl.GetRandomValue(0, 20)
				if test > 18 {
					wood, err := utils.GetItemByName("wood")

					if err != nil {
						return
					}

					g.ObjectManager.CreateNewObject(BUSH, rl.Vector2{X: float32(x), Y: float32(y)}, 20, rl.DarkGray, false, true, utils.ML_LOW, *wood)
				}
			}

			g.WorldGrid[x][y] = int(tileType)
		}
	}
}

func (g *WorldGenerator) DrawWorld(camera *rl.Camera2D, inventoryVisible *bool) {
	tMap, err := utils.GetTextureMap("world-tiles")

	if err != nil {
		fmt.Println(err)
		return
	}

	for y := range g.MapHeight {
		for x := range g.MapWidth {
			//Draw world tiles
			pixelType := g.GetPixelType(rl.GetImageColor(g.WorldNoiseImage, int32(x), int32(y)))

			cellPosition := rl.Vector2{X: float32(x * g.CellSize), Y: float32(y * g.CellSize)}

			switch pixelType {
			case WATER:
				tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 0, Y: 0}, cellPosition, g.CellSize)
				break
			case GRASS:
				tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 1, Y: 0}, cellPosition, g.CellSize)
				break
			case SAND:
				tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 2, Y: 0}, cellPosition, g.CellSize)
				break
			case STONE:
				tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 3, Y: 0}, cellPosition, g.CellSize)
				break
			case SNOW:
				tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 4, Y: 0}, cellPosition, g.CellSize)
				break
			}

			//Draw object tiles
			g.ObjectManager.DrawObjects(x, y, camera)

			//Draws outline around tile that is being hovered over by the mouse, doesn't show if inventory is open
			if !*inventoryVisible {
				worldPos := rl.GetScreenToWorld2D(rl.Vector2{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY())}, (*camera))

				if int32(worldPos.X) >= int32(cellPosition.X) && int32(worldPos.X) <= int32(cellPosition.X)+int32(g.CellSize) {
					if int32(worldPos.Y) >= int32(cellPosition.Y) && int32(worldPos.Y) <= int32(cellPosition.Y)+int32(g.CellSize) {
						rect := rl.Rectangle{X: cellPosition.X, Y: cellPosition.Y, Width: float32(g.CellSize), Height: float32(g.CellSize)}
						rl.DrawRectangleLinesEx(rect, 2, rl.Color{R: 40, G: 40, B: 40, A: 220})
					}
				}
			}

		}
	}
}

func (g *WorldGenerator) DrawWorldToConsole() {
	fmt.Println("###WORLD###")
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
