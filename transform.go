package p5go

// Push saves the current drawing style settings and transformations.
func (c *Canvas) Push() {
	c.p5Instance.Call("push")
}

// Pop restores the previous drawing style settings and transformations.
func (c *Canvas) Pop() {
	c.p5Instance.Call("pop")
}

// Translate translates the canvas by the specified x and y values.
func (c *Canvas) Translate(x, y float64, z ...float64) {
	if len(z) > 0 {
		c.p5Instance.Call("translate", x, y, z[0])
	} else {
		c.p5Instance.Call("translate", x, y)
	}
}

// Rotate rotates the canvas by the specified angle.
func (c *Canvas) Rotate(angle float64) {
	c.p5Instance.Call("rotate", angle)
}

// RotateX rotates the canvas around the x-axis by the specified angle.
func (c *Canvas) RotateX(angle float64) {
	c.p5Instance.Call("rotateX", angle)
}

// RotateY rotates the canvas around the y-axis by the specified angle.
func (c *Canvas) RotateY(angle float64) {
	c.p5Instance.Call("rotateY", angle)
}

// RotateZ rotates the canvas around the z-axis by the specified angle.
func (c *Canvas) RotateZ(angle float64) {
	c.p5Instance.Call("rotateZ", angle)
}

// Scale scales the canvas by the specified factor.
func (c *Canvas) Scale(s float64) {
	c.p5Instance.Call("scale", s)
}

// ShearX shears the canvas along the x-axis by the specified angle.
func (c *Canvas) ShearX(angle float64) {
	c.p5Instance.Call("shearX", angle)
}

// ShearY shears the canvas along the y-axis by the specified angle.
func (c *Canvas) ShearY(angle float64) {
	c.p5Instance.Call("shearY", angle)
}

// ApplyMatrix applies a transformation matrix to the canvas.
func (c *Canvas) ApplyMatrix(a, b, c1, d, e, f float64) {
	c.p5Instance.Call("applyMatrix", a, b, c1, d, e, f)
}

// ResetMatrix resets the transformation matrix.
func (c *Canvas) ResetMatrix() {
	c.p5Instance.Call("resetMatrix")
}

// OrbitControl represents a control for orbiting around an object
func (c *Canvas) OrbitControl(opts ...any) {
	c.p5Instance.Call("orbitControl", opts...)
}
