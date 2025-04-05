package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Size      float32
	Speed     float32
	FlipSpeed float32
	Pos       rl.Vector2
	TargetY   float32
	Flipped   bool
	IsMovingY bool
}

func NewCreature(size, speed, flipSpeed float32, position rl.Vector2, flipped bool) Creature {
	return Creature{
		Size:      size,
		Speed:     speed,
		FlipSpeed: flipSpeed,
		Pos:       position,
		TargetY:   position.Y,
		Flipped:   flipped,
		IsMovingY: false,
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

func (c *Creature) FlipCreature(platform Platform) {

	var newY float32
	if c.Flipped {
		newY = float32(rl.GetScreenHeight() - int(c.Size))
	} else {
		newY = 0
	}

	if CheckCollision(
		c.Pos.X, newY, c.Size, c.Size,
		platform.Pos.X, platform.Pos.Y, platform.Size.X, platform.Size.Y,
	) {
		return
	}

	c.Flipped = !c.Flipped
	c.TargetY = newY
	c.IsMovingY = true
}

// handle smooth floating of the player
func (c *Creature) UpdateVerticalMovement(platform *Platform) {
	if !c.IsMovingY {
		return //not floating
	}

	//calculate distance to move the frame
	moveStep := c.FlipSpeed * rl.GetFrameTime()
	//calculate difference between current and target Y
	diffY := c.TargetY - c.Pos.Y

	moveDirection := float32(math.Copysign(1, float64(diffY)))
	potentialNewY := c.Pos.Y + moveStep*moveDirection

	platformOnScreen := platform.Pos.X+platform.Size.X > 0 //check if platform on screen

	if platformOnScreen {
		//check for collision with the platform
		playerNextYRect := rl.Rectangle{X: c.Pos.X, Y: potentialNewY, Width: c.Size, Height: c.Size}
		platformRect := rl.Rectangle{X: platform.Pos.X, Y: platform.Pos.Y, Width: platform.Size.X, Height: platform.Size.Y}

		collides := rl.CheckCollisionRecs(playerNextYRect, platformRect)

		if collides {
			if moveDirection > 0 {
				c.Pos.Y = platform.Pos.Y - c.Size
				c.IsMovingY = false
			} else if moveDirection < 0 {
				c.Pos.Y = platform.Pos.Y + platform.Size.Y
				c.IsMovingY = false
			} else {
				c.Pos.Y = c.TargetY
				c.IsMovingY = false
			}
			return
		}
	}

	if math.Abs(float64(diffY)) <= float64(moveStep) {
		c.Pos.Y = c.TargetY
		c.IsMovingY = false //stop moving
	} else {
		c.Pos.Y += moveStep * moveDirection
	}
}
