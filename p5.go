package p5go

import (
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
	p5Instance js.Value
	global     js.Value
)

func init() {
	global = js.Global()
}

// Init initializes the p5 instance
func Init(query string) {
	document := global.Get("document")
	container := document.Call("querySelector", query)
	container.Set("innerHTML", "")

	p5Constructor := global.Get("p5")
	sketch := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		p5Instance = args[0]
		setupSketch()
		return nil
	})

	p5Constructor.New(sketch, container)
}

func setupSketch() {
	if jsPreload := global.Get("preload"); jsPreload.Type() != js.TypeUndefined {
		p5Instance.Set("preload", preload)
	}
	if jsSetup := global.Get("setup"); jsSetup.Type() != js.TypeUndefined {
		p5Instance.Set("setup", setup)
	}
	if jsDraw := global.Get("draw"); jsDraw.Type() != js.TypeUndefined {
		p5Instance.Set("draw", draw)
	}
	// Event handlers
	setupEventHandler("mouseMoved")
	setupEventHandler("mouseDragged")
	setupEventHandler("mousePressed")
	setupEventHandler("mouseReleased")
	setupEventHandler("mouseClicked")
	setupEventHandler("doubleClicked")
	setupEventHandler("mouseWheel")
	setupEventHandler("keyPressed")
	setupEventHandler("keyReleased")
	setupEventHandler("keyTyped")
}

func setupEventHandler(name string) {
	if jsHandler := global.Get(name); jsHandler.Type() != js.TypeUndefined {
		p5Instance.Set(name, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			switch name {
			case "mouseMoved":
				mouseMoved.Invoke(args[0])
			case "mouseDragged":
				mouseDragged.Invoke(args[0])
			case "mousePressed":
				mousePressed.Invoke(args[0])
			case "mouseReleased":
				mouseReleased.Invoke(args[0])
			case "mouseClicked":
				mouseClicked.Invoke(args[0])
			case "doubleClicked":
				doubleClicked.Invoke(args[0])
			case "mouseWheel":
				mouseWheel.Invoke(args[0])
			case "keyPressed":
				keyPressed.Invoke(args[0])
			case "keyReleased":
				keyReleased.Invoke(args[0])
			case "keyTyped":
				keyTyped.Invoke(args[0])
			}
			return nil
		}))
	}
}

// Basic p5.js function wrappers
func CreateCanvas(w, h int) {
	p5Instance.Call("createCanvas", w, h)
}

func Background(args ...interface{}) {
	p5Instance.Call("background", args...)
}

func Fill(args ...interface{}) {
	p5Instance.Call("fill", args...)
}

func Stroke(args ...interface{}) {
	p5Instance.Call("stroke", args...)
}

func NoFill() {
	p5Instance.Call("noFill")
}

func NoStroke() {
	p5Instance.Call("noStroke")
}

func Ellipse(x, y, w, h float64) {
	p5Instance.Call("ellipse", x, y, w, h)
}

func Rect(x, y, w, h float64) {
	p5Instance.Call("rect", x, y, w, h)
}

func Line(x1, y1, x2, y2 float64) {
	p5Instance.Call("line", x1, y1, x2, y2)
}

func Triangle(x1, y1, x2, y2, x3, y3 float64) {
	p5Instance.Call("triangle", x1, y1, x2, y2, x3, y3)
}

func Point(x, y float64) {
	p5Instance.Call("point", x, y)
}

// Math functions
func Random(min, max float64) float64 {
	return p5Instance.Call("random", min, max).Float()
}

func Map(value, start1, stop1, start2, stop2 float64) float64 {
	return p5Instance.Call("map", value, start1, stop1, start2, stop2).Float()
}

// Color functions
func Color(args ...interface{}) js.Value {
	return p5Instance.Call("color", args...)
}

// Vector represents a p5.Vector
type Vector struct {
	v js.Value
}

func CreateVector(x, y, z float64) Vector {
	return Vector{v: p5Instance.Call("createVector", x, y, z)}
}

func (v Vector) X() float64 {
	return v.v.Get("x").Float()
}

func (v Vector) Y() float64 {
	return v.v.Get("y").Float()
}

func (v Vector) Z() float64 {
	return v.v.Get("z").Float()
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v: v.v.Call("add", other.v)}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{v: v.v.Call("sub", other.v)}
}

func (v Vector) Mult(n float64) Vector {
	return Vector{v: v.v.Call("mult", n)}
}

func (v Vector) Div(n float64) Vector {
	return Vector{v: v.v.Call("div", n)}
}

func (v Vector) Mag() float64 {
	return v.v.Call("mag").Float()
}

func (v Vector) MagSq() float64 {
	return v.v.Call("magSq").Float()
}

func (v Vector) Dot(other Vector) float64 {
	return v.v.Call("dot", other.v).Float()
}

func (v Vector) Cross(other Vector) Vector {
	return Vector{v: v.v.Call("cross", other.v)}
}

func (v Vector) Dist(other Vector) float64 {
	return v.v.Call("dist", other.v).Float()
}

func (v Vector) Normalize() Vector {
	return Vector{v: v.v.Call("normalize")}
}

func (v Vector) Limit(max float64) Vector {
	return Vector{v: v.v.Call("limit", max)}
}

func (v Vector) SetMag(len float64) Vector {
	return Vector{v: v.v.Call("setMag", len)}
}

func (v Vector) Heading() float64 {
	return v.v.Call("heading").Float()
}

func (v Vector) Rotate(angle float64) Vector {
	return Vector{v: v.v.Call("rotate", angle)}
}

func (v Vector) Lerp(other Vector, amt float64) Vector {
	return Vector{v: v.v.Call("lerp", other.v, amt)}
}

func (v Vector) Equals(other Vector) bool {
	return v.v.Call("equals", other.v).Bool()
}

// Event handlers (to be implemented in user code)
var (
	preload       js.Func
	setup         js.Func
	draw          js.Func
	mouseMoved    js.Func
	mouseDragged  js.Func
	mousePressed  js.Func
	mouseReleased js.Func
	mouseClicked  js.Func
	doubleClicked js.Func
	mouseWheel    js.Func
	keyPressed    js.Func
	keyReleased   js.Func
	keyTyped      js.Func
)

func SetPreload(f func()) {
	preload = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
}

func SetSetup(f func()) {
	setup = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
}

func SetDraw(f func()) {
	draw = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
}

func SetMouseMoved(f func(event js.Value)) {
	mouseMoved = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetMouseDragged(f func(event js.Value)) {
	mouseDragged = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetMousePressed(f func(event js.Value)) {
	mousePressed = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetMouseReleased(f func(event js.Value)) {
	mouseReleased = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetMouseClicked(f func(event js.Value)) {
	mouseClicked = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetDoubleClicked(f func(event js.Value)) {
	doubleClicked = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetMouseWheel(f func(event js.Value)) {
	mouseWheel = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetKeyPressed(f func(event js.Value)) {
	keyPressed = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetKeyReleased(f func(event js.Value)) {
	keyReleased = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

func SetKeyTyped(f func(event js.Value)) {
	keyTyped = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0])
		return nil
	})
}

// Example usage:
//
// func main() {
// 	p5.SetSetup(func() {
// 		p5.CreateCanvas(400, 400)
// 	})
//
// 	p5.SetDraw(func() {
// 		p5.Background(220)
// 		p5.Fill(255, 0, 0)
// 		p5.Ellipse(200, 200, 50, 50)
// 	})
//
// 	p5.SetMousePressed(func(event js.Value) {
// 		x := event.Get("mouseX").Float()
// 		y := event.Get("mouseY").Float()
// 		p5.Ellipse(x, y, 20, 20)
// 	})
//
// 	p5.Init("main")
//
// 	// Prevent the program from exiting
// 	select {}
//
