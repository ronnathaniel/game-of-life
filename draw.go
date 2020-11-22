
package main

import (
	"fmt"
	pixel "github.com/faiface/pixel"
	imdraw "github.com/faiface/pixel/imdraw"
	cnames "golang.org/x/image/colornames"
)



func DrawCell (imd *imdraw.IMDraw, c *Cell) {
	imd.Color = cnames.Darkolivegreen
	imd.Push(c.Pos(), c.EndPos())
	imd.Rectangle(0)
}

func DrawGrid(imd *imdraw.IMDraw, grid *Grid) {

	fmt.Print(grid)

	DrawCell(imd, NewCell(pixel.V(200, 200)))
}

