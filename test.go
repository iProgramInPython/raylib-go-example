package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {
	r.InitWindow(800, 450, "Raylib Go Plus")
	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		r.DrawText("Woo! Raylib-Go-Plus! Now with ++", 20, 20, 20, r.GopherBlue)
		r.EndDrawing()
	}
	r.CloseWindow()
}

// This documentation is taken from raylib-goplpus
// I'd like to give credits and thanks to my friend Lachee :)
