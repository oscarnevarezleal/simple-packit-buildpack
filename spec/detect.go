package spec

import (
	"encoding/json"
	"github.com/cloudfoundry/packit"
	"os"
	"fmt"
	"path/filepath"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {

		fmt.Printf("Detecting -> %s", "spec.json ?")

		// The DetectContext includes a WorkingDir field that specifies the
		// location of the application source code. This field can be combined with
		// other paths to find and inspect files included in the application source
		// code that is provided to the buildpack.
		file, err := os.Open(filepath.Join(context.WorkingDir, "spec.json"))

		if err != nil {
			fmt.Printf("Spec file '%s' was not found", filepath.Join(context.WorkingDir, "spec.json"))
			return packit.DetectResult{}, fmt.Errorf("spec file not found")
		}

		// The spec.json file includes a declaration of what versions of php
		// are acceptable. For example:
		//   {
		//     "php": {
		//       "version": ">=0.10.3 <0.12"
		//     }
		//   }
		var config struct {
			PhpConfig struct {
				Version string `json:"version"`
			} `json:"php"`
		}

		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			fmt.Printf("	--> An error ocurred while parsing spec file: '%s'", err)
			return packit.DetectResult{}, fmt.Errorf("invalid spec file")
		}

		// Once the spec.json file has been parsed, the detect phase can return
		// a result that indicates the provision of xxxxx and the requirement of
		// xxxxx. As can be seen below, the BuildPlanRequirement may also include
		// optional metadata information to such as the source of the version
		// information for a given requirement.

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Requires: []packit.BuildPlanRequirement{
					{
						Name:    "php-dist",
						Version: config.PhpConfig.Version,
						Metadata: map[string]string{
							"version-source": "spec.json",
						},
					},
				},
			},
		}, nil
	}
}
