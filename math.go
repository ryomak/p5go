package p5go

// Random returns a random number between the specified min and max values.
func (c *Canvas) Random(min, max float64) float64 {
	return c.p5Instance.Call("random", min, max).Float()
}

// Map maps a value from one range to another.
func (c *Canvas) Map(value, start1, stop1, start2, stop2 float64) float64 {
	return c.p5Instance.Call("map", value, start1, stop1, start2, stop2).Float()
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
