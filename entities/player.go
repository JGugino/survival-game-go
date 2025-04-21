package entities

import (
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position         rl.Vector2
	Health           int
	MaxHealth        int
	Speed            float32
	Width            int
	Height           int
	Direction        utils.FacingDirection
	HoldingLocations HoldingLocations
	Moving           bool
}

type HoldingLocations struct {
	ForwardHold rl.Vector2
	LeftHold    rl.Vector2
	RightHold   rl.Vector2
}

func (p *Player) Init() {
	p.HoldingLocations.ForwardHold = rl.Vector2{X: p.Position.X, Y: p.Position.Y}
	p.HoldingLocations.LeftHold = rl.Vector2{X: p.Position.X, Y: p.Position.Y}
	p.HoldingLocations.RightHold = rl.Vector2{X: p.Position.X, Y: p.Position.Y}
}

func (p *Player) Update(worldWidth int, worldHeight int) {
	p.HoldingLocations.ForwardHold = rl.Vector2{X: p.Position.X + float32(p.Width)/2 + 6, Y: p.Position.Y + 10}
	p.HoldingLocations.LeftHold = rl.Vector2{X: p.Position.X + 14, Y: p.Position.Y + 6}
	p.HoldingLocations.RightHold = rl.Vector2{X: p.Position.X + 14, Y: p.Position.Y + 6}

	//Limit X
	if p.Position.X <= 0 {
		p.Position.X = 0
	}
	if p.Position.X >= float32(worldWidth)-float32(p.Width) {
		p.Position.X = float32(worldWidth) - float32(p.Width)
	}

	//Limit Y
	if p.Position.Y <= 0 {
		p.Position.Y = 0
	}
	if p.Position.Y >= float32(worldHeight)-float32(p.Width) {
		p.Position.Y = float32(worldHeight) - float32(p.Width)
	}

}
func (p *Player) HandleInput() {
	if rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= p.Speed
		p.Direction = utils.UP
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= p.Speed
		p.Direction = utils.LEFT
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += p.Speed
		p.Direction = utils.DOWN
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += p.Speed
		p.Direction = utils.RIGHT
		p.Moving = true
	} else {
		if p.Moving {
			p.Moving = false
		}
	}
}
func (p *Player) Draw() {
	tMap, err := utils.GetTextureMap("player")

	if err != nil {
		return
	}

	position := rl.Vector2{X: p.Position.X, Y: p.Position.Y}
	if p.Direction == utils.DOWN {
		tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 0, Y: 0}, position, p.Width)
	} else if p.Direction == utils.UP {
		tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 1, Y: 0}, position, p.Width)
	} else if p.Direction == utils.LEFT {
		tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 2, Y: 0}, position, p.Width)
	} else if p.Direction == utils.RIGHT {
		tMap.DrawTextureAtPositionWithScaling(rl.Vector2{X: 3, Y: 0}, position, p.Width)
	}
}

func (p *Player) DrawHealthBar(barPosition rl.Vector2, textPosition rl.Vector2) {
	var healthBarWidth int32 = 200
	var healthBarHeight int32 = 40
	rl.DrawRectangle(int32(barPosition.X), int32(barPosition.Y), healthBarWidth, healthBarHeight, rl.White)
	rl.DrawRectangle(int32(barPosition.X), int32(barPosition.Y), healthBarWidth*int32(p.Health)/int32(p.MaxHealth), healthBarHeight, rl.Red)
	rl.DrawText("Health", int32(textPosition.X), int32(textPosition.Y), 20, rl.Black)
}

func (p *Player) MoveToWorldPosition(position rl.Vector2) {
	p.Position = position
}
