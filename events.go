package p5go

import "syscall/js"

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
