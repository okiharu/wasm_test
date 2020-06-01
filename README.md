# wasm_test

A program to find 'SUPINF' from a randomly generated string.

I made this to verify the difference in processing speed between JavaScript and WebAssembly. 
The results are only a guide.

## Getting Started
Go to the findSupinf directory:

```
$ cd $HOME/src/github.com/okiharu/wasm_test/findSupinf
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ go run ../server.go .
```
Access to
http://localhost:8080/
