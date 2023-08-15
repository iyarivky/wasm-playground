# wasm-playground
just bunch of code for learning wasm

## GO

#### For compile Go to WASM
```
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```
#### idk about this, but i think this lil js code is important

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```