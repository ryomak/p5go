package main

import (
	"fmt"

	"github.com/ryomak/p5go"
)

func main() {
	err := p5go.Run("main",
		p5go.Setup(func(p *p5go.Canvas) {
			p.CreateCanvas(400, 400)
			p.Background(128, 200, 128)
		}),
		p5go.Draw(func(p *p5go.Canvas) {
			p.Fill(0)
			p.Ellipse(200, 200, 50, 50)
			p.Text("Hello, p5go", 50, 20)
		}),
	)
	fmt.Println(err)
	select {}
}
