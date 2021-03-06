
package main

import (
    pixel "github.com/faiface/pixel"
)

type Cell struct {
	alive       bool
    position    pixel.Vec
    dimensions  pixel.Vec
}

func NewCell(v pixel.Vec) *Cell {
	return &Cell{
		alive:      false,
        position:   v,
        dimensions: pixel.V(CELL_W, CELL_H),
	}
}

func (c *Cell) IsAlive() bool {
	return c.alive
}

func (c *Cell) Populate() {
    c.alive = true
}

func (c *Cell) Die() {
    c.alive = false
}

func (c *Cell) Pos() pixel.Vec {
    return c.position
}

func (c *Cell) Dimen() pixel.Vec {
    return c.dimensions
}

func (c *Cell) EndPos() pixel.Vec {
    return c.Pos().Add(c.Dimen())
}
