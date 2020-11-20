
package main

import (
    "fmt"

	pixel "github.com/faiface/pixel"
	imdraw "github.com/faiface/pixel/imdraw"
	cnames "golang.org/x/image/colornames"
)

func DrawSquare (imd *imdraw.IMDraw, c *Cell) {
	imd.Color = cnames.Green

	// imd.Push(pixel.V(100, 100), pixel.V(200, 200))
    imd.Push(c.Pos(), c.EndPos())

	imd.Rectangle(0)

    fmt.Println("drawing rectangle at ", c.Pos().Add(c.EndPos()))
}

func DrawGridInitial(imd *imdraw.IMDraw, g *Grid) {
    imd.Color = pixel.RGB(1, 0, 0)
    imd.Push(pixel.V(100, 100))
    imd.Color = pixel.RGB(1, 0, 0)
    imd.Push(pixel.V(250, 250))

    imd.Polygon(3)
}

func DrawGrid(imd *imdraw.IMDraw, g *Grid) {
    fmt.Println("starting to draw grid")
    // b := g.Board()

    // for _, row := range(b) {
    //     for _, cell := range(row) {
    //         DrawSquare(imd, cell)
    //     }
    // }
}

func draw_pixel_example(imd *imdraw.IMDraw) {
	imd.Color = cnames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.EndShape = imdraw.SharpEndShape
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(30)
}
