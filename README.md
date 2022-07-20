#### TinyGo requirements

- In addition to TinyGo, add following into system path:
https://github.com/WebAssembly/binaryen/releases

- Please note that your golang package must be ‘main’.


#### Build WASM

Bash 
```bash

GOOS=js GOARCH=wasm go build -o main.wasm main.go

tinygo build -o tiny-main.wasm -target wasi main.go

```

Powershell
```bash
$env:GOARCH="wasm"
$env:GOOS="js"
go build -o main.wasm main.go
```

Copy JS file from GOROOT/misc/wasm/wasm_exec.js

```bash
set /a ROOT=go env GOROOT
cp 
```

