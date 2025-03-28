package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Name  string
	Level int
	Speed float32
	Pos   rl.Vector2
}

func InitCreature(name string, level int, speed float32, position rl.Vector2) Creature {
	return Creature{
		Name:  name,
		Level: level,
		Speed: speed,
		Pos:   position,
	}
}

func (c Creature) DrawCreature() {
	rl.DrawCircle(int32(c.Pos.X), int32(c.Pos.Y), 30, rl.Blue)
}

func (c *Creature) MoveCreature(xpos, ypos float32) {
	c.Pos.X += xpos * c.Speed * rl.GetFrameTime()
	c.Pos.Y += ypos * c.Speed * rl.GetFrameTime()
}
