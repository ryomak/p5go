// Package p5go provides a bridge between Go and p5.js, allowing you to create interactive visuals in Go
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

// Execute initializes the p5 p5Instance
func Execute(query string, opts ...Option) error {
	// Get container
	container := global.Get("document").Call("querySelector", query)
	if container.IsNull() {
		return errors.New(fmt.Sprintf("%s is not match", query))
	}
	container.Set("innerHTML", "")

	p5 := &P5Instance{
		p5Instance:   js.Undefined(),
		funcHandlers: map[string]js.Func{},
	}

	sketch := js.FuncOf(func(this js.Value, args []js.Value) any {
		p5.p5Instance = args[0]
		for _, opt := range opts {
			opt(p5)
		}

		for method, handler := range p5.funcHandlers {
			p5.p5Instance.Set(method, handler)
		}
		return nil
	})

	p5Constructor := global.Get("p5")
	p5Constructor.New(sketch, container)

	if err := p5.Validate(); err != nil {
		return err
	}

	return nil
}

type Option func(p *P5Instance)

func Preload(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["preload"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func Setup(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["setup"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func Draw(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["draw"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MouseMoved(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mouseMoved"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MouseDragged(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mouseDragged"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MousePressed(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mousePressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MouseReleased(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mouseReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MouseClicked(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mouseClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func DoubleClicked(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["doubleClicked"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func MouseWheel(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["mouseWheel"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func KeyPressed(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["keyPressed"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func KeyReleased(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["keyReleased"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

func KeyTyped(handler func(p *P5Instance)) Option {
	return func(p *P5Instance) {
		p.funcHandlers["keyTyped"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(p)
			return nil
		})
	}
}

type P5Instance struct {
	p5Instance   js.Value
	funcHandlers map[string]js.Func
}

func (p *P5Instance) Validate() error {
	if p.p5Instance.Type() == js.TypeUndefined {
		return errors.New("p5.js is not loaded")
	}
	if p.funcHandlers["setup"].IsUndefined() {
		return errors.New("setup function is not defined")
	}
	if p.funcHandlers["draw"].IsUndefined() {
		return errors.New("draw function is not defined")
	}
	return nil
}

func (p *P5Instance) CreateCanvas(w, h int) {
	p.p5Instance.Call("createCanvas", w, h)
}

func (p *P5Instance) Background(args ...interface{}) {
	p.p5Instance.Call("background", args...)
}

func (p *P5Instance) Fill(args ...interface{}) {
	p.p5Instance.Call("fill", args...)
}

func (p *P5Instance) Stroke(args ...interface{}) {
	p.p5Instance.Call("stroke", args...)
}

func (p *P5Instance) NoFill() {
	p.p5Instance.Call("noFill")
}

func (p *P5Instance) NoStroke() {
	p.p5Instance.Call("noStroke")
}

func (p *P5Instance) Ellipse(x, y, w, h float64) {
	p.p5Instance.Call("ellipse", x, y, w, h)
}

func (p *P5Instance) Rect(x, y, w, h float64) {
	p.p5Instance.Call("rect", x, y, w, h)
}

func (p *P5Instance) Line(x1, y1, x2, y2 float64) {
	p.p5Instance.Call("line", x1, y1, x2, y2)
}

func (p *P5Instance) Triangle(x1, y1, x2, y2, x3, y3 float64) {
	p.p5Instance.Call("triangle", x1, y1, x2, y2, x3, y3)
}

func (p *P5Instance) Point(x, y float64) {
	p.p5Instance.Call("point", x, y)
}

func (p *P5Instance) Arc(x, y, w, h, start, stop float64) {
	p.p5Instance.Call("arc", x, y, w, h, start, stop)
}

func (p *P5Instance) Bezier(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	p.p5Instance.Call("bezier", x1, y1, x2, y2, x3, y3, x4, y4)
}

func (p *P5Instance) QuadraticVertex(cx, cy, x, y float64) {
	p.p5Instance.Call("quadraticVertex", cx, cy, x, y)
}

func (p *P5Instance) Curve(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	p.p5Instance.Call("curve", x1, y1, x2, y2, x3, y3, x4, y4)
}

func (p *P5Instance) Text(str string, x, y float64) {
	p.p5Instance.Call("text", str, x, y)
}

func (p *P5Instance) TextSize(size float64) {
	p.p5Instance.Call("textSize", size)
}

func (p *P5Instance) Push() {
	p.p5Instance.Call("push")
}

func (p *P5Instance) Pop() {
	p.p5Instance.Call("pop")
}

func (p *P5Instance) Translate(x, y float64) {
	p.p5Instance.Call("translate", x, y)
}

func (p *P5Instance) Rotate(angle float64) {
	p.p5Instance.Call("rotate", angle)
}

func (p *P5Instance) Scale(s float64) {
	p.p5Instance.Call("scale", s)
}

func (p *P5Instance) SaveCanvas(filename, extension string) {
	p.p5Instance.Call("saveCanvas", filename, extension)
}

func (p *P5Instance) LoadImage(path string) js.Value {
	return p.p5Instance.Call("loadImage", path)
}

func (p *P5Instance) Image(img js.Value, x, y, w, h float64) {
	p.p5Instance.Call("image", img, x, y, w, h)
}

func (p *P5Instance) FrameRate(fps float64) {
	p.p5Instance.Call("frameRate", fps)
}

func (p *P5Instance) Random(min, max float64) float64 {
	return p.p5Instance.Call("random", min, max).Float()
}

func (p *P5Instance) Map(value, start1, stop1, start2, stop2 float64) float64 {
	return p.p5Instance.Call("map", value, start1, stop1, start2, stop2).Float()
}

func (p *P5Instance) BeginShape() {
	p.p5Instance.Call("beginShape")
}

func (p *P5Instance) Vertex(x, y float64) {
	p.p5Instance.Call("vertex", x, y)
}

func (p *P5Instance) EndShape(mode ...string) {
	if len(mode) > 0 {
		p.p5Instance.Call("endShape", mode[0])
	} else {
		p.p5Instance.Call("endShape")
	}
}

func (p *P5Instance) BezierVertex(cx1, cy1, cx2, cy2, x, y float64) {
	p.p5Instance.Call("bezierVertex", cx1, cy1, cx2, cy2, x, y)
}

func (p *P5Instance) CurveVertex(x, y float64) {
	p.p5Instance.Call("curveVertex", x, y)
}

func (p *P5Instance) BeginContour() {
	p.p5Instance.Call("beginContour")
}

func (p *P5Instance) EndContour() {
	p.p5Instance.Call("endContour")
}

func (p *P5Instance) Close() {
	p.p5Instance.Call("close")
}

func (p *P5Instance) TextAlign(align string) {
	p.p5Instance.Call("textAlign", align)
}

// Example usage:
//
// func main() {
// 	p5go.Execute("#container",
// 		p5go.Preload(func(p *p5go.P5Instance) {
// 			// Preload assets
// 		}),
// 		p5go.Setup(func(p *p5go.P5Instance) {
// 			p.CreateCanvas(400, 400)
// 			p.Background(255)
// 		}),
// 		p5go.Draw(func(p *p5go.P5Instance) {
// 			p.Fill(0)
// 			p.Ellipse(200, 200, 50, 50)
// 		}),
// 	)
//
// 	// Prevent the program from exiting
// 	select {}
// }
