
package main

import (
	//pgl "github.com/faiface/pixel/pixelgl"
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

// TODO: add dynamic initial state before running
//func DrawInitialState (imd *imdraw.IMDraw) {
//
//}
