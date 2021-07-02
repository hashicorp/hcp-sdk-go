package main

import (
	"encoding/json"
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
// 	...
// 	"x-go-type": {
// 		"import": {
// 			"package": "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models",
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
	log.Print("Loading shared type definitions")
	sharedDefs, err := loadSharedDefinitions()
	if err != nil {
		log.Fatalf("failed to add load shared type definitions: %v", err)
	}

	log.Print("Loading service specs")
	svcDocs, err := loadServiceSpecs("specs")
	if err != nil {
		log.Fatalf("failed to add load service specs: %v", err)
	}

	// At this point, no changes are made to the original specs.
	log.Print("Copying external type definitions")
	err = copyExternalTypes(svcDocs)
	if err != nil {
		log.Fatalf("failed to copy external type definitions: %v", err)
	}

	// After this point, the original specs are overwritten by updated specs.
	for path, svcDoc := range svcDocs {

		sp := svcDoc.Spec()

		// TODO: validate spec

		log.Printf("Adding shared type extension to spec %q", path)
		updatedSpec, err := addSharedExtension(sp, sharedDefs)
		if err != nil {
			log.Fatalf("failed to add shared extension: %v", err)
		}

		json, err := json.MarshalIndent(updatedSpec, "", "    ")
		if err != nil {
			log.Fatalf("failed to marshal json for path %q: %v", path, err)
		}

		// Overwrite existing json with the updated json spec
		log.Printf("Overwriting spec %q", path)
		ioutil.WriteFile(path, json, os.ModePerm)
	}
}

// loadSharedDefinitions parses both service and shared specs and returns a map of every shared type name
func loadSharedDefinitions() (map[string]bool, error) {
	sharedDefs := make(map[string]bool)

	// Walk the shared specs to find all internal shared type definitions.
	err := filepath.Walk("specs/hashicorp/cloud", func(path string, info os.FileInfo, err error) error {
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
		log.Fatal(fmt.Errorf("error walking the path 'specs/hashicorp/cloud': %w", err))
	}

	// Walk the service specs to find any external shared type definitions.
	err = filepath.Walk("specs", func(path string, info os.FileInfo, err error) error {
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
				if !strings.HasPrefix(def.Name, "hashicorp.cloud") {
					sharedDefs[def.Name] = true
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(fmt.Errorf("error walking the path 'specs': %w", err))
	}

	return sharedDefs, nil
}

// loadServiceSpecs parses the spec json at the given path, unmarshals the json into a Document struct,
// and saves that Document in a map for subsequent validation or mutation.
func loadServiceSpecs(svcPath string) (svcDocs map[string]*loads.Document, err error) {
	svcDocs = make(map[string]*loads.Document)

	// Walk over the spec directory
	err = filepath.Walk(svcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed accessing path %q: %w", path, err)
		}

		// These subdirectories contains shared specs, which are in a different format than the service specs.
		subDirToSkip := []string{"hashicorp", "external"}

		if info.IsDir() && (info.Name() == subDirToSkip[0] || info.Name() == subDirToSkip[1]) {
			return filepath.SkipDir
		} else if info.IsDir() {
			return nil
		}

		// Parse the found spec file into an analyzed struct that can be mutated
		svcDoc, err := loads.JSONSpec(path)
		if err != nil {
			return fmt.Errorf("failed to load spec at path %q: %w", path, err)
		}

		svcDocs[path] = svcDoc

		return nil
	})
	if err != nil {
		return nil, err
	}

	return svcDocs, nil
}

// addSharedExtensions loops over each repetition of a shared definition in a service spec and adds the type that it should reuse.
// Without adding the type reuse extension, separate copies of each shared definition are generated alongside the service-specific type definitions.
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

// copyExternalTypes traverses struct representations of the service specs and saves any imported types that have been defined in external protos,
// and copies all of those type definitions into a single swagger file under the shared specs directory.
// Those external types can then be generated as shared models in the SDK for reuse across any service SDK that requires them.
func copyExternalTypes(svcDocs map[string]*loads.Document) error {
	externalTypes := make(map[string]spec.Schema)

	// This map of documents contains go struct copies of the service spec json. The service specs themselves are never overwritten,
	// so any manipulation of the documents only affects how the external types spec gets overwritten.
	for path, svcDoc := range svcDocs {

		sp := svcDoc.Spec()

		log.Printf("Copying external type definitions at %q", path)
		for name, def := range sp.SwaggerProps.Definitions {

			// The prefix "hashicorp.cloud" indicates a type that is either defined by a cloud service proto or one of the shared hashicorp/cloud protos.
			// Any other type is deemed external, even if defined by another 'hashicorp' repo
			if !strings.HasPrefix(name, "hashicorp.cloud") {
				// The type definition may include the XGoType extension from a prior transformation of the service spec.
				// When the type definition is saved in the external spec, the extension must be removed to ensure the type can be generated as a shared model.
				delete(def.VendorExtensible.Extensions, "x-go-type")

				externalTypes[name] = def
			}
		}
	}

	log.Printf("Adding external type definitions to external.swagger.json")

	externalSpecPath := "specs/external/external.swagger.json"

	externalDoc, err := loads.JSONSpec(externalSpecPath)
	if err != nil {
		return fmt.Errorf("failed to load spec at path %q: %v", externalSpecPath, err)
	}

	updatedExternalSpect := externalDoc.Spec()
	updatedExternalSpect.SwaggerProps.Definitions = externalTypes

	json, err := json.MarshalIndent(updatedExternalSpect, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal json for path %q: %v", externalSpecPath, err)
	}

	// This helper wraps up by overwriting the existing external.swagger.json with the updated external spec,
	// which contains all the latest external type definitions copied from the service specs.
	ioutil.WriteFile(externalSpecPath, json, os.ModePerm)

	return nil
}
