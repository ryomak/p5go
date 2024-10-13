// Package p5go provides a bridge between Go and p5.js, allowing you to create interactive visuals in Go
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
	global     = js.Global()
)

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
			if handler, ok := eventHandlers[name]; ok {
				handler(args[0])
			}
			return nil
		}))
	}
}

// p5.js wrappers
func CreateCanvas(w, h int) {
	p5Instance.Call("createCanvas", w, h)
}

func Background(args ...interface{}) {
	p5Instance.Call("background", args...)
}

func Fill(args ...interface{}) {
	p5Instance.Call("fill", args...)
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

// Event handlers to be implemented in user code
var eventHandlers = make(map[string]func(js.Value))

func SetEventHandler(eventName string, handler func(event js.Value)) {
	eventHandlers[eventName] = handler
}

// Example usage:
//
// func main() {
// 	p5go.SetEventHandler("setup", func(event js.Value) {
// 		p5go.CreateCanvas(400, 400)
// 	})
//
// 	p5go.SetEventHandler("draw", func(event js.Value) {
// 		p5go.Background(220)
// 		p5go.Fill(255, 0, 0)
// 		p5go.Ellipse(200, 200, 50, 50)
// 	})
//
// 	p5go.Init("main")
//
// 	// Prevent the program from exiting
// 	select {}
// }
