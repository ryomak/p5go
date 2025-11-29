# p5go

[![Go Reference](https://pkg.go.dev/badge/github.com/ryomak/p5go.svg)](https://pkg.go.dev/github.com/ryomak/p5go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ryomak/p5go)](https://go.dev/)

Go wrapper for [p5.js](https://p5js.org/) - Create interactive graphics and animations in Go that run in the browser via WebAssembly.

Inspired by [p5rb](https://github.com/ongaeshi/p5rb).

## Installation

```bash
go get github.com/ryomak/p5go
```

**Requirements:** Go 1.22+

## Quick Start

```go
package main

import "github.com/ryomak/p5go"

func main() {
	p5go.Run("#container",
		p5go.Setup(func(c *p5go.Canvas) {
			c.CreateCanvas(400, 400)
		}),
		p5go.Draw(func(c *p5go.Canvas) {
			c.Background(220)
			c.Fill(255, 0, 0)
			c.Circle(200, 200, 50)
		}),
	)
	select {}
}
```

Build and run:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm
# Serve with any HTTP server
```

## Examples

See the [examples directory](example/) for more:

- [Simple](example/simple/) - Basic shapes and drawing
- [Playground](example/playground/) - Interactive animation with mouse events

## Features

- Drawing: shapes, text, images, custom paths
- Styling: colors, strokes, fills, transformations
- Events: mouse and keyboard input
- Export: save as image or animated GIF
- Graphics buffers for layered compositions

## API

For the complete API reference, see [pkg.go.dev](https://pkg.go.dev/github.com/ryomak/p5go).

Common methods:

```go
// Canvas
c.CreateCanvas(width, height)
c.Background(r, g, b)

// Shapes
c.Circle(x, y, diameter)
c.Rect(x, y, w, h)
c.Line(x1, y1, x2, y2)

// Styling
c.Fill(r, g, b)
c.Stroke(r, g, b)
c.StrokeWeight(weight)

// Transform
c.Translate(x, y)
c.Rotate(angle)
c.Push() / c.Pop()

// Events
p5go.MousePressed(handler)
p5go.KeyPressed(handler)

// Export
c.SaveCanvas(filename, ext)
c.SaveGif(filename, seconds)
```

## Contributing

Pull requests are welcome! Many p5.js features are still being added.

## Resources

- [p5.js Reference](https://p5js.org/reference/)
- [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
