package utils

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScrollbox() {
	panelRec := rl.Rectangle{X: 20, Y: 40, Width: 200, Height: 150}
	panelContentRec := rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 300}
	panelView := rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 300}
	panelScroll := rl.Vector2{X: 99, Y: -20}
	//mouseCell := rl.Vector2{X: 0, Y: 0}

	showContentArea := false

	rl.DrawText(fmt.Sprintf("[%.1f, %.1f]", panelScroll.X, panelScroll.Y), 4, 4, 20, rl.Red)

	gui.ScrollPanel(panelRec, "", panelContentRec, &panelScroll, &panelView)

	rl.BeginScissorMode(int32(panelView.X), int32(panelView.Y), int32(panelView.Width), int32(panelView.Height))

	gui.Button(rl.Rectangle{X: panelContentRec.X + 5, Y: panelContentRec.Y + 5, Width: panelContentRec.Width - 10, Height: 80}, "Pickaxe")
	rl.EndScissorMode()

	if showContentArea {
		rl.DrawRectangle(
			int32(panelRec.X+panelScroll.X),
			int32(panelRec.Y+panelScroll.Y),
			int32(panelContentRec.Width),
			int32(panelContentRec.Height),
			rl.Fade(rl.Red, 0.1),
		)
	}

	DrawStyleEditControls()

	showContentArea = gui.CheckBox(rl.Rectangle{X: 565, Y: 80, Width: 20, Height: 20}, "SHOW CONTENT AREA", showContentArea)

	panelContentRec.Width = gui.SliderBar(rl.Rectangle{X: 590, Y: 385, Width: 145, Height: 15},
		"WIDTH",
		fmt.Sprintf("%.1f", panelContentRec.Width),
		panelContentRec.Width,
		1, 600)
	panelContentRec.Height = gui.SliderBar(rl.Rectangle{X: 590, Y: 410, Width: 145, Height: 15},
		"HEIGHT",
		fmt.Sprintf("%.1f", panelContentRec.Height),
		panelContentRec.Height, 1, 400)
}

// Draw and process scroll bar style edition controls
func DrawStyleEditControls() {
	// ScrollPanel style controls
	//----------------------------------------------------------
	gui.GroupBox(rl.Rectangle{X: 550, Y: 170, Width: 220, Height: 205}, "SCROLLBAR STYLE")

	var style int32

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.BORDER_WIDTH))
	gui.Label(rl.Rectangle{X: 555, Y: 195, Width: 110, Height: 10}, "BORDER_WIDTH")
	gui.Spinner(rl.Rectangle{X: 670, Y: 190, Width: 90, Height: 20}, "", &style, 0, 6, false)
	gui.SetStyle(gui.SCROLLBAR, gui.BORDER_WIDTH, int64(style))

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.ARROWS_SIZE))
	gui.Label(rl.Rectangle{X: 555, Y: 220, Width: 110, Height: 10}, "ARROWS_SIZE")
	gui.Spinner(rl.Rectangle{X: 670, Y: 215, Width: 90, Height: 20}, "", &style, 4, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.ARROWS_SIZE, int64(style))

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING))
	gui.Label(rl.Rectangle{X: 555, Y: 245, Width: 110, Height: 10}, "SLIDER_PADDING")
	gui.Spinner(rl.Rectangle{X: 670, Y: 240, Width: 90, Height: 20}, "", &style, 0, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING, int64(style))

	style = boolToint32(gui.CheckBox(rl.Rectangle{X: 565, Y: 280, Width: 20, Height: 20}, "ARROWS_VISIBLE", int32Tobool(int32(gui.GetStyle(gui.SCROLLBAR, gui.ARROWS_VISIBLE)))))
	gui.SetStyle(gui.SCROLLBAR, gui.ARROWS_VISIBLE, int64(style))

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING))
	gui.Label(rl.Rectangle{X: 555, Y: 325, Width: 110, Height: 10}, "SLIDER_PADDING")
	gui.Spinner(rl.Rectangle{X: 670, Y: 320, Width: 90, Height: 20}, "", &style, 0, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING, int64(style))

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_WIDTH))
	gui.Label(rl.Rectangle{X: 555, Y: 350, Width: 110, Height: 10}, "SLIDER_WIDTH")
	gui.Spinner(rl.Rectangle{X: 670, Y: 345, Width: 90, Height: 20}, "", &style, 2, 100, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_WIDTH, int64(style))

	var text string
	if gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE) == gui.SCROLLBAR_LEFT_SIDE {
		text = "SCROLLBAR: LEFT"
	} else {
		text = "SCROLLBAR: RIGHT"
	}
	style = boolToint32(gui.Toggle(rl.Rectangle{X: 560, Y: 110, Width: 200, Height: 35}, text, int32Tobool(int32(gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE)))))
	gui.SetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE, int64(style))
	//----------------------------------------------------------

	// ScrollBar style controls
	//----------------------------------------------------------
	gui.GroupBox(rl.Rectangle{X: 550, Y: 20, Width: 220, Height: 135}, "SCROLLPANEL STYLE")

	style = int32(gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_WIDTH))
	gui.Label(rl.Rectangle{X: 555, Y: 35, Width: 110, Height: 10}, "SCROLLBAR_WIDTH")
	gui.Spinner(rl.Rectangle{X: 670, Y: 30, Width: 90, Height: 20}, "", &style, 6, 30, false)
	gui.SetStyle(gui.LISTVIEW, gui.SCROLLBAR_WIDTH, int64(style))

	style = int32(gui.GetStyle(gui.DEFAULT, gui.BORDER_WIDTH))
	gui.Label(rl.Rectangle{X: 555, Y: 60, Width: 110, Height: 10}, "BORDER_WIDTH")
	gui.Spinner(rl.Rectangle{X: 670, Y: 55, Width: 90, Height: 20}, "", &style, 0, 20, false)
	gui.SetStyle(gui.DEFAULT, gui.BORDER_WIDTH, int64(style))
	//----------------------------------------------------------
}

func boolToint32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func int32Tobool(v int32) bool {
	return 0 < v
}
