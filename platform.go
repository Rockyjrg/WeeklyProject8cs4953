package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Platform struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Speed float32
}

func NewPlatform(pos, size rl.Vector2, speed float32) Platform {
	return Platform{
		Pos:   pos,
		Size:  size,
		Speed: speed,
	}
}

func (p *Platform) Update(speedMultiplier float32) {
	p.Pos.X -= p.Speed * speedMultiplier * rl.GetFrameTime()
}

func (p *Platform) Draw() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Size.X), int32(p.Size.Y), rl.DarkGray)
}
