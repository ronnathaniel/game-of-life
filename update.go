
package main

import (
	"fmt"
	imdraw "github.com/faiface/pixel/imdraw"
)

func UpdateGrid(imd *imdraw.IMDraw, g *Grid) {
	for y, row := range g.Board() {
		for x, col := range row {
			UpdateCell(col, g, x, y)
			DrawCell(imd, col)
		}
	}
}

// TODO: redo logic, make look good. switch is ugly
func UpdateCell(c *Cell, g *Grid, xPos, yPos int) {

	_ = fmt.Sprintln(c)

	neighbors := CellLiveNeighbors(g, xPos, yPos)

	if neighbors  > 0 {
		fmt.Println("cell live neighbors: ", neighbors)
	}



	//if neighbors < 2 || 3 < neighbors {
	//	c.Die()
	//} else {
	//
	//}


	if c.IsAlive() {

		if 1 < neighbors && neighbors < 4 {
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

