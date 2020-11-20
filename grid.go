
package main

import (
    fmt     "fmt"
    pixel   "github.com/faiface/pixel"
)

type Grid struct {
	b [][]*Cell
}

func NewGrid(rows, cols int) *Grid {
    cells := [][]*Cell{}

    for j := 0; j < rows; j++ {
        r := []*Cell{}

        for i := 0; i < cols; cols++ {
            cell := NewCell(pixel.V(float64(i), float64(j)))
            // cells = append(cells[j], newCell)
            r = append(r, cell)
        }
        // cells.append(r)
        cells = append(cells, r)
    }

    fmt.Println(cells)


    return &Grid{
        b:  cells,
    }

    // append(cells, r)

}

func (g *Grid) Board() [][]*Cell {
    return g.b
}
