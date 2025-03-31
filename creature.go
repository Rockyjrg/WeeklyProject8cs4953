package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Size  float32
	Speed float32
	Pos   rl.Vector2
}

func NewCreature(size, speed float32, position rl.Vector2) Creature {
	return Creature{
		Size:  size,
		Speed: speed,
		Pos:   position,
	}
}

func (c Creature) DrawCreature() {
	rl.DrawRectangle(int32(c.Pos.X), int32(c.Pos.Y), int32(c.Size), int32(c.Size), rl.Blue)
}

func (c *Creature) MoveCreature(xpos, ypos float32) {
	c.Pos.X += xpos * c.Speed * rl.GetFrameTime()
	c.Pos.Y += ypos * c.Speed * rl.GetFrameTime()

	if c.Pos.Y < 0 {
		c.Pos.Y = 0
	} else if c.Pos.Y+c.Size > float32(rl.GetScreenHeight()) {
		c.Pos.Y = float32(rl.GetScreenHeight()) - c.Size
	}
}
