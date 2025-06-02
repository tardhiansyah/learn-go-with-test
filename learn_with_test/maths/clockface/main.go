package main

import (
	"os"
	"time"

	clockface "example.com/learnGoWithTest/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
