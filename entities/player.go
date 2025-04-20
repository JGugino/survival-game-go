package entities

import (
	"github.com/JGugino/survival-game-go/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position  rl.Vector2
	Health    int
	MaxHealth int
	Speed     float32
	Width     int
	Height    int
	Direction utils.FacingDirection
	Moving    bool
}

func (p *Player) Update(worldWidth int, worldHeight int) {
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

func (p *Player) MoveToWorldPosition(position rl.Vector2) {
	p.Position = position
}
