package p5go

// Lighting and material APIs for the WEBGL renderer.
// All of these are no-ops in 2D mode — switch to WEBGL via CreateCanvas(w, h, WEBGL)
// before calling them.

// AmbientLight sets the ambient light color.
// Accepts the same forms as Fill: (gray), (gray, alpha), (r, g, b),
// (r, g, b, a), or a CSS color string / Color struct.
func (c *Canvas) AmbientLight(args ...any) {
	c.p5Instance.Call("ambientLight", args...)
}

// DirectionalLight sets a directional light with color (r,g,b) and direction (x,y,z).
// Direction values are in the range -1..1.
func (c *Canvas) DirectionalLight(r, g, b, x, y, z float64) {
	c.p5Instance.Call("directionalLight", r, g, b, x, y, z)
}

// PointLight sets a point light at (x, y, z) emitting color (r, g, b).
func (c *Canvas) PointLight(r, g, b, x, y, z float64) {
	c.p5Instance.Call("pointLight", r, g, b, x, y, z)
}

// SpotLight sets a spotlight at (x, y, z) emitting color (r, g, b)
// pointing in direction (rx, ry, rz). angle and concentration are optional.
func (c *Canvas) SpotLight(r, g, b, x, y, z, rx, ry, rz float64, opts ...float64) {
	args := []any{r, g, b, x, y, z, rx, ry, rz}
	for _, v := range opts {
		args = append(args, v)
	}
	c.p5Instance.Call("spotLight", args...)
}

// Lights enables a default ambient + directional light setup. Convenient for quick 3D scenes.
func (c *Canvas) Lights() {
	c.p5Instance.Call("lights")
}

// NoLights disables all active lights for subsequent geometry.
func (c *Canvas) NoLights() {
	c.p5Instance.Call("noLights")
}

// LightFalloff sets attenuation factors (constant, linear, quadratic) for point / spot lights.
func (c *Canvas) LightFalloff(constant, linear, quadratic float64) {
	c.p5Instance.Call("lightFalloff", constant, linear, quadratic)
}

// SpecularColor sets the color of highlights for subsequent specular materials.
func (c *Canvas) SpecularColor(args ...any) {
	c.p5Instance.Call("specularColor", args...)
}

// AmbientMaterial sets the material color used for ambient lighting.
func (c *Canvas) AmbientMaterial(args ...any) {
	c.p5Instance.Call("ambientMaterial", args...)
}

// SpecularMaterial sets the material highlight color for specular reflections.
func (c *Canvas) SpecularMaterial(args ...any) {
	c.p5Instance.Call("specularMaterial", args...)
}

// EmissiveMaterial sets a self-illuminating material color (not affected by lights).
func (c *Canvas) EmissiveMaterial(args ...any) {
	c.p5Instance.Call("emissiveMaterial", args...)
}

// NormalMaterial colors geometry by surface normal — no lights required.
// Useful for debugging and quick 3D visualizations.
func (c *Canvas) NormalMaterial() {
	c.p5Instance.Call("normalMaterial")
}

// Shininess sets the highlight sharpness for the current specular material.
func (c *Canvas) Shininess(s float64) {
	c.p5Instance.Call("shininess", s)
}

// Camera positions the camera at (x,y,z) looking at (cx,cy,cz) with up vector (ux,uy,uz).
func (c *Canvas) Camera(x, y, z, cx, cy, cz, ux, uy, uz float64) {
	c.p5Instance.Call("camera", x, y, z, cx, cy, cz, ux, uy, uz)
}

// Perspective configures a perspective projection.
// fovy in radians, aspect = width/height, near/far are clip distances.
func (c *Canvas) Perspective(fovy, aspect, near, far float64) {
	c.p5Instance.Call("perspective", fovy, aspect, near, far)
}

// Ortho configures an orthographic projection.
func (c *Canvas) Ortho(left, right, bottom, top, near, far float64) {
	c.p5Instance.Call("ortho", left, right, bottom, top, near, far)
}
