package p5go

// MouseX returns the current x-coordinate of the mouse.
func (c *Canvas) MouseX() float64 {
	return c.p5Instance.Get("mouseX").Float()
}

// PMouseX returns the previous x-coordinate of the mouse.
func (c *Canvas) PMouseX() float64 {
	return c.p5Instance.Get("pmouseX").Float()
}

// MouseY returns the current y-coordinate of the mouse.
func (c *Canvas) MouseY() float64 {
	return c.p5Instance.Get("mouseY").Float()
}

// PMouseY returns the previous y-coordinate of the mouse.
func (c *Canvas) PMouseY() float64 {
	return c.p5Instance.Get("pmouseY").Float()
}

// MouseIsPressed returns true if the mouse is currently pressed.
func (c *Canvas) MouseIsPressed() bool {
	return c.p5Instance.Get("mouseIsPressed").Bool()
}

// MovedX returns the amount the mouse has moved along the x-axis.
func (c *Canvas) MovedX() float64 {
	return c.p5Instance.Get("movedX").Float()
}

// MovedY returns the amount the mouse has moved along the y-axis.
func (c *Canvas) MovedY() float64 {
	return c.p5Instance.Get("movedY").Float()
}

// MouseButton returns the current mouse button being pressed.
func (c *Canvas) MouseButton() string {
	return c.p5Instance.Get("mouseButton").String()
}

// Key returns the current key being pressed.
func (c *Canvas) Key() string {
	return c.p5Instance.Get("key").String()
}

// KeyCode returns the key code of the current key being pressed.
func (c *Canvas) KeyCode() int {
	return c.p5Instance.Get("keyCode").Int()
}

// KeyIsPressed returns true if a key is currently pressed.
func (c *Canvas) KeyIsPressed() bool {
	return c.p5Instance.Get("keyIsPressed").Bool()
}
