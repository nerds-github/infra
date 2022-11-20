// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}