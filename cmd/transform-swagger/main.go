package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/iancoleman/strcase"
)

// XGoType contains the fields of the go-swagger extension x-go-type
// In its final form, the extension looks like this on the generated type definition:
// "hashicorp.cloud.common.PaginationRequest": {
// 	"properties": { ... }
//	...
//	"x-go-type": {
//		"import": {
//		"package": "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
// 			"alias": "cloud"
// 		},
// 		"type": "HashicorpCloudCommonPaginationRequest"
// 	}
// }
// ref: https://goswagger.io/use/models/schemas.html#types-reusability
type XGoType struct {
	Import Import `json:"import"`
	Type   string `json:"type"`
}

// Import contains the import info used in the go-swagger extension x-go-type
type Import struct {
	Package string `json:"package"`
	Alias   string `json:"alias"`
}

func main() {

	// These flags collect:
	// - the path of the service spec, i.e. specs/consul-foo-service/preview/2021-07-09/hcp.swagger.json
	// - the path of the shared spec directory, should always end with 'specs/cloud-shared'
	// - the path of the external spec, which should always end with 'specs/external/external.swagger.json'

	svcPathPtr := flag.String("service", "", "the path of the service specs")
	sharedPathPtr := flag.String("shared", "", "the path of the shared specs")
	externalSpecPathPtr := flag.String("external", "", "the path of the external types spec")

	flag.Parse()

	svcPath := *svcPathPtr
	sharedPath := *sharedPathPtr
	externalSpecPath := *externalSpecPathPtr

	log.Print("Creating map of shared type definitions")
	sharedDefs, err := loadSharedDefinitions(sharedPath, svcPath)
	if err != nil {
		log.Fatalf("failed to add load shared type definitions: %v", err)
	}

	// No changes are made to the original spec at this step.
	log.Printf("Copying external type definitions from spec %q", svcPath)
	err = copyExternalTypes(externalSpecPath, svcPath)
	if err != nil {
		log.Fatalf("failed to copy external type definitions: %v", err)
	}

	svcDoc, err := loads.JSONSpec(svcPath)
	if err != nil {
		log.Fatalf("failed to load spec at path %q: %v", svcPath, err)
	}

	sp := svcDoc.Spec()

	log.Printf("Adding shared type extension to spec %q", svcPath)
	updatedSpec, err := addSharedExtension(sp, sharedDefs)
	if err != nil {
		log.Fatalf("failed to add shared extension: %v", err)
	}

	json, err := json.MarshalIndent(updatedSpec, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal json for path %q: %v", svcPath, err)
	}

	// Overwrite original spec with the transformed spec.
	log.Printf("Overwriting spec %q", svcPath)
	err = ioutil.WriteFile(svcPath, json, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to overwrite spec at path %q: %v", svcPath, err)
	}
}

// loadSharedDefinitions parses both the service spec and shared specs and returns a map of every shared type name
func loadSharedDefinitions(sharedPath, svcPath string) (map[string]bool, error) {
	sharedDefs := make(map[string]bool)

	// Walk the shared specs to find all internal shared type definitions.
	err := filepath.Walk(sharedPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed accessing path %q: %w", path, err)
		}

		if info.IsDir() {
			return nil
		}

		doc, err := loads.JSONSpec(path)
		if err != nil {
			return fmt.Errorf("failed to load spec at path %q: %w", path, err)
		}

		for _, def := range doc.Analyzer.AllDefinitions() {
			if def.TopLevel {
				sharedDefs[def.Name] = true
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(fmt.Errorf("error walking the path %q: %w", sharedPath, err))
	}

	doc, err := loads.JSONSpec(svcPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load spec at path %q: %w", svcPath, err)
	}

	for _, def := range doc.Analyzer.AllDefinitions() {
		if def.TopLevel {
			if !strings.HasPrefix(def.Name, "hashicorp.cloud") {
				sharedDefs[def.Name] = true
			}
		}
	}

	return sharedDefs, nil
}

// addSharedExtensions loops over each shared type definition in a service spec and adds the type that it should reuse.
// Without adding the type reuse extension, separate copies of each shared type definition are generated alongside the service-specific type definitions.
func addSharedExtension(apiSpec *spec.Swagger, sharedDefs map[string]bool) (*spec.Swagger, error) {
	for sharedDefName := range sharedDefs {

		var def spec.Schema
		var ok bool
		if def, ok = apiSpec.SwaggerProps.Definitions[sharedDefName]; !ok {
			continue
		}

		// The shared definition name gets transformed into a camelcased Go type.
		// example: hashicorp.cloud.common.PaginationRequest -> HashicorpCloudCommonPaginationRequest
		genTypeName := strcase.ToCamel(sharedDefName)

		// This struct contains all the data, like which package to import, needed for the SDK client generator to ensure the service reuses
		// the generated shared type, rather than a service-specific duplicate of that shared type.
		ext := XGoType{
			Import: Import{
				Package: "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
				Alias:   "cloud",
			},
			Type: genTypeName,
		}

		def.VendorExtensible.AddExtension("x-go-type", ext)

		// The value of 'def' is a copy of the definition's schema, so here the original must be overwritten by the copy with the extension.
		apiSpec.SwaggerProps.Definitions[sharedDefName] = def
	}

	return apiSpec, nil
}

// copyExternalTypes reads a given service spec and finds any imported types that have been defined in external protos,
// then copies those type definitions into an 'external' spec file. Those external types can then be generated as shared models
// in the SDK for reuse across any service SDK that requires them.
func copyExternalTypes(externalSpecPath, svcPath string) error {
	externalTypes := make(map[string]spec.Schema)

	// Create a map of existing external definitions.
	externalDoc, err := loads.JSONSpec(externalSpecPath)
	if err != nil {
		return fmt.Errorf("failed to load spec at path %q: %w", externalSpecPath, err)
	}

	for _, def := range externalDoc.Analyzer.AllDefinitions() {
		if def.TopLevel {
			externalTypes[def.Name] = *def.Schema
		}
	}

	// Load the service spec.
	svcDoc, err := loads.JSONSpec(svcPath)
	if err != nil {
		return fmt.Errorf("failed to load spec at path %q: %w", svcPath, err)
	}

	sp := svcDoc.Spec()

	// Read the service spec to find any external type definitions.
	log.Printf("Copying external type definitions from %q", svcPath)
	for name, def := range sp.SwaggerProps.Definitions {

		// The prefix "hashicorp.cloud" indicates a type that is either defined by a cloud service proto or one of the shared hashicorp/cloud protos.
		// Any other type is deemed external, even if defined by another 'hashicorp' repo.
		if !strings.HasPrefix(name, "hashicorp.cloud") {
			// The type definition may include the XGoType extension from a prior transformation of the service spec.
			// When the type definition is saved in the external spec, the extension must be removed to ensure the type can be generated as a shared model.
			delete(def.VendorExtensible.Extensions, "x-go-type")

			externalTypes[name] = def
		}
	}

	log.Printf("Saving updated external type definitions to %q", externalSpecPath)
	updatedExternalSpec := externalDoc.Spec()
	updatedExternalSpec.SwaggerProps.Definitions = externalTypes

	json, err := json.MarshalIndent(updatedExternalSpec, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal json for path %q: %v", externalSpecPath, err)
	}

	// Overwrite existing external.swagger.json with the updated external definitions.
	err = ioutil.WriteFile(externalSpecPath, json, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to overwrite spec at path %q: %v", svcPath, err)
	}

	return nil
}
