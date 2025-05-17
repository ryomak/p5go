// Package p5go provides a bridge between Go and p5.js.
package p5go

import (
	"errors"
	"fmt"
	"math"
	"syscall/js"
)

// RendererMode represents the rendering mode for the canvas
type RendererMode string

const (
	// Renderer modes
	P2D   RendererMode = "p2d"
	WEBGL RendererMode = "webgl"
)

// CursorStyle represents the cursor style
type CursorStyle string

const (
	// Environment
	ARROW CursorStyle = "default"
	CROSS CursorStyle = "crosshair"
	HAND  CursorStyle = "pointer"
	MOVE  CursorStyle = "move"
	TEXT  CursorStyle = "text"
	WAIT  CursorStyle = "wait"
)

// AngleMode represents the angle mode
type AngleMode string

const (
	// Trigonometry
	PI                   = math.Pi
	HALF_PI              = math.Pi / 2
	QUARTER_PI           = math.Pi / 4
	TWO_PI               = math.Pi * 2
	TAU                  = TWO_PI
	DEGREES    AngleMode = "degrees"
	RADIANS    AngleMode = "radians"
)

// ColorMode represents the color mode
type ColorMode string

const (
	// Color modes
	RGB ColorMode = "rgb"
	HSB ColorMode = "hsb"
	HSL ColorMode = "hsl"
)

// DrawingMode represents the drawing mode
type DrawingMode string

const (
	// Drawing modes
	CORNER   DrawingMode = "corner"
	CORNERS  DrawingMode = "corners"
	RADIUS   DrawingMode = "radius"
	RIGHT    DrawingMode = "right"
	LEFT     DrawingMode = "left"
	CENTER   DrawingMode = "center"
	TOP      DrawingMode = "top"
	BOTTOM   DrawingMode = "bottom"
	BASELINE DrawingMode = "alphabetic"
)

// ShapeType represents the type of shape
// P5.jsの beginShape で使う型
// https://p5js.org/reference/#/p5/beginShape
// "POINTS", "LINES", "TRIANGLES", "TRIANGLE_FAN", "TRIANGLE_STRIP", "QUADS", "QUAD_STRIP", "TESS"
type ShapeType string

const (
	POINTS         ShapeType = "POINTS"
	LINES          ShapeType = "LINES"
	LINE_STRIP     ShapeType = "LINE_STRIP"
	LINE_LOOP      ShapeType = "LINE_LOOP"
	TRIANGLES      ShapeType = "TRIANGLES"
	TRIANGLE_FAN   ShapeType = "TRIANGLE_FAN"
	TRIANGLE_STRIP ShapeType = "TRIANGLE_STRIP"
	QUADS          ShapeType = "QUADS"
	QUAD_STRIP     ShapeType = "QUAD_STRIP"
	TESS           ShapeType = "TESS"
	CLOSE          ShapeType = "CLOSE"
	OPEN           ShapeType = "OPEN"
	CHORD          ShapeType = "CHORD"
	PIE            ShapeType = "PIE"
	PROJECT        ShapeType = "PROJECT"
	SQUARE         ShapeType = "SQUARE"
	ROUND          ShapeType = "ROUND"
	BEVEL          ShapeType = "BEVEL"
	MITER          ShapeType = "MITER"
)

// BlendMode represents the blending mode
type BlendMode string

const (
	// Blend modes
	BLEND      BlendMode = "source-over"
	REMOVE     BlendMode = "destination-out"
	ADD        BlendMode = "lighter"
	DARKEST    BlendMode = "darken"
	LIGHTEST   BlendMode = "lighten"
	DIFFERENCE BlendMode = "difference"
	SUBTRACT   BlendMode = "subtract"
	EXCLUSION  BlendMode = "exclusion"
	MULTIPLY   BlendMode = "multiply"
	SCREEN     BlendMode = "screen"
	REPLACE    BlendMode = "copy"
	OVERLAY    BlendMode = "overlay"
	HARD_LIGHT BlendMode = "hard-light"
	SOFT_LIGHT BlendMode = "soft-light"
	DODGE      BlendMode = "color-dodge"
	BURN       BlendMode = "color-burn"
)

// FilterType represents the type of filter
type FilterType string

const (
	// Image filters
	THRESHOLD FilterType = "threshold"
	GRAY      FilterType = "gray"
	OPAQUE    FilterType = "opaque"
	INVERT    FilterType = "invert"
	POSTERIZE FilterType = "posterize"
	DILATE    FilterType = "dilate"
	ERODE     FilterType = "erode"
	BLUR      FilterType = "blur"
)

// TextStyle represents the style of text
type TextStyle string

const (
	// Typography
	NORMAL     TextStyle = "normal"
	ITALIC     TextStyle = "italic"
	BOLD       TextStyle = "bold"
	BOLDITALIC TextStyle = "bold italic"
)

// WebGLMode represents the WebGL mode
type WebGLMode string

const (
	// Web GL specific
	IMMEDIATE WebGLMode = "immediate"
	IMAGE     WebGLMode = "image"
	NEAREST   WebGLMode = "nearest"
	REPEAT    WebGLMode = "repeat"
	CLAMP     WebGLMode = "clamp"
	MIRROR    WebGLMode = "mirror"
)

// Orientation represents the device orientation
type Orientation string

const (
	// Device orientation
	LANDSCAPE Orientation = "landscape"
	PORTRAIT  Orientation = "portrait"
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

// MouseDragged sets the mouseDragged handler with a MouseDraggedEvent
func MouseDragged(handler MouseDraggedHandler) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseDragged"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			e := MouseDraggedEvent{
				X:       c.MouseX(),
				Y:       c.MouseY(),
				Button:  c.MouseButton(),
				Pressed: c.MouseIsPressed(),
			}
			handler(c, e)
			return nil
		})
	}
}

// MousePressed sets the mousePressed handler with a MouseEvent
func MousePressed(handler MousePressedHandler) Func {
	return func(c *Canvas) {
		c.funcHandlers["mousePressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			e := MouseEvent{
				X:       c.MouseX(),
				Y:       c.MouseY(),
				Button:  c.MouseButton(),
				Pressed: c.MouseIsPressed(),
			}
			handler(c, e)
			return nil
		})
	}
}

// MouseReleased sets the mouseReleased handler with a MouseReleasedEvent
func MouseReleased(handler MouseReleasedHandler) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			e := MouseReleasedEvent{
				X:       c.MouseX(),
				Y:       c.MouseY(),
				Button:  c.MouseButton(),
				Pressed: c.MouseIsPressed(),
			}
			handler(c, e)
			return nil
		})
	}
}

// MouseClicked sets the mouseClicked handler with a MouseClickedEvent
func MouseClicked(handler MouseClickedHandler) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			e := MouseClickedEvent{
				X:       c.MouseX(),
				Y:       c.MouseY(),
				Button:  c.MouseButton(),
				Pressed: c.MouseIsPressed(),
			}
			handler(c, e)
			return nil
		})
	}
}

// DoubleClicked sets the doubleClicked handler with a DoubleClickedEvent
func DoubleClicked(handler DoubleClickedHandler) Func {
	return func(c *Canvas) {
		c.funcHandlers["doubleClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			e := DoubleClickedEvent{
				X:       c.MouseX(),
				Y:       c.MouseY(),
				Button:  c.MouseButton(),
				Pressed: c.MouseIsPressed(),
			}
			handler(c, e)
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
func (c *Canvas) CreateCanvas(w, h int, opts ...any) {
	c.width = float64(w)
	c.height = float64(h)
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
func (c *Canvas) Point(x, y float64, z ...float64) {
	if len(z) > 0 {
		c.p5Instance.Call("point", x, y, z[0])
	} else {
		c.p5Instance.Call("point", x, y)
	}
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
func (c *Canvas) Image(img any, opts ...any) {
	c.p5Instance.Call("image", append([]any{img}, opts...)...)
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
func (c *Canvas) BeginShape(kind ...ShapeType) {
	if len(kind) > 0 {
		c.p5Instance.Call("beginShape", js.Global().Get(string(kind[0])))
	} else {
		c.p5Instance.Call("beginShape")
	}
}

// Vertex adds a vertex to the current shape.
func (c *Canvas) Vertex(x, y float64) {
	c.p5Instance.Call("vertex", x, y)
}

// EndShape ends recording vertices for a shape.
func (c *Canvas) EndShape(mode ...ShapeType) {
	if len(mode) > 0 {
		c.p5Instance.Call("endShape", js.Global().Get(string(mode[0])))
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
func (c *Canvas) TextAlign(align DrawingMode) {
	c.p5Instance.Call("textAlign", string(align))
}

// TextWrap sets the wrap mode for text.
func (c *Canvas) TextWrap(w DrawingMode) {
	c.p5Instance.Call("textWrap", string(w))
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
func (c *Canvas) ColorMode(mode ColorMode, max ...float64) {
	if len(max) > 0 {
		c.p5Instance.Call("colorMode", string(mode), max[0])
	} else {
		c.p5Instance.Call("colorMode", string(mode))
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
func (c *Canvas) AngleMode(mode AngleMode) {
	c.p5Instance.Call("angleMode", string(mode))
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
func (c *Canvas) StrokeCap(cap ShapeType) {
	c.p5Instance.Call("strokeCap", string(cap))
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
func (c *Canvas) TextStyle(style TextStyle) {
	c.p5Instance.Call("textStyle", string(style))
}

// TextWidth returns the width of the specified text.
func (c *Canvas) TextWidth(text string) float64 {
	return c.p5Instance.Call("textWidth", text).Float()
}

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

// Width returns the width of the canvas.
func (c *Canvas) Width() float64 {
	return c.width
}

// Height returns the height of the canvas.
func (c *Canvas) Height() float64 {
	return c.height
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
func (c *Canvas) CreateGraphics(w, h float64, renderer ...any) *Canvas {
	return &Canvas{p5Instance: c.p5Instance.Call("createGraphics", append([]any{w, h}, renderer...)...)}
}

// BlendMode sets the blending mode for the canvas.
func (c *Canvas) BlendMode(mode BlendMode) {
	c.p5Instance.Call("blendMode", string(mode))
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

// EllipseMode sets the location from which ellipses are drawn.
func (c *Canvas) EllipseMode(mode DrawingMode) {
	c.p5Instance.Call("ellipseMode", string(mode))
}

// RectMode sets the location from which rectangles are drawn.
func (c *Canvas) RectMode(mode DrawingMode) {
	c.p5Instance.Call("rectMode", string(mode))
}

// StrokeJoin sets the style of the joints which connect line segments.
func (c *Canvas) StrokeJoin(join ShapeType) {
	c.p5Instance.Call("strokeJoin", string(join))
}

// Smooth draws all geometry with smooth (anti-aliased) edges.
func (c *Canvas) Smooth() {
	c.p5Instance.Call("smooth")
}

// NoSmooth draws all geometry with jagged (aliased) edges.
func (c *Canvas) NoSmooth() {
	c.p5Instance.Call("noSmooth")
}

// CaptureKind is a type that represents the kind of capture.
type CaptureKind string

const (
	CaptureKindVIDEO CaptureKind = "VIDEO"
	CaptureKindIMAGE CaptureKind = "IMAGE"
)

// CreateCapture creates a capture object.
func (c *Canvas) CreateCapture(kind CaptureKind) js.Value {
	return c.p5Instance.Call("createCapture", string(kind))
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

// Color represents a color with RGBA components
type Color struct {
	R, G, B, A float64
}

// Vector represents a 2D vector
type Vector struct {
	X, Y float64
}

// Rectangle represents a rectangle with position and size
type Rectangle struct {
	Position Vector
	Size     Vector
}

// Circle represents a circle with center position and diameter
type Circle struct {
	Position Vector
	Diameter float64
}

// Line represents a line with start and end points
type Line struct {
	Start, End Vector
}

// Triangle represents a triangle with three vertices
type Triangle struct {
	V1, V2, V3 Vector
}

// FillRGB sets the fill color using RGB values
func (c *Canvas) FillRGB(r, g, b float64) {
	c.Fill(r, g, b)
}

// FillRGBA sets the fill color using RGBA values
func (c *Canvas) FillRGBA(r, g, b, a float64) {
	c.Fill(r, g, b, a)
}

// FillColor sets the fill color using a Color struct
func (c *Canvas) FillColor(color Color) {
	c.Fill(color.R, color.G, color.B, color.A)
}

// StrokeRGB sets the stroke color using RGB values
func (c *Canvas) StrokeRGB(r, g, b float64) {
	c.Stroke(r, g, b)
}

// StrokeRGBA sets the stroke color using RGBA values
func (c *Canvas) StrokeRGBA(r, g, b, a float64) {
	c.Stroke(r, g, b, a)
}

// StrokeColor sets the stroke color using a Color struct
func (c *Canvas) StrokeColor(color Color) {
	c.Stroke(color.R, color.G, color.B, color.A)
}

// DrawRect draws a rectangle using a Rectangle struct
func (c *Canvas) DrawRect(r Rectangle) {
	c.Rect(r.Position.X, r.Position.Y, r.Size.X, r.Size.Y)
}

// DrawCircle draws a circle using a Circle struct
func (c *Canvas) DrawCircle(circle Circle) {
	c.Circle(circle.Position.X, circle.Position.Y, circle.Diameter)
}

// DrawLine draws a line using a Line struct
func (c *Canvas) DrawLine(line Line) {
	c.Line(line.Start.X, line.Start.Y, line.End.X, line.End.Y)
}

// DrawTriangle draws a triangle using a Triangle struct
func (c *Canvas) DrawTriangle(t Triangle) {
	c.Triangle(t.V1.X, t.V1.Y, t.V2.X, t.V2.Y, t.V3.X, t.V3.Y)
}

// OrbitControl represents a control for orbiting around an object
func (c *Canvas) OrbitControl(opts ...any) {
	c.p5Instance.Call("orbitControl", opts...)
}

// MouseEvent represents a mouse event
type MouseEvent struct {
	X, Y    float64
	Button  string
	Pressed bool
}

// MousePressedHandler is a type for mouse pressed event handlers
type MousePressedHandler func(c *Canvas, e MouseEvent)

// MouseDraggedEvent represents a mouse dragged event
type MouseDraggedEvent struct {
	X, Y    float64
	Button  string
	Pressed bool
}

// MouseDraggedHandler is a type for mouse dragged event handlers
type MouseDraggedHandler func(c *Canvas, e MouseDraggedEvent)

// MouseReleasedEvent represents a mouse released event
type MouseReleasedEvent struct {
	X, Y    float64
	Button  string
	Pressed bool
}

// MouseReleasedHandler is a type for mouse released event handlers
type MouseReleasedHandler func(c *Canvas, e MouseReleasedEvent)

// MouseClickedEvent represents a mouse clicked event
type MouseClickedEvent struct {
	X, Y    float64
	Button  string
	Pressed bool
}

// MouseClickedHandler is a type for mouse clicked event handlers
type MouseClickedHandler func(c *Canvas, e MouseClickedEvent)

// DoubleClickedEvent represents a double clicked event
type DoubleClickedEvent struct {
	X, Y    float64
	Button  string
	Pressed bool
}

// DoubleClickedHandler is a type for double clicked event handlers
type DoubleClickedHandler func(c *Canvas, e DoubleClickedEvent)
