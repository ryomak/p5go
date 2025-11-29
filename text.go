package p5go

// Text draws text on the canvas.
func (c *Canvas) Text(str string, x, y float64) {
	c.p5Instance.Call("text", str, x, y)
}

// TextFont sets the font and size for text.
func (c *Canvas) TextFont(font string, size float64) {
	c.p5Instance.Call("textFont", font, size)
}

// TextSize sets the size for text.
func (c *Canvas) TextSize(size float64) {
	c.p5Instance.Call("textSize", size)
}

// TextAlign sets the alignment for text.
func (c *Canvas) TextAlign(align DrawingMode) {
	c.p5Instance.Call("textAlign", string(align))
}

// TextWrap sets the wrap mode for text.
func (c *Canvas) TextWrap(w DrawingMode) {
	c.p5Instance.Call("textWrap", string(w))
}

// TextAscent returns the ascent of the current font.
func (c *Canvas) TextAscent() float64 {
	return c.p5Instance.Call("textAscent").Float()
}

// TextDescent returns the descent of the current font.
func (c *Canvas) TextDescent() float64 {
	return c.p5Instance.Call("textDescent").Float()
}

// TextLeading sets the leading for text.
func (c *Canvas) TextLeading(leading float64) {
	c.p5Instance.Call("textLeading", leading)
}

// TextStyle sets the style for text.
func (c *Canvas) TextStyle(style TextStyle) {
	c.p5Instance.Call("textStyle", string(style))
}

// TextWidth returns the width of the specified text.
func (c *Canvas) TextWidth(text string) float64 {
	return c.p5Instance.Call("textWidth", text).Float()
}
