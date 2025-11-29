package p5go

// Cursor sets the cursor style.
func (c *Canvas) Cursor(style CursorStyle) {
	c.p5Instance.Call("cursor", string(style))
}

// NoCursor hides the cursor.
func (c *Canvas) NoCursor() {
	c.p5Instance.Call("noCursor")
}

// WindowWidth returns the width of the window.
func (c *Canvas) WindowWidth() float64 {
	return c.p5Instance.Get("windowWidth").Float()
}

// WindowHeight returns the height of the window.
func (c *Canvas) WindowHeight() float64 {
	return c.p5Instance.Get("windowHeight").Float()
}

// FrameRate sets the frame rate for the canvas.
func (c *Canvas) FrameRate(fps float64) {
	c.p5Instance.Call("frameRate", fps)
}

// FrameCount returns the number of frames that have been displayed.
func (c *Canvas) FrameCount() int {
	return c.p5Instance.Get("frameCount").Int()
}

// GetFrameRate returns the current frame rate.
func (c *Canvas) GetFrameRate() float64 {
	return c.p5Instance.Get("frameRate").Float()
}

// Loop starts the draw loop.
func (c *Canvas) Loop() {
	c.p5Instance.Call("loop")
}

// NoLoop stops the draw loop.
func (c *Canvas) NoLoop() {
	c.p5Instance.Call("noLoop")
}

// IsLooping returns true if the draw loop is currently running.
func (c *Canvas) IsLooping() bool {
	return c.p5Instance.Call("isLooping").Bool()
}

// Redraw redraws the canvas.
func (c *Canvas) Redraw() {
	c.p5Instance.Call("redraw")
}

// Save saves the canvas as an image file.
func (c *Canvas) Save(filename string) {
	c.p5Instance.Call("save", filename)
}

// SaveFrames saves a sequence of frames as image files.
func (c *Canvas) SaveFrames(filename string, extension string, duration float64, fps float64) {
	c.p5Instance.Call("saveFrames", filename, extension, duration, fps)
}

// SaveCanvas saves the canvas as an image file.
func (c *Canvas) SaveCanvas(filename, extension string) {
	c.p5Instance.Call("saveCanvas", filename, extension)
}

// SaveGif saves the canvas as a GIF file.
func (c *Canvas) SaveGif(name string, second float64) {
	c.p5Instance.Call("saveGif", name, second)
}
