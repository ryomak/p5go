package p5go

import "syscall/js"

// ColorMode sets the color mode for the canvas.
func (c *Canvas) ColorMode(mode ColorMode, max ...float64) {
	if len(max) > 0 {
		c.p5Instance.Call("colorMode", string(mode), max[0])
	} else {
		c.p5Instance.Call("colorMode", string(mode))
	}
}

// Color creates a color object.
func (c *Canvas) Color(args ...any) js.Value {
	return c.p5Instance.Call("color", args...)
}

// Alpha returns the alpha value of a color.
func (c *Canvas) Alpha(color js.Value) float64 {
	return c.p5Instance.Call("alpha", color).Float()
}

// Red returns the red value of a color.
func (c *Canvas) Red(color js.Value) float64 {
	return c.p5Instance.Call("red", color).Float()
}

// Green returns the green value of a color.
func (c *Canvas) Green(color js.Value) float64 {
	return c.p5Instance.Call("green", color).Float()
}

// Blue returns the blue value of a color.
func (c *Canvas) Blue(color js.Value) float64 {
	return c.p5Instance.Call("blue", color).Float()
}

// Brightness returns the brightness value of a color.
func (c *Canvas) Brightness(color js.Value) float64 {
	return c.p5Instance.Call("brightness", color).Float()
}

// Hue returns the hue value of a color.
func (c *Canvas) Hue(color js.Value) float64 {
	return c.p5Instance.Call("hue", color).Float()
}

// Saturation returns the saturation value of a color.
func (c *Canvas) Saturation(color js.Value) float64 {
	return c.p5Instance.Call("saturation", color).Float()
}

// LerpColor interpolates between two colors.
func (c *Canvas) LerpColor(c1 js.Value, c2 js.Value, amt float64) js.Value {
	return c.p5Instance.Call("lerpColor", c1, c2, amt)
}

// BlendMode sets the blending mode for the canvas.
func (c *Canvas) BlendMode(mode BlendMode) {
	c.p5Instance.Call("blendMode", string(mode))
}
