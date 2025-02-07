// Package p5go provides a bridge between Go and p5.js.
package p5go

import (
	"errors"
	"fmt"
	"math"
	"syscall/js"
)

// Constants
const (
	// Renderer modes
	P2D   = "p2d"
	WEBGL = "webgl"

	// Environment
	ARROW = "default"
	CROSS = "crosshair"
	HAND  = "pointer"
	MOVE  = "move"
	TEXT  = "text"
	WAIT  = "wait"

	// Trigonometry
	PI         = math.Pi
	HALF_PI    = math.Pi / 2
	QUARTER_PI = math.Pi / 4
	TWO_PI     = math.Pi * 2
	TAU        = TWO_PI
	DEGREES    = "degrees"
	RADIANS    = "radians"

	// Color modes
	RGB = "rgb"
	HSB = "hsb"
	HSL = "hsl"

	// Drawing modes
	CORNER   = "corner"
	CORNERS  = "corners"
	RADIUS   = "radius"
	RIGHT    = "right"
	LEFT     = "left"
	CENTER   = "center"
	TOP      = "top"
	BOTTOM   = "bottom"
	BASELINE = "alphabetic"

	// Shape modes
	POINTS         = 0x0000
	LINES          = 0x0001
	LINE_STRIP     = 0x0003
	LINE_LOOP      = 0x0002
	TRIANGLES      = 0x0004
	TRIANGLE_FAN   = 0x0006
	TRIANGLE_STRIP = 0x0005
	QUADS          = "quads"
	QUAD_STRIP     = "quad_strip"
	TESS           = "tess"
	CLOSE          = "close"
	OPEN           = "open"
	CHORD          = "chord"
	PIE            = "pie"
	PROJECT        = "square"
	SQUARE         = "butt"
	ROUND          = "round"
	BEVEL          = "bevel"
	MITER          = "miter"

	// Blend modes
	BLEND      = "source-over"
	REMOVE     = "destination-out"
	ADD        = "lighter"
	DARKEST    = "darken"
	LIGHTEST   = "lighten"
	DIFFERENCE = "difference"
	SUBTRACT   = "subtract"
	EXCLUSION  = "exclusion"
	MULTIPLY   = "multiply"
	SCREEN     = "screen"
	REPLACE    = "copy"
	OVERLAY    = "overlay"
	HARD_LIGHT = "hard-light"
	SOFT_LIGHT = "soft-light"
	DODGE      = "color-dodge"
	BURN       = "color-burn"

	// Image filters
	THRESHOLD = "threshold"
	GRAY      = "gray"
	OPAQUE    = "opaque"
	INVERT    = "invert"
	POSTERIZE = "posterize"
	DILATE    = "dilate"
	ERODE     = "erode"
	BLUR      = "blur"

	// Typography
	NORMAL     = "normal"
	ITALIC     = "italic"
	BOLD       = "bold"
	BOLDITALIC = "bold italic"

	// Web GL specific
	IMMEDIATE = "immediate"
	IMAGE     = "image"
	NEAREST   = "nearest"
	REPEAT    = "repeat"
	CLAMP     = "clamp"
	MIRROR    = "mirror"

	// Device orientation
	LANDSCAPE = "landscape"
	PORTRAIT  = "portrait"
)

var (
	global = js.Global()
)

// Run initializes the p5 p5Instance
func Run(query string, fs ...Func) error {
	// Get container
	container := global.Get("document").Call("querySelector", query)
	if container.IsNull() {
		return errors.New(fmt.Sprintf("%s is not match", query))
	}
	container.Set("innerHTML", "")

	// P5.jsがロードされていない場合は追加
	if global.Get("p5").IsUndefined() {
		doc := global.Get("document")
		script := doc.Call("createElement", "script")
		script.Set("src", "https://cdn.jsdelivr.net/npm/p5@1.11.2/lib/p5.min.js")
		doc.Get("head").Call("appendChild", script)

		ch := make(chan struct{})
		script.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			close(ch)
			return nil
		}))
		<-ch
	}

	c := &Canvas{
		p5Instance:   js.Undefined(),
		funcHandlers: map[string]js.Func{},
	}

	sketch := js.FuncOf(func(this js.Value, args []js.Value) any {
		c.p5Instance = args[0]
		for _, f := range fs {
			f(c)
		}

		for method, handler := range c.funcHandlers {
			c.p5Instance.Set(method, handler)
		}
		return nil
	})

	p5Constructor := global.Get("p5")
	p5Constructor.New(sketch, container)

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

// Func is a type that represents a function that takes a Canvas pointer as an argument.
type Func func(c *Canvas)

// Preload sets the preload handler for the canvas.
func Preload(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["preload"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// Setup sets the setup handler for the canvas.
func Setup(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["setup"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// Draw sets the draw handler for the canvas.
func Draw(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["draw"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MouseMoved sets the mouseMoved handler for the canvas.
func MouseMoved(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseMoved"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MouseDragged sets the mouseDragged handler for the canvas.
func MouseDragged(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseDragged"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MousePressed sets the mousePressed handler for the canvas.
func MousePressed(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mousePressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MouseReleased sets the mouseReleased handler for the canvas.
func MouseReleased(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MouseClicked sets the mouseClicked handler for the canvas.
func MouseClicked(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// DoubleClicked sets the doubleClicked handler for the canvas.
func DoubleClicked(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["doubleClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// MouseWheel sets the mouseWheel handler for the canvas.
func MouseWheel(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseWheel"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// KeyPressed sets the keyPressed handler for the canvas.
func KeyPressed(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyPressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// KeyReleased sets the keyReleased handler for the canvas.
func KeyReleased(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// KeyTyped sets the keyTyped handler for the canvas.
func KeyTyped(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyTyped"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// Canvas represents a p5.js canvas.
type Canvas struct {
	p5Instance   js.Value
	funcHandlers map[string]js.Func
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
func (c *Canvas) CreateCanvas(w, h int) {
	c.p5Instance.Call("createCanvas", w, h)
}

// Background sets the background color of the canvas.
func (c *Canvas) Background(args ...any) {
	c.p5Instance.Call("background", args...)
}

// Fill sets the fill color for shapes.
func (c *Canvas) Fill(args ...any) {
	c.p5Instance.Call("fill", args...)
}

// Stroke sets the stroke color for shapes.
func (c *Canvas) Stroke(args ...any) {
	c.p5Instance.Call("stroke", args...)
}

// NoFill disables filling shapes.
func (c *Canvas) NoFill() {
	c.p5Instance.Call("noFill")
}

// NoStroke disables drawing the stroke for shapes.
func (c *Canvas) NoStroke() {
	c.p5Instance.Call("noStroke")
}

// Ellipse draws an ellipse on the canvas.
func (c *Canvas) Ellipse(x, y, w, h float64) {
	c.p5Instance.Call("ellipse", x, y, w, h)
}

// Rect draws a rectangle on the canvas.
func (c *Canvas) Rect(x, y, w, h float64) {
	c.p5Instance.Call("rect", x, y, w, h)
}

// Line draws a line on the canvas.
func (c *Canvas) Line(x1, y1, x2, y2 float64) {
	c.p5Instance.Call("line", x1, y1, x2, y2)
}

// Triangle draws a triangle on the canvas.
func (c *Canvas) Triangle(x1, y1, x2, y2, x3, y3 float64) {
	c.p5Instance.Call("triangle", x1, y1, x2, y2, x3, y3)
}

// Point draws a point on the canvas.
func (c *Canvas) Point(x, y float64) {
	c.p5Instance.Call("point", x, y)
}

// Arc draws an arc on the canvas.
func (c *Canvas) Arc(x, y, w, h, start, stop float64) {
	c.p5Instance.Call("arc", x, y, w, h, start, stop)
}

// Bezier draws a bezier curve on the canvas.
func (c *Canvas) Bezier(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	c.p5Instance.Call("bezier", x1, y1, x2, y2, x3, y3, x4, y4)
}

// QuadraticVertex draws a quadratic vertex on the canvas.
func (c *Canvas) QuadraticVertex(cx, cy, x, y float64) {
	c.p5Instance.Call("quadraticVertex", cx, cy, x, y)
}

// Curve draws a curve on the canvas.
func (c *Canvas) Curve(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	c.p5Instance.Call("curve", x1, y1, x2, y2, x3, y3, x4, y4)
}

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

// Push saves the current drawing style settings and transformations.
func (c *Canvas) Push() {
	c.p5Instance.Call("push")
}

// Pop restores the previous drawing style settings and transformations.
func (c *Canvas) Pop() {
	c.p5Instance.Call("pop")
}

// Translate translates the canvas by the specified x and y values.
func (c *Canvas) Translate(x, y float64) {
	c.p5Instance.Call("translate", x, y)
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

// SaveCanvas saves the canvas as an image file.
func (c *Canvas) SaveCanvas(filename, extension string) {
	c.p5Instance.Call("saveCanvas", filename, extension)
}

// LoadImage loads an image from the specified path.
func (c *Canvas) LoadImage(path string) js.Value {
	return c.p5Instance.Call("loadImage", path)
}

// Image draws an image on the canvas.
func (c *Canvas) Image(img js.Value, x, y, w, h float64) {
	c.p5Instance.Call("image", img, x, y, w, h)
}

// FrameRate sets the frame rate for the canvas.
func (c *Canvas) FrameRate(fps float64) {
	c.p5Instance.Call("frameRate", fps)
}

// Random returns a random number between the specified min and max values.
func (c *Canvas) Random(min, max float64) float64 {
	return c.p5Instance.Call("random", min, max).Float()
}

// Map maps a value from one range to another.
func (c *Canvas) Map(value, start1, stop1, start2, stop2 float64) float64 {
	return c.p5Instance.Call("map", value, start1, stop1, start2, stop2).Float()
}

// BeginShape begins recording vertices for a shape.
func (c *Canvas) BeginShape(option ...any) {
	if len(option) > 0 {
		c.p5Instance.Call("beginShape", option[0])
	} else {
		c.p5Instance.Call("beginShape")
	}
}

// Vertex adds a vertex to the current shape.
func (c *Canvas) Vertex(x, y float64) {
	c.p5Instance.Call("vertex", x, y)
}

// EndShape ends recording vertices for a shape.
func (c *Canvas) EndShape(mode ...string) {
	if len(mode) > 0 {
		c.p5Instance.Call("endShape", mode[0])
	} else {
		c.p5Instance.Call("endShape")
	}
}

// BezierVertex adds a bezier vertex to the current shape.
func (c *Canvas) BezierVertex(cx1, cy1, cx2, cy2, x, y float64) {
	c.p5Instance.Call("bezierVertex", cx1, cy1, cx2, cy2, x, y)
}

// CurveVertex adds a curve vertex to the current shape.
func (c *Canvas) CurveVertex(x, y float64) {
	c.p5Instance.Call("curveVertex", x, y)
}

// BeginContour begins recording vertices for a contour.
func (c *Canvas) BeginContour() {
	c.p5Instance.Call("beginContour")
}

// EndContour ends recording vertices for a contour.
func (c *Canvas) EndContour() {
	c.p5Instance.Call("endContour")
}

// Close closes the current shape.
func (c *Canvas) Close() {
	c.p5Instance.Call("close")
}

// TextAlign sets the alignment for text.
func (c *Canvas) TextAlign(align string) {
	c.p5Instance.Call("textAlign", align)
}

// TextWrap sets the wrap mode for text.
func (c *Canvas) TextWrap(w string) {
	c.p5Instance.Call("textWrap", w)
}

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

// SaveGif saves the canvas as a GIF file.
func (c *Canvas) SaveGif(name string, second float64) {
	c.p5Instance.Call("saveGif", name, second)
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

// ColorMode sets the color mode for the canvas.
func (c *Canvas) ColorMode(mode string, max ...float64) {
	if len(max) > 0 {
		c.p5Instance.Call("colorMode", mode, max[0])
	} else {
		c.p5Instance.Call("colorMode", mode)
	}
}

// Acos returns the arccosine of a value.
func (c *Canvas) Acos(value float64) float64 {
	return c.p5Instance.Call("acos", value).Float()
}

// Cos returns the cosine of a value.
func (c *Canvas) Cos(value float64) float64 {
	return c.p5Instance.Call("cos", value).Float()
}

// AngleMode sets the angle mode for the canvas.
func (c *Canvas) AngleMode(mode string) {
	c.p5Instance.Call("angleMode", mode)
}

// Asin returns the arcsine of a value.
func (c *Canvas) Asin(value float64) float64 {
	return c.p5Instance.Call("asin", value).Float()
}

// Atan returns the arctangent of a value.
func (c *Canvas) Atan(value float64) float64 {
	return c.p5Instance.Call("atan", value).Float()
}

// Atan2 returns the arctangent of y/x.
func (c *Canvas) Atan2(y, x float64) float64 {
	return c.p5Instance.Call("atan2", y, x).Float()
}

// Sin returns the sine of a value.
func (c *Canvas) Sin(value float64) float64 {
	return c.p5Instance.Call("sin", value).Float()
}

// Tan returns the tangent of a value.
func (c *Canvas) Tan(value float64) float64 {
	return c.p5Instance.Call("tan", value).Float()
}

// Degrees converts a value from radians to degrees.
func (c *Canvas) Degrees(value float64) float64 {
	return c.p5Instance.Call("degrees", value).Float()
}

// Radians converts a value from degrees to radians.
func (c *Canvas) Radians(value float64) float64 {
	return c.p5Instance.Call("radians", value).Float()
}

// StrokeWeight sets the weight of the stroke.
func (c *Canvas) StrokeWeight(weight float64) {
	c.p5Instance.Call("strokeWeight", weight)
}

// StrokeCap sets the style of the stroke cap.
func (c *Canvas) StrokeCap(cap string) {
	c.p5Instance.Call("strokeCap", cap)
}

// Erase enables the eraser tool.
func (c *Canvas) Erase(opt ...any) {
	c.p5Instance.Call("erase", opt...)
}

// NoErase disables the eraser tool.
func (c *Canvas) NoErase() {
	c.p5Instance.Call("noErase")
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

// Circle draws a circle on the canvas.
func (c *Canvas) Circle(x, y, d float64) {
	c.p5Instance.Call("circle", x, y, d)
}

// Square draws a square on the canvas.
func (c *Canvas) Square(x, y, s float64) {
	c.p5Instance.Call("square", x, y, s)
}

// Color creates a color object.
func (c *Canvas) Color(args ...any) js.Value {
	return c.p5Instance.Call("color", args...)
}

// Clear clears the canvas.
func (c *Canvas) Clear() {
	c.p5Instance.Call("clear")
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
func (c *Canvas) TextStyle(style string) {
	c.p5Instance.Call("textStyle", style)
}

// TextWidth returns the width of the specified text.
func (c *Canvas) TextWidth(text string) float64 {
	return c.p5Instance.Call("textWidth", text).Float()
}

// Cursor sets the cursor style.
func (c *Canvas) Cursor(style string) {
	c.p5Instance.Call("cursor", style)
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

// Width returns the width of the canvas.
func (c *Canvas) Width() float64 {
	return c.p5Instance.Get("width").Float()
}

// Height returns the height of the canvas.
func (c *Canvas) Height() float64 {
	return c.p5Instance.Get("height").Float()
}

// ApplyMatrix applies a transformation matrix to the canvas.
func (c *Canvas) ApplyMatrix(a, b, c1, d, e, f float64) {
	c.p5Instance.Call("applyMatrix", a, b, c1, d, e, f)
}

// ResetMatrix resets the transformation matrix.
func (c *Canvas) ResetMatrix() {
	c.p5Instance.Call("resetMatrix")
}

// Abs returns the absolute value of the given number.
func (c *Canvas) Abs(n float64) float64 {
	return c.p5Instance.Call("abs", n).Float()
}

// Ceil returns the smallest integer greater than or equal to the given number.
func (c *Canvas) Ceil(n float64) float64 {
	return c.p5Instance.Call("ceil", n).Float()
}

// Constrain limits a number to be within a specified range.
func (c *Canvas) Constrain(n, low, high float64) float64 {
	return c.p5Instance.Call("constrain", n, low, high).Float()
}

// Dist calculates the distance between two points.
func (c *Canvas) Dist(x1, y1, x2, y2 float64) float64 {
	return c.p5Instance.Call("dist", x1, y1, x2, y2).Float()
}

// Exp returns Euler's number e raised to the power of the given number.
func (c *Canvas) Exp(n float64) float64 {
	return c.p5Instance.Call("exp", n).Float()
}

// Floor returns the largest integer less than or equal to the given number.
func (c *Canvas) Floor(n float64) float64 {
	return c.p5Instance.Call("floor", n).Float()
}

// Lerp performs a linear interpolation between two values.
func (c *Canvas) Lerp(start, stop, amt float64) float64 {
	return c.p5Instance.Call("lerp", start, stop, amt).Float()
}

// Log returns the natural logarithm (base e) of the given number.
func (c *Canvas) Log(n float64) float64 {
	return c.p5Instance.Call("log", n).Float()
}

// Mag calculates the magnitude of a vector.
func (c *Canvas) Mag(x, y float64) float64 {
	return c.p5Instance.Call("mag", x, y).Float()
}

// Max returns the largest value from a list of numbers.
func (c *Canvas) Max(args ...float64) float64 {
	values := make([]any, len(args))
	for i, v := range args {
		values[i] = v
	}
	return c.p5Instance.Call("max", values...).Float()
}

// Min returns the smallest value from a list of numbers.
func (c *Canvas) Min(args ...float64) float64 {
	values := make([]any, len(args))
	for i, v := range args {
		values[i] = v
	}
	return c.p5Instance.Call("min", values...).Float()
}

// Norm normalizes a number from another range into a value between 0 and 1.
func (c *Canvas) Norm(value, start, stop float64) float64 {
	return c.p5Instance.Call("norm", value, start, stop).Float()
}

// Pow returns the result of raising a number to a power.
func (c *Canvas) Pow(n, e float64) float64 {
	return c.p5Instance.Call("pow", n, e).Float()
}

// Round returns the nearest integer to the given number.
func (c *Canvas) Round(n float64) float64 {
	return c.p5Instance.Call("round", n).Float()
}

// Sq returns the square of the given number.
func (c *Canvas) Sq(n float64) float64 {
	return c.p5Instance.Call("sq", n).Float()
}

// Sqrt returns the square root of the given number.
func (c *Canvas) Sqrt(n float64) float64 {
	return c.p5Instance.Call("sqrt", n).Float()
}

// CreateGraphics creates and returns a new graphics buffer.
func (c *Canvas) CreateGraphics(w, h float64, renderer ...string) js.Value {
	if len(renderer) > 0 {
		return c.p5Instance.Call("createGraphics", w, h, renderer[0])
	}
	return c.p5Instance.Call("createGraphics", w, h)
}

// BlendMode sets the blending mode for the canvas.
func (c *Canvas) BlendMode(mode string) {
	c.p5Instance.Call("blendMode", mode)
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
func (c *Canvas) Filter(filterType string, value ...float64) {
	if len(value) > 0 {
		c.p5Instance.Call("filter", filterType, value[0])
	} else {
		c.p5Instance.Call("filter", filterType)
	}
}

// Blend blends a region of pixels using a specified blend mode.
func (c *Canvas) Blend(sx, sy, sw, sh, dx, dy, dw, dh float64, blendMode string) {
	c.p5Instance.Call("blend", sx, sy, sw, sh, dx, dy, dw, dh, blendMode)
}

// Mask applies an image as a mask to the canvas.
func (c *Canvas) Mask(img js.Value) {
	c.p5Instance.Call("mask", img)
}

// EllipseMode sets the location from which ellipses are drawn.
func (c *Canvas) EllipseMode(mode string) {
	c.p5Instance.Call("ellipseMode", mode)
}

// RectMode sets the location from which rectangles are drawn.
func (c *Canvas) RectMode(mode string) {
	c.p5Instance.Call("rectMode", mode)
}

// StrokeJoin sets the style of the joints which connect line segments.
func (c *Canvas) StrokeJoin(join string) {
	c.p5Instance.Call("strokeJoin", join)
}

// Smooth draws all geometry with smooth (anti-aliased) edges.
func (c *Canvas) Smooth() {
	c.p5Instance.Call("smooth")
}

// NoSmooth draws all geometry with jagged (aliased) edges.
func (c *Canvas) NoSmooth() {
	c.p5Instance.Call("noSmooth")
}
