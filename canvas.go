package p5go

import (
	"errors"
	"syscall/js"
)

// Canvas represents a p5.js canvas.
type Canvas struct {
	p5Instance   js.Value
	funcHandlers map[string]js.Func
	width        float64
	height       float64
}

// Validate checks if the p5.js instance and required handlers are set.
func (c *Canvas) Validate() error {
	if c.p5Instance.Type() == js.TypeUndefined {
		return errors.New("p5.js is not loaded")
	}
	if c.funcHandlers["setup"].IsUndefined() {
		return errors.New("setup function is not defined")
	}
	if c.funcHandlers["draw"].IsUndefined() {
		return errors.New("draw function is not defined")
	}
	return nil
}

// CreateCanvas creates a new canvas with the specified width and height.
func (c *Canvas) CreateCanvas(w, h int, opts ...RendererMode) {
	c.width = float64(w)
	c.height = float64(h)
	if len(opts) > 0 {
		c.p5Instance.Call("createCanvas", w, h, string(opts[0]))
	} else {
		c.p5Instance.Call("createCanvas", w, h)
	}
}

// Width returns the width of the canvas.
func (c *Canvas) Width() float64 {
	return c.width
}

// Height returns the height of the canvas.
func (c *Canvas) Height() float64 {
	return c.height
}

// Size sets the size of the canvas.
func (c *Canvas) Size(width, height float64) {
	c.width = width
	c.height = height
	c.p5Instance.Call("size", width, height)
}

// Hide hides the canvas.
func (c *Canvas) Hide() {
	c.p5Instance.Call("hide")
}

// Clear clears the canvas.
func (c *Canvas) Clear() {
	c.p5Instance.Call("clear")
}

// CreateGraphics creates and returns a new graphics buffer.
func (c *Canvas) CreateGraphics(w, h float64, renderer ...any) *Canvas {
	return &Canvas{p5Instance: c.p5Instance.Call("createGraphics", append([]any{w, h}, renderer...)...)}
}

// CreateCapture creates a capture object.
func (c *Canvas) CreateCapture(kind CaptureKind) js.Value {
	return c.p5Instance.Call("createCapture", string(kind))
}
