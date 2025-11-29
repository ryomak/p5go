package p5go

import "math"

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

// CaptureKind is a type that represents the kind of capture.
type CaptureKind string

const (
	CaptureKindVIDEO CaptureKind = "VIDEO"
	CaptureKindIMAGE CaptureKind = "IMAGE"
)
