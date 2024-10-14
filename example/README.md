
## Usage
### hosting
```bash
go build -o hosting ./server
./hosting
```

### build simple wasm
```bash
GOARCH=wasm GOOS=js go build -o main.wasm ./simple
``` 


### check browser
```bash
open http://localhost:3000
```



