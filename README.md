# p5go
p5.js with go.wasm.

## Usage
```go
package main

import (
    "github.com/ryomak/p5go"
)

func main() {
    p5go.Execute("#container",
        p5go.Preload(func(p *p5go.P5Instance) {
            // Preload assets
        }),
        p5go.Setup(func(p *p5go.P5Instance) {
            p.CreateCanvas(400, 400)
            p.Background(255)
        }),
        p5go.Draw(func(p *p5go.P5Instance) {
            p.Fill(0)
            p.Ellipse(200, 200, 50, 50)
        }),
    )

    // Prevent the program from exiting
    select {}
}

```

## example
see [example](https://github.com/ryomak/p5go/example)

