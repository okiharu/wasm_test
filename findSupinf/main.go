package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

var clickCount float64
var avgTime float64
var totalTime float64

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("findByGo", js.FuncOf(findByGo))
	js.Global().Set("create", js.FuncOf(create))
}

func findByGo(this js.Value, i []js.Value) interface{} {
	//start
	start := time.Now()

	str := js.Global().Get("document").Call("getElementById", "randomStr").Get("value").String()
	countSup := 0
	for i := 0; i < len(str)-5; i++ {
		if str[i:i+6] == "SUPINF" {
			countSup++
		}
	}
	js.Global().Get("document").Call("getElementById", "result").Set("value", countSup)

	//end
	end := time.Now()
	time := (end.Sub(start)).Seconds() * 1000
	fmt.Printf("%f\n", time)
	js.Global().Get("document").Call("getElementById", "spdGo").Set("value", time)

	culcAvgTime(time)
	return nil
}

func culcAvgTime(time float64) {
	clickCount++
	totalTime += time
	avgTime = totalTime / clickCount
	js.Global().Get("document").Call("getElementById", "avgGo").Set("value", avgTime)
}

func create(this js.Value, i []js.Value) interface{} {
	supinf := []rune("SUPINF")
	num, _ := strconv.Atoi(js.Global().Get("document").Call("getElementById", "num").Get("value").String())
	rndRune := make([]rune, num)
	for i := 0; i < num; i++ {
		rndRune[i] = supinf[rand.Intn(6)]
	}
	js.Global().Get("document").Call("getElementById", "randomStr").Set("value", string(rndRune))

	initVals()
	return nil
}

func initVals() {
	clickCount = 0
	avgTime = 0
	totalTime = 0
}
