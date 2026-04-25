package main

import (
	"github.com/ryomak/p5go"
)

// Wireframe Sphere — stroke-only spheres of varying detail.
// Showcases Sphere() and Torus() with NoFill + Stroke.

func main() {
	p5go.Run("main",
		p5go.Setup(func(c *p5go.Canvas) {
			c.CreateCanvas(800, 600, p5go.WEBGL)
		}),
		p5go.Draw(draw),
	)
	select {}
}

func draw(c *p5go.Canvas) {
	c.Background(12, 14, 22)
	c.OrbitControl()

	t := float64(c.FrameCount()) * 0.01

	c.NoFill()
	c.StrokeWeight(0.6)

	c.Push()
	c.RotateY(t)
	c.RotateX(t * 0.7)
	c.Stroke(120, 220, 200)
	c.Sphere(180, 18, 12)
	c.Pop()

	c.Push()
	c.RotateY(-t * 1.3)
	c.RotateZ(t * 0.5)
	c.Stroke(220, 120, 200)
	c.Torus(220, 8, 32, 8)
	c.Pop()

	c.Push()
	c.RotateX(t * 1.5)
	c.Stroke(220, 220, 120, 180)
	c.Sphere(80, 8, 6)
	c.Pop()
}
