package main

import (
	"os"
	"tests/math/svg"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
