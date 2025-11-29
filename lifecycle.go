package p5go

import (
	"errors"
	"fmt"
	"syscall/js"
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

	// P5.jsがロードされていない場合は追加
	if global.Get("p5").IsUndefined() {
		doc := global.Get("document")
		script := doc.Call("createElement", "script")
		script.Set("src", "https://cdn.jsdelivr.net/npm/p5@1.11.2/lib/p5.min.js")
		doc.Get("head").Call("appendChild", script)

		ch := make(chan struct{})
		script.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			close(ch)
			return nil
		}))
		<-ch
	}

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

// Func is a type that represents a function that takes a Canvas pointer as an argument.
type Func func(c *Canvas)

// Preload sets the preload handler for the canvas.
func Preload(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["preload"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// Setup sets the setup handler for the canvas.
func Setup(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["setup"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}

// Draw sets the draw handler for the canvas.
func Draw(handler func(c *Canvas)) Func {
	return func(c *Canvas) {
		c.funcHandlers["draw"] = js.FuncOf(func(value js.Value, args []js.Value) any {
			handler(c)
			return nil
		})
	}
}
