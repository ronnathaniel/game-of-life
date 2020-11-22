
package main

import (
	pixel "github.com/faiface/pixel"
	imdraw "github.com/faiface/pixel/imdraw"
	pgl "github.com/faiface/pixel/pixelgl"
	cnames "golang.org/x/image/colornames"
)


func Run() {
	config := pgl.WindowConfig{
		Title:  "Samo Game of Life",
		Bounds: pixel.R(0, 0, SCREEN_W, SCREEN_H),
		VSync:  true,
		//AlwaysOnTop: true,
		TransparentFramebuffer: true,
	}

	window, err := pgl.NewWindow(config)

	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	grid := NewGrid(ROWS, COLS)

	for ! window.Closed() {
		window.Clear(cnames.Lightslategrey)

		if ! grid.IsSetup() {
			if window.Pressed(pgl.MouseButtonLeft) {
				grid.Plot(imd, window.MousePosition())
			}
			if window.Pressed(pgl.KeySpace) {
				grid.FinishedSetup()
			}
		} else {
			UpdateGrid(imd, grid)
		}

		imd.Draw(window)
		window.Update()
	}
}
