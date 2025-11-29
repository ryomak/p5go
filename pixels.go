package p5go

import "syscall/js"

// LoadImage loads an image from the specified path.
func (c *Canvas) LoadImage(path string) js.Value {
	return c.p5Instance.Call("loadImage", path)
}

// Image draws an image on the canvas.
func (c *Canvas) Image(img any, opts ...any) {
	c.p5Instance.Call("image", append([]any{img}, opts...)...)
}

// LoadPixels loads the pixel data for the canvas into the pixels[] array.
func (c *Canvas) LoadPixels() {
	c.p5Instance.Call("loadPixels")
}

// UpdatePixels updates the canvas with the data in the pixels[] array.
func (c *Canvas) UpdatePixels() {
	c.p5Instance.Call("updatePixels")
}

// Get retrieves the color of any pixel or grabs a section of an image.
func (c *Canvas) Get(x, y float64) js.Value {
	return c.p5Instance.Call("get", x, y)
}

// Set changes the color of any pixel or writes an image into the canvas.
func (c *Canvas) Set(x, y float64, color js.Value) {
	c.p5Instance.Call("set", x, y, color)
}

// Copy copies a region of pixels from one image to another.
func (c *Canvas) Copy(srcImage js.Value, sx, sy, sw, sh, dx, dy, dw, dh float64) {
	c.p5Instance.Call("copy", srcImage, sx, sy, sw, sh, dx, dy, dw, dh)
}

// Filter applies a filter to the canvas.
func (c *Canvas) Filter(filterType FilterType, value ...float64) {
	if len(value) > 0 {
		c.p5Instance.Call("filter", string(filterType), value[0])
	} else {
		c.p5Instance.Call("filter", string(filterType))
	}
}

// Blend blends a region of pixels using a specified blend mode.
func (c *Canvas) Blend(sx, sy, sw, sh, dx, dy, dw, dh float64, blendMode BlendMode) {
	c.p5Instance.Call("blend", sx, sy, sw, sh, dx, dy, dw, dh, string(blendMode))
}

// Mask applies an image as a mask to the canvas.
func (c *Canvas) Mask(img js.Value) {
	c.p5Instance.Call("mask", img)
}
