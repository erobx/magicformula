package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 600
	height = 500
	size   = width * height
)

type App struct {
	button  *Button
	bgColor color.Color
}

func NewApp() *App {
	a := &App{}
	a.button = &Button{
		Rect: image.Rect(16, 16, 144, 48),
		Text: "Button 1",
	}
	a.button.SetOnPressed(func(b *Button) {
		fmt.Println("Button pressed")
	})
	a.bgColor = color.White

	return a
}

func (a *App) Update() error {
	a.button.Update()
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(a.bgColor)
	a.button.Draw(screen)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (a *App) Run(title string) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}
}
