
package main

import (
	"fmt"
	pixel "github.com/faiface/pixel"
)

type Grid struct {
	b [][]*Cell
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

			// TODO: CHANGE THIS - 100 means all first 100 rows are populated initially
			if j < 100 {
				c.Populate()
			}
			t = append(t, c)
		}
		b = append(b, t)
	}

    return &Grid{
    	b: b,
	}

}

func (g *Grid) Board() [][]*Cell {
    return g.b
}

func IsOutOfBounds(x, y int) bool {
	if x < 0 || y < 0 || x >= COLS || y >= ROWS {
		return false
	}

	return true
}

func CellLiveNeighbors(g *Grid, xPos, yPos int) (count int) {
	count = 0

	for horiz := -1; horiz < 2; horiz++ {
		for vert := -1; vert < 2; vert++ {

			if (horiz == 0 && vert == 0) || IsOutOfBounds(xPos, yPos) {
				continue
			}

			fmt.Println("vert: ", vert, "horiz: ", horiz)
			if g.Board()[vert][horiz].IsAlive() {
				count++
				fmt.Println("count, ", count)
			}
		}
	}
	return
}