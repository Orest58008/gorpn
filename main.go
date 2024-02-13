package main

import (
	"log"
	"os"
	"strconv"
	"fmt"
	"math"
)

func main() {
	// Initializing stack
	var stack []float64

	// Binary operations
	binaryOps := map[string] func(float64, float64) float64 {
		"+":  func(x, y float64) float64 { return x + y },
		"-":  func(x, y float64) float64 { return x - y },
		"x":  func(x, y float64) float64 { return x * y },
		"/":  func(x, y float64) float64 { return x / y },
		"xx": func(x, y float64) float64 { return math.Pow(x, y) },
	}

	// Processing args
	for i, e := range os.Args {
		if i == 0 { continue // Don't process program's name
		} else if op, ok := binaryOps[e]; ok {
			if len(stack) < 2 { log.Fatal("Provide more arguments!") }

			// I do love Golang
			y, stack := stack[len(stack)-1], stack[:len(stack)-1]
			x, stack := stack[len(stack)-1], stack[:len(stack)-1]

			stack = append(stack, op(x, y))
		} else if num, err := strconv.ParseFloat(e, 64); err == nil {
			stack = append(stack, num)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("%G\n",stack[0])
}
