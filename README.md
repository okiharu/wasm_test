# wasm_test

A program to find 'SUPINF' from a randomly generated string.

I made this to verify the difference in processing speed between JavaScript and WebAssembly. 
The results are only a guide.

## Getting Started
At /etc/wasm_test/findSupinf

```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ go run ../server.go .
```
Access to
http://localhost:8080/
