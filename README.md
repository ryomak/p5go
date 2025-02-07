# p5go
p5go provides a bridge between Go and p5.js.
inspired by https://github.com/ongaeshi/p5rb

* Pull requests are welcome as there are missing features.

## Usage

```go
package main

import (
	"github.com/ryomak/p5go"
)

func main() {
	p5go.Run("#container",
		p5go.Setup(func(c *p5go.Canvas) {
			c.CreateCanvas(400, 400)
			c.Background(255)
		}),
		p5go.Draw(func(c *p5go.Canvas) {
			c.Fill(0)
			c.Ellipse(200, 200, 50, 50)
		}),
	)

	// Prevent the program from exiting
	select {}
}

```

## example
see [example](https://github.com/ryomak/p5go/tree/main/example)

