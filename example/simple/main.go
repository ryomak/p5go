package main

import (
	"fmt"

	"github.com/ryomak/p5go"
)

func main() {
	err := p5go.Run("main",
		p5go.Setup(func(c *p5go.Canvas) {
			c.CreateCanvas(400, 400)
			c.Background(128, 200, 128)
		}),
		p5go.Draw(func(c *p5go.Canvas) {
			c.Fill(0)
			c.Ellipse(200, 200, 50, 50)
			c.Text("Hello, p5go", 50, 20)
		}),
	)
	fmt.Println(err)
	select {}
}
