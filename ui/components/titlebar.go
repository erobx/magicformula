package components

import rl "github.com/gen2brain/raylib-go/raylib"

type TitleBar struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	col    rl.Color
}

func NewTitleBar(width, height int32, col rl.Color) *TitleBar {
	return &TitleBar{
		posX:   0,
		posY:   0,
		width:  width,
		height: height,
		col:    col,
	}
}

func (tb *TitleBar) DrawTitleBar() {
	rl.DrawRectangle(tb.posX, tb.posY, tb.width, tb.height, tb.col)
}
