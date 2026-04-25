package main

import (
	"github.com/ryomak/p5go"
)

// Solar System — sphere primitives + lighting.
// Showcases AmbientLight, PointLight, AmbientMaterial, SpecularMaterial,
// and the new Sphere() primitive.

type planet struct {
	orbit  float64
	radius float64
	speed  float64
	r, g, b float64
}

var planets = []planet{
	{orbit: 80, radius: 8, speed: 1.6, r: 200, g: 180, b: 160},   // Mercury
	{orbit: 120, radius: 14, speed: 1.2, r: 230, g: 200, b: 120}, // Venus
	{orbit: 170, radius: 16, speed: 1.0, r: 100, g: 160, b: 240}, // Earth
	{orbit: 220, radius: 12, speed: 0.8, r: 220, g: 100, b: 80},  // Mars
	{orbit: 290, radius: 28, speed: 0.5, r: 220, g: 180, b: 140}, // Jupiter
	{orbit: 360, radius: 22, speed: 0.35, r: 230, g: 210, b: 150}, // Saturn
}

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
	c.Background(8, 10, 24)
	c.OrbitControl()

	t := float64(c.FrameCount()) * 0.01

	c.AmbientLight(40)
	c.PointLight(255, 240, 200, 0, 0, 0)

	c.Push()
	c.NoStroke()
	c.EmissiveMaterial(255, 220, 120)
	c.Sphere(36)
	c.Pop()

	for i, p := range planets {
		c.Push()
		angle := t*p.speed + float64(i)*1.1
		x := c.Cos(angle) * p.orbit
		z := c.Sin(angle) * p.orbit
		c.Translate(x, 0, z)
		c.NoStroke()
		c.AmbientMaterial(p.r*0.4, p.g*0.4, p.b*0.4)
		c.SpecularMaterial(p.r, p.g, p.b)
		c.Shininess(40)
		c.Sphere(p.radius)
		c.Pop()
	}
}
