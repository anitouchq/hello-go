package main

import (
	"fmt"
	"runtime"

	"github.com/anitouchq/hello-go/greeting"
)

func main() {
	fmt.Println(greeting.Message())
	fmt.Println(greeting.Details(runtime.GOOS, runtime.GOARCH))
	fmt.Println("Built and published by GitHub Actions.")
}
