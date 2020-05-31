# wasm_test

A program to find a SUPINF from a randomly generated string.

## Getting Started
At /etc/wasm_test/findSupinf

```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ go run ../server.go .
```
Access to
http://localhost:8080/