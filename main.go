package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// collision detection
func CheckCollision(aX, aY, aW, aH, bX, bY, bW, bH float32) bool {
	return aX < bX+bW &&
		aX+aW > bX &&
		aY < bY+bH &&
		aY+aH > bY
}

func main() {
	rl.InitWindow(800, 400, "Practice for midterm and Game Jam.")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := NewCreature(50, 10, 300, rl.NewVector2(30, float32(rl.GetScreenHeight())-50), false)

	gameOver := false
	score := float32(0.0)

	var obstacles []Obstacle
	var spawnTimer float32 = 0
	spawnInterval := float32(1.5)

	//increase game over time
	gameSpeedMultiplier := float32(1.0)

	platform := NewPlatform(rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())/2-10), rl.NewVector2(float32(rand.IntN(200))+100, 20), 200)

	for !rl.WindowShouldClose() {

		// --- input handling ---
		if !gameOver && rl.IsKeyPressed(rl.KeySpace) {
			player.FlipCreature(platform)
		}

		// --- reset game ---
		if gameOver && rl.IsKeyPressed(rl.KeyR) {
			player = NewCreature(50, 10, 300, rl.NewVector2(30, float32(rl.GetScreenHeight())-50), false)
			platform = NewPlatform(rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())/2-10), rl.NewVector2(float32(rand.IntN(200))+100, 20), 200)
			obstacles = nil
			gameOver = false
			score = 0.0
			gameSpeedMultiplier = 1.0
			spawnTimer = 0
		}

		if !gameOver {
			//update score
			score += rl.GetFrameTime()

			//increase game speed slightly over time
			gameSpeedMultiplier += 0.02 * rl.GetFrameTime()

			spawnInterval = float32(math.Max(0.5, 1.5-(float64(gameSpeedMultiplier)-1.0)*0.2))

			//update player
			player.UpdateVerticalMovement(&platform)
			platform.Update(gameSpeedMultiplier)
			if platform.Pos.X+platform.Size.X < 0 {
				platform.Pos.X = float32(rl.GetScreenWidth())
				platform.Size.X = float32(rand.IntN(200)) + 100 //new random width
				platform.Pos.Y = float32(rl.GetScreenHeight())/2 - 10 + float32(rand.IntN(40)-20)
			}

			//spawn obstacles
			spawnTimer += rl.GetFrameTime()
			if spawnTimer >= spawnInterval {
				obstacles = append(obstacles, NewObstacle(rl.GetScreenWidth(), rl.GetScreenHeight(), platform.Pos.Y, gameSpeedMultiplier))
				spawnTimer = 0
			}

			//remove off-screen obstacles
			newObstacles := obstacles[:0] //empty slice

			for i := range obstacles {
				obstacles[i].Update(gameSpeedMultiplier)

				if obstacles[i].Pos.X+obstacles[i].Size.X > 0 {
					//keep if still on screen
					newObstacles = append(newObstacles, obstacles[i])
					if CheckCollision(
						player.Pos.X, player.Pos.Y, player.Size, player.Size,
						obstacles[i].Pos.X, obstacles[i].Pos.Y, obstacles[i].Size.X, obstacles[i].Size.Y,
					) {
						gameOver = true
						player.IsMovingY = false
						break
					}
				}

			}
			obstacles = newObstacles
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if gameOver {
			finalScoreText := fmt.Sprintf("Final Time: %.2f seconds", score)
			rl.DrawText("Game over hehe! Click R to try again", 180, 150, 30, rl.Red)
			rl.DrawText(finalScoreText, (int32(rl.GetScreenWidth())-rl.MeasureText(finalScoreText, 20))/2, int32(rl.GetScreenHeight())/2+10, 20, rl.DarkGray)
		} else {
			player.DrawCreature()
			for i := range obstacles {
				obstacles[i].Draw()
			}
			platform.Draw()

			//draw text
			scoreText := fmt.Sprintf("Time: %.2f", score)
			rl.DrawText(scoreText, 10, 10, 20, rl.DarkGray)
		}

		rl.EndDrawing()
	}
}
