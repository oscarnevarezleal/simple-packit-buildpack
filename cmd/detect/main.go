package main

import (
	"fmt"
	"oscarnevarezleal/simple-packit-buildpack/spec"
	"github.com/cloudfoundry/packit"
)

func main() {
	fmt.Printf("Detecting -> %s", "spec.json ?")

	packit.Detect(spec.Detect())
}
