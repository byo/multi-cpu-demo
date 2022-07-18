package main

import (
	"fmt"
	"runtime"
)

func getOsArch() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}

func main() {
	fmt.Printf(
		"Hello world from %s!\n",
		getOsArch(),
	)
}
