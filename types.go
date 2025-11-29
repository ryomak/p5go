package p5go

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
