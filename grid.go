
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
			c := NewCell(pixel.V(float64(i), float64(j)))
			fmt.Println("adding cell", c)
			t = append(t, c)
		}
		b = append(b, t)
	}

	fmt.Println("finished creating grid")

    return &Grid{
    	b: b,
	}

}

func (g *Grid) Board() [][]*Cell {
    return g.b
}
