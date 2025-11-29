// Package p5go provides a bridge between Go and p5.js.
//
// This package allows you to create interactive graphics and animations using p5.js from Go code
// that compiles to WebAssembly. The package is organized into the following components:
//
// - Constants and type definitions (constants.go)
// - Canvas management (canvas.go)
// - Lifecycle functions like Setup and Draw (lifecycle.go)
// - Event handlers for mouse and keyboard (events.go)
// - Drawing primitives and shapes (drawing.go)
// - Color manipulation (color.go)
// - Transformations (transform.go)
// - Mathematical functions (math.go)
// - Text rendering (text.go)
// - Pixel and image operations (pixels.go)
// - Input handling (input.go)
// - Environment and frame control (environment.go)
// - Helper types and structures (types.go)
//
// Example usage:
//
//	err := p5go.Run("#app",
//		p5go.Setup(func(c *p5go.Canvas) {
//			c.CreateCanvas(400, 400)
//		}),
//		p5go.Draw(func(c *p5go.Canvas) {
//			c.Background(220)
//			c.Ellipse(200, 200, 50, 50)
//		}),
//	)
package p5go
