package world

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureIdentifier struct {
	TextureMapId      string
	TexturePosition   rl.Vector2
	ContainerDrawSize int
	WorldDrawSize     int
}

type TextureMap struct {
	Id          string
	Width       int
	Height      int
	TextureSize int
	Texture     rl.Texture2D
}

func (t *TextureMap) DrawTextureAtPosition(texturePosition rl.Vector2, drawLocation rl.Vector2) {
	rl.DrawTextureRec(t.Texture, rl.Rectangle{X: float32(int(texturePosition.X) * t.TextureSize), Y: float32(int(texturePosition.Y) * t.TextureSize), Width: float32(t.TextureSize), Height: float32(t.TextureSize)}, drawLocation, rl.White)
}
func (t *TextureMap) DrawTextureAtPositionWithScalingPro(texturePosition rl.Vector2, drawLocation rl.Vector2, drawWidth, drawHeight int) {
	textureSelectRect := rl.Rectangle{X: float32(int(texturePosition.X) * t.TextureSize), Y: float32(int(texturePosition.Y) * t.TextureSize), Width: float32(t.TextureSize), Height: float32(t.TextureSize)}
	textureDrawRect := rl.Rectangle{X: drawLocation.X, Y: drawLocation.Y, Width: float32(drawWidth), Height: float32(drawHeight)}

	rl.DrawTexturePro(t.Texture, textureSelectRect, textureDrawRect, rl.Vector2{X: float32(drawWidth) / 2, Y: float32(drawHeight) / 2}, 0, rl.White)
}
func (t *TextureMap) DrawTextureAtPositionWithScaling(texturePosition rl.Vector2, drawLocation rl.Vector2, drawSize int) {
	textureSelectRect := rl.Rectangle{X: float32(int(texturePosition.X) * t.TextureSize), Y: float32(int(texturePosition.Y) * t.TextureSize), Width: float32(t.TextureSize), Height: float32(t.TextureSize)}
	textureDrawRect := rl.Rectangle{X: drawLocation.X, Y: drawLocation.Y, Width: float32(drawSize), Height: float32(drawSize)}

	rl.DrawTexturePro(t.Texture, textureSelectRect, textureDrawRect, rl.Vector2{X: float32(drawSize) / 2, Y: float32(drawSize) / 2}, 0, rl.White)
}

var loadedTextureMaps map[string]*TextureMap = make(map[string]*TextureMap, 0)

func GetTextureMap(id string) (*TextureMap, error) {
	tMap, ok := loadedTextureMaps[id]

	if !ok {
		return &TextureMap{}, errors.New("no-texture-map")
	}

	return tMap, nil
}

func LoadNewTextureMap(mapId string, mapPath string, textureSize int) error {
	if _, ok := loadedTextureMaps[mapId]; !ok {
		tMap := rl.LoadTexture(mapPath)
		loadedTextureMaps[mapId] = &TextureMap{
			Id:          mapId,
			Width:       int(tMap.Width),
			Height:      int(tMap.Height),
			TextureSize: textureSize,
			Texture:     tMap,
		}

		return nil
	}

	return errors.New("texture-map-id-exists")
}

func UnloadTextureMaps() {
	if len(loadedTextureMaps) <= 0 {
		return
	}

	for _, m := range loadedTextureMaps {
		rl.UnloadTexture(m.Texture)
	}
}
