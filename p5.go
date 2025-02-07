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

type Func func(c *Canvas)

func Preload(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["preload"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func Setup(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["setup"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func Draw(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["draw"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MouseMoved(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseMoved"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MouseDragged(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseDragged"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MousePressed(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mousePressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MouseReleased(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MouseClicked(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func DoubleClicked(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["doubleClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func MouseWheel(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["mouseWheel"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func KeyPressed(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyPressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func KeyReleased(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

func KeyTyped(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["keyTyped"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

type Canvas struct {
	p5Instance   js.Value
	funcHandlers map[string]js.Func
}

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

func (c *Canvas) CreateCanvas(w, h int) {
	c.p5Instance.Call("createCanvas", w, h)
}

func (c *Canvas) Background(args ...any) {
	c.p5Instance.Call("background", args...)
}

func (c *Canvas) Fill(args ...any) {
	c.p5Instance.Call("fill", args...)
}

func (c *Canvas) Stroke(args ...any) {
	c.p5Instance.Call("stroke", args...)
}

func (c *Canvas) NoFill() {
	c.p5Instance.Call("noFill")
}

func (c *Canvas) NoStroke() {
	c.p5Instance.Call("noStroke")
}

func (c *Canvas) Ellipse(x, y, w, h float64) {
	c.p5Instance.Call("ellipse", x, y, w, h)
}

func (c *Canvas) Rect(x, y, w, h float64) {
	c.p5Instance.Call("rect", x, y, w, h)
}

func (c *Canvas) Line(x1, y1, x2, y2 float64) {
	c.p5Instance.Call("line", x1, y1, x2, y2)
}

func (c *Canvas) Triangle(x1, y1, x2, y2, x3, y3 float64) {
	c.p5Instance.Call("triangle", x1, y1, x2, y2, x3, y3)
}

func (c *Canvas) Point(x, y float64) {
	c.p5Instance.Call("point", x, y)
}

func (c *Canvas) Arc(x, y, w, h, start, stop float64) {
	c.p5Instance.Call("arc", x, y, w, h, start, stop)
}

func (c *Canvas) Bezier(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	c.p5Instance.Call("bezier", x1, y1, x2, y2, x3, y3, x4, y4)
}

func (c *Canvas) QuadraticVertex(cx, cy, x, y float64) {
	c.p5Instance.Call("quadraticVertex", cx, cy, x, y)
}

func (c *Canvas) Curve(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	c.p5Instance.Call("curve", x1, y1, x2, y2, x3, y3, x4, y4)
}

func (c *Canvas) Text(str string, x, y float64) {
	c.p5Instance.Call("text", str, x, y)
}

func (c *Canvas) TextFont(font string, size float64) {
	c.p5Instance.Call("textFont", font, size)
}

func (c *Canvas) TextSize(size float64) {
	c.p5Instance.Call("textSize", size)
}

func (c *Canvas) Push() {
	c.p5Instance.Call("push")
}

func (c *Canvas) Pop() {
	c.p5Instance.Call("pop")
}

func (c *Canvas) Translate(x, y float64) {
	c.p5Instance.Call("translate", x, y)
}

func (c *Canvas) Rotate(angle float64) {
	c.p5Instance.Call("rotate", angle)
}

func (c *Canvas) RotateX(angle float64) {
	c.p5Instance.Call("rotateX", angle)
}

func (c *Canvas) RotateY(angle float64) {
	c.p5Instance.Call("rotateY", angle)
}

func (c *Canvas) RotateZ(angle float64) {
	c.p5Instance.Call("rotateZ", angle)
}

func (c *Canvas) Scale(s float64) {
	c.p5Instance.Call("scale", s)
}

func (c *Canvas) ShearX(angle float64) {
	c.p5Instance.Call("shearX", angle)
}

func (c *Canvas) ShearY(angle float64) {
	c.p5Instance.Call("shearY", angle)
}

func (c *Canvas) SaveCanvas(filename, extension string) {
	c.p5Instance.Call("saveCanvas", filename, extension)
}

func (c *Canvas) LoadImage(path string) js.Value {
	return c.p5Instance.Call("loadImage", path)
}

func (c *Canvas) Image(img js.Value, x, y, w, h float64) {
	c.p5Instance.Call("image", img, x, y, w, h)
}

func (c *Canvas) FrameRate(fps float64) {
	c.p5Instance.Call("frameRate", fps)
}

func (c *Canvas) Random(min, max float64) float64 {
	return c.p5Instance.Call("random", min, max).Float()
}

func (c *Canvas) Map(value, start1, stop1, start2, stop2 float64) float64 {
	return c.p5Instance.Call("map", value, start1, stop1, start2, stop2).Float()
}

func (c *Canvas) BeginShape(option ...any) {
	if len(option) > 0 {
		c.p5Instance.Call("beginShape", option[0])
	} else {
		c.p5Instance.Call("beginShape")
	}
}

func (c *Canvas) Vertex(x, y float64) {
	c.p5Instance.Call("vertex", x, y)
}

func (c *Canvas) EndShape(mode ...string) {
	if len(mode) > 0 {
		c.p5Instance.Call("endShape", mode[0])
	} else {
		c.p5Instance.Call("endShape")
	}
}

func (c *Canvas) BezierVertex(cx1, cy1, cx2, cy2, x, y float64) {
	c.p5Instance.Call("bezierVertex", cx1, cy1, cx2, cy2, x, y)
}

func (c *Canvas) CurveVertex(x, y float64) {
	c.p5Instance.Call("curveVertex", x, y)
}

func (c *Canvas) BeginContour() {
	c.p5Instance.Call("beginContour")
}

func (c *Canvas) EndContour() {
	c.p5Instance.Call("endContour")
}

func (c *Canvas) Close() {
	c.p5Instance.Call("close")
}

func (c *Canvas) TextAlign(align string) {
	c.p5Instance.Call("textAlign", align)
}

func (c *Canvas) TextWrap(w string) {
	c.p5Instance.Call("textWrap", w)
}

func (c *Canvas) MouseX() float64 {
	return c.p5Instance.Get("mouseX").Float()
}

func (c *Canvas) PMouseX() float64 {
	return c.p5Instance.Get("pmouseX").Float()
}

func (c *Canvas) MouseY() float64 {
	return c.p5Instance.Get("mouseY").Float()
}

func (c *Canvas) PMouseY() float64 {
	return c.p5Instance.Get("pmouseY").Float()
}

func (c *Canvas) MouseIsPressed() bool {
	return c.p5Instance.Get("mouseIsPressed").Bool()
}

func (c *Canvas) MovedX() float64 {
	return c.p5Instance.Get("movedX").Float()
}

func (c *Canvas) MovedY() float64 {
	return c.p5Instance.Get("movedY").Float()
}

func (c *Canvas) MouseButton() string {
	return c.p5Instance.Get("mouseButton").String()
}

func (c *Canvas) SaveGif(name string, second float64) {
	c.p5Instance.Call("saveGif", name, second)
}

func (c *Canvas) Key() string {
	return c.p5Instance.Get("key").String()
}

func (c *Canvas) KeyCode() int {
	return c.p5Instance.Get("keyCode").Int()
}

func (c *Canvas) KeyIsPressed() bool {
	return c.p5Instance.Get("keyIsPressed").Bool()
}

func (c *Canvas) ColorMode(mode string, max ...float64) {
	if len(max) > 0 {
		c.p5Instance.Call("colorMode", mode, max[0])
	} else {
		c.p5Instance.Call("colorMode", mode)
	}
}

func (c *Canvas) Acos(value float64) float64 {
	return c.p5Instance.Call("acos", value).Float()
}

func (c *Canvas) Cos(value float64) float64 {
	return c.p5Instance.Call("cos", value).Float()
}

func (c *Canvas) AngleMode(mode string) {
	c.p5Instance.Call("angleMode", mode)
}

func (c *Canvas) Asin(value float64) float64 {
	return c.p5Instance.Call("asin", value).Float()
}

func (c *Canvas) Atan(value float64) float64 {
	return c.p5Instance.Call("atan", value).Float()
}

func (c *Canvas) Atan2(y, x float64) float64 {
	return c.p5Instance.Call("atan2", y, x).Float()
}

func (c *Canvas) Sin(value float64) float64 {
	return c.p5Instance.Call("sin", value).Float()
}

func (c *Canvas) Tan(value float64) float64 {
	return c.p5Instance.Call("tan", value).Float()
}

func (c *Canvas) Degrees(value float64) float64 {
	return c.p5Instance.Call("degrees", value).Float()
}

func (c *Canvas) Radians(value float64) float64 {
	return c.p5Instance.Call("radians", value).Float()
}

func (c *Canvas) StrokeWeight(weight float64) {
	c.p5Instance.Call("strokeWeight", weight)
}

func (c *Canvas) StrokeCap(cap string) {
	c.p5Instance.Call("strokeCap", cap)
}

func (c *Canvas) Erase(opt ...any) {
	c.p5Instance.Call("erase", opt...)
}

func (c *Canvas) NoErase() {
	c.p5Instance.Call("noErase")
}

func (c *Canvas) CreateGraphics(width, height int, opt ...any) js.Value {
	if len(opt) > 0 {
		return c.p5Instance.Call("createGraphics", width, height, opt...)
	}
	return c.p5Instance.Call("createGraphics", width, height)
}
