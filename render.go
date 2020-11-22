
package main

import (
	// "fmt"

	pixel "github.com/faiface/pixel"
	imdraw "github.com/faiface/pixel/imdraw"
	pgl "github.com/faiface/pixel/pixelgl"
	cnames "golang.org/x/image/colornames"
)


func Run() {
	config := pgl.WindowConfig{
		Title:  "SamoGamO",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}

	window, err := pgl.NewWindow(config)

	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	grid := NewGrid(ROWS, COLS)

	for ! window.Closed() {
		window.Clear(cnames.Lightslategrey)

		UpdateGrid(imd, grid)

		imd.Draw(window)
		window.Update()
	}
}
