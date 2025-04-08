package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type FacingDirection int32

const (
	UP    FacingDirection = 0
	DOWN  FacingDirection = 1
	LEFT  FacingDirection = 2
	RIGHT FacingDirection = 3
)

type Player struct {
	Position  rl.Vector2
	Health    int
	Speed     float32
	Width     int
	Height    int
	Direction FacingDirection
	Moving    bool
}

func (p *Player) Init() {

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
		p.Direction = UP
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= p.Speed
		p.Direction = LEFT
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += p.Speed
		p.Direction = DOWN
		p.Moving = true
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += p.Speed
		p.Direction = RIGHT
		p.Moving = true
	} else {
		if p.Moving {
			p.Moving = false
		}
	}
}
func (p *Player) Draw() {
	rl.DrawRectangleRounded(rl.Rectangle{Width: float32(p.Width), Height: float32(p.Height), X: p.Position.X, Y: p.Position.Y}, 0.1, 1, rl.Yellow)
}

func (p *Player) MoveToWorldPosition(position rl.Vector2) {
	p.Position = position
}
