package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Obstacle struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Speed float32
	Color rl.Color
}

func NewObstacle(screenWidth, screenHeight int, platformY float32, gameSpeedMultiplier float32) Obstacle {
	//want to get a random size for the obstacles
	width := float32(rand.IntN(30) + 30)
	height := float32(rand.IntN(30) + 30)

	//start the obstacle on the right side, either top or bottom of screen
	yPos := float32(0) //top
	spawnType := rand.IntN(3)

	switch spawnType {
	case 0: // spawn at top
		yPos = 0
	case 1: // spawn at bottom
		yPos = float32(rl.GetScreenHeight()) - height
	case 2: // spawn near middle
		offset := float32(rand.IntN(60) - 30)
		yPos = platformY + offset

		// clamp y position to be within the screen bounds
		if yPos < 0 {
			yPos = 0
		} else if yPos+height > float32(rl.GetScreenHeight()) {
			yPos = float32(rl.GetScreenHeight()) - height
		}
	}

	baseSpeed := float32(rand.IntN(100) + 150)
	currentSpeed := baseSpeed * (1 + (gameSpeedMultiplier-1.0)*0.5)

	return Obstacle{
		Pos:   rl.NewVector2(float32(screenWidth), yPos),
		Size:  rl.NewVector2(width, height),
		Speed: currentSpeed,
		Color: rl.Color{R: uint8(rand.IntN(255)), G: uint8(rand.IntN(255)), B: uint8(rand.IntN(255)), A: 255},
	}
}

func (o *Obstacle) Update(speedMultiplier float32) {
	o.Pos.X -= o.Speed * speedMultiplier * rl.GetFrameTime()
}

func (o *Obstacle) Draw() {
	rl.DrawRectangle(int32(o.Pos.X), int32(o.Pos.Y), int32(o.Size.X), int32(o.Size.Y), o.Color)
}
