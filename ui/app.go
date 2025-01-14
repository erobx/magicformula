package ui

import (
	"github.com/erobx/magicformula/ui/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 500
	height = 500
)

type App struct {
}

func (a *App) Run(title string) {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	tb := components.NewTitleBar(width, height, rl.RayWhite)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		tb.DrawTitleBar()

		rl.EndDrawing()
	}
}
