package test

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// should be where the spec file is located
func findBaseUrl() (string, error) {
	// this code currently depends on concrete folder names of the project, so it's prone to breakage. Need to refactor ASAP!
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	splitted := strings.Split(workingDirectory, "/")
	elytraIndex := slices.Index(splitted, "backend") + 1
	basePathSplitted := append(splitted[0:elytraIndex], "docs")
	return strings.Join(basePathSplitted, "/"), nil
}

func ParseAndBuildSpec() v3.Document {

	baseDir, err := findBaseUrl()
	if err != nil {
		panic("Couldn't find the directory of spec file")
	}
	specFile, err := os.ReadFile(baseDir + "/openapi.spec.yaml")
	if err != nil {
		panic(fmt.Sprintf("Error retrieving the file: %v", err))
	}
	spec, err := libopenapi.NewDocumentWithConfiguration(specFile, &datamodel.DocumentConfiguration{
		BasePath: baseDir,
	})
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file: %v", err))
	}

	v3Model, errors := spec.BuildV3Model()

	// if anything went wrong when building the v3 model, a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported",
			len(errors)))
	}

	return v3Model.Model
}
