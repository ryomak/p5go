
## Usage
### hosting
```bash
go build -o hosting ./server
./hosting
```

### simple
```bash
GOARCH=wasm GOOS=js go build -o main.wasm ./simple
``` 

