
package main

import (
	imdraw "github.com/faiface/pixel/imdraw"
)

func UpdateGrid(imd *imdraw.IMDraw, g *Grid) {
	for y, row := range g.Board() {
		for x, cell := range row {
			UpdateCell(cell, g, x, y)
			DrawCell(imd, cell)
		}
	}
}

func UpdateCell(c *Cell, g *Grid, xPos, yPos int) {

	neighbors := g.CellLiveNeighbors(xPos, yPos)

	if c.IsAlive() {

		if 2 <= neighbors && neighbors <= 3 {
			c.Populate()
		} else {
			c.Die()
		}
	} else {
		if neighbors == 3 {
			c.Populate()
		}
	}
}

