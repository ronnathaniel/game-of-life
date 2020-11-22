
package main

import (
	imdraw "github.com/faiface/pixel/imdraw"
	cnames "golang.org/x/image/colornames"
)



func DrawCell (imd *imdraw.IMDraw, c *Cell) {

	if c.IsAlive() {
		imd.Color = cnames.Darkolivegreen
	} else {
		imd.Color = cnames.Lightslategrey
	}
	imd.Push(c.Pos(), c.EndPos())
	imd.Rectangle(0)
}

func DrawGrid(imd *imdraw.IMDraw, grid *Grid) {

	for _, row := range grid.Board() {
		for _, col := range row {
			DrawCell(imd, col)
		}
	}
}

