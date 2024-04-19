package main

import (
	"os"
	"tests/math"
	"time"
)

func main() {
	t := time.Now()
	math.SVGWriter(os.Stdout, t)
}
