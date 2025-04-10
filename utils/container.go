package utils

import "github.com/JGugino/survival-game-go/world"

type Container struct {
	Width        int
	Height       int
	ScreenWidth  int
	ScreenHeight int
	ItemGrid     [9][4]world.ItemId
}
