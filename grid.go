
package main

import (
	pixel "github.com/faiface/pixel"
	imdraw "github.com/faiface/pixel/imdraw"
)

type Grid struct {
	b 			[][]*Cell
	settingUp	bool
}

// grid should be thought of as any 2D matrix
// j goes down, i goes right
// 0,0 is topleft
func NewGrid(rows, cols int) *Grid {
	b := [][]*Cell{}
    t := []*Cell{}

    for j := 0; j < rows; j++ {
		t = []*Cell{}
		for i := 0; i < cols; i++ {
			c := NewCell(pixel.V(float64(i * CELL_W), float64(j * CELL_H)))
			t = append(t, c)
		}
		b = append(b, t)
	}

    return &Grid{
    	b: b,
    	settingUp: true,
	}

}

func (g *Grid) Board() [][]*Cell {
    return g.b
}

func (g *Grid) IsSetup() bool {
	return ! g.settingUp
}

func (g *Grid) FinishedSetup() {
	g.settingUp = false
}

func (g *Grid) RestartSetup() {
	g = NewGrid(ROWS, COLS)
	g.settingUp = true
}

func (g *Grid) Plot(imd *imdraw.IMDraw, p pixel.Vec) {
	xPlot, yPlot := p.XY()

	i := int(xPlot) / CELL_W
	j := int(yPlot) / CELL_H

	cell := g.Board()[j][i]
	cell.Populate()
	DrawCell(imd, cell)
}


func IsOutOfBounds(x, y int) bool {
	if x < 0 || y < 0 || x >= COLS || y >= ROWS {
		return true
	}

	return false
}

func (g *Grid) CellLiveNeighbors(xPos, yPos int) (count int) {
	count = 0
	var neighborXPos, neighborYPos int

	for horiz := -1; horiz < 2; horiz++ {
		for vert := -1; vert < 2; vert++ {

			neighborXPos = xPos + horiz
			neighborYPos = yPos + vert

			if IsOutOfBounds(neighborXPos, neighborYPos) || (horiz == 0 && vert == 0) {
				continue
			}

			if g.Board()[neighborYPos][neighborXPos].IsAlive() {
				count++
			}
		}
	}
	return
}