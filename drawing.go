package p5go

import "syscall/js"

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

// Circle draws a circle on the canvas.
func (c *Canvas) Circle(x, y, d float64) {
	c.p5Instance.Call("circle", x, y, d)
}

// Square draws a square on the canvas.
func (c *Canvas) Square(x, y, s float64) {
	c.p5Instance.Call("square", x, y, s)
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
func (c *Canvas) Vertex(x, y float64, z ...float64) {
	if len(z) > 0 {
		c.p5Instance.Call("vertex", x, y, z[0])
	} else {
		c.p5Instance.Call("vertex", x, y)
	}
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

// StrokeWeight sets the weight of the stroke.
func (c *Canvas) StrokeWeight(weight float64) {
	c.p5Instance.Call("strokeWeight", weight)
}

// StrokeCap sets the style of the stroke cap.
func (c *Canvas) StrokeCap(cap ShapeType) {
	c.p5Instance.Call("strokeCap", string(cap))
}

// StrokeJoin sets the style of the joints which connect line segments.
func (c *Canvas) StrokeJoin(join ShapeType) {
	c.p5Instance.Call("strokeJoin", string(join))
}

// Erase enables the eraser tool.
func (c *Canvas) Erase(opt ...any) {
	c.p5Instance.Call("erase", opt...)
}

// NoErase disables the eraser tool.
func (c *Canvas) NoErase() {
	c.p5Instance.Call("noErase")
}

// EllipseMode sets the location from which ellipses are drawn.
func (c *Canvas) EllipseMode(mode DrawingMode) {
	c.p5Instance.Call("ellipseMode", string(mode))
}

// RectMode sets the location from which rectangles are drawn.
func (c *Canvas) RectMode(mode DrawingMode) {
	c.p5Instance.Call("rectMode", string(mode))
}

// Smooth draws all geometry with smooth (anti-aliased) edges.
func (c *Canvas) Smooth() {
	c.p5Instance.Call("smooth")
}

// NoSmooth draws all geometry with jagged (aliased) edges.
func (c *Canvas) NoSmooth() {
	c.p5Instance.Call("noSmooth")
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

// Box draws a 3D box. width / height / depth are optional (defaults to 50).
// Requires WEBGL renderer.
func (c *Canvas) Box(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("box", opts...)
	} else {
		c.p5Instance.Call("box")
	}
}

// Sphere draws a 3D sphere. radius is optional (default 50).
// detailX and detailY control segment count. Requires WEBGL renderer.
func (c *Canvas) Sphere(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("sphere", opts...)
	} else {
		c.p5Instance.Call("sphere")
	}
}

// Ellipsoid draws a 3D ellipsoid with the given radii along x/y/z.
// detailX and detailY are optional. Requires WEBGL renderer.
func (c *Canvas) Ellipsoid(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("ellipsoid", opts...)
	} else {
		c.p5Instance.Call("ellipsoid")
	}
}

// Cylinder draws a 3D cylinder. radius / height / detailX / detailY / bottomCap / topCap are optional.
// Requires WEBGL renderer.
func (c *Canvas) Cylinder(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("cylinder", opts...)
	} else {
		c.p5Instance.Call("cylinder")
	}
}

// Cone draws a 3D cone. radius / height / detailX / detailY / cap are optional.
// Requires WEBGL renderer.
func (c *Canvas) Cone(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("cone", opts...)
	} else {
		c.p5Instance.Call("cone")
	}
}

// Plane draws a 3D plane on the XY axes. width / height / detailX / detailY are optional.
// Requires WEBGL renderer.
func (c *Canvas) Plane(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("plane", opts...)
	} else {
		c.p5Instance.Call("plane")
	}
}

// Torus draws a 3D torus. radius / tubeRadius / detailX / detailY are optional.
// Requires WEBGL renderer.
func (c *Canvas) Torus(opts ...any) {
	if len(opts) > 0 {
		c.p5Instance.Call("torus", opts...)
	} else {
		c.p5Instance.Call("torus")
	}
}
