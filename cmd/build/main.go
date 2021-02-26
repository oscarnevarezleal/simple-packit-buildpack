package main

import (
	"oscarnevarezleal/simple-packit-buildpack/spec"
	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Build(spec.Build())
}
