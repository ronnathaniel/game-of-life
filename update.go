
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

func UpdateCell(c *Cell, g *Grid, xPos, yPos int) {

	fmt.Println(c)

	neighbors := CellLiveNeighbors(g, xPos, yPos)
	fmt.Println("cell live neighbors: ", neighbors)

	//if neighbors < 2 || 3 < neighbors {
	//	c.Die()
	//}


}

