
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
![スクリーンショット 2024-10-14 10 58 19](https://github.com/user-attachments/assets/055517ac-0b7b-415a-a560-233a03edfa34)


