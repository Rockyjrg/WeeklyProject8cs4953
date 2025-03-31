package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "Practice for midterm and Game Jam.")

	defer rl.CloseWindow()
	player := NewCreature(50, 10, rl.NewVector2(30, 30))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		player.DrawCreature()

		//movement
		if rl.IsKeyDown(rl.KeyW) {
			player.MoveCreature(0, -20)
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.MoveCreature(0, 20)
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.MoveCreature(-20, 0)
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.MoveCreature(20, 0)
		}

		rl.EndDrawing()
	}
}
