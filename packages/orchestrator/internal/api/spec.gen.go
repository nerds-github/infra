// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
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

	"H4sIAAAAAAAC/7RWTW+jSBD9K6h2j8g4ye6F22YT7Vob7UTj0VwiHypQhh433aS7SGJZ/PdRN+bDBqwc",
	"MieMq7vq1Xv1wQESXZRakWIL8QFKNFgQk/FvFlX6rN9Xd+5FKIihRM4hBIUFQTywh2DopRKGUojZVBSC",
	"TXIq0F3kfekPsxEqg7qu3WFbamXJR/lzuXSPRCsmxe4nlqUUCbLQKvphtfJYOn+/G9pCDL9FPfSosdro",
	"3hhtmhgp2cSI0jmBGNZkXskEdLS3+DyA5pLL3uiSDIsGV6JTcs9TR/5w4G0hbLUpkCEGofjmGsI2V6GY",
	"MjJQh1CQtZjNOeqvdPQMqXyCY6DWy6YO4X96WzfET2CWghQ3gp05Dk/lvBx2qGzn0wWfjYxSNGyOwj5X",
	"QqYzkC7iJfU6b7F3wu4mjVthKDGY7Mh8J2NFUz+jY3mV0SNmNAT9rLUkVM4slGVUCc0g2JFRJC/5L/B9",
	"dfTxQCrj3J36UMEwpsg4LWGJagYRExZzJoMziZzL3rjvb7Qi9CqGR6VPGDrnY1KEgWwd2imaBhQMVdr4",
	"rhZqq30egqVL5P76NvjrcQUhvLZawHJxtVi6vHVJCksBMdwsloslhH5+ecGjnFA2mmTE4+7815uDJKfE",
	"wXWV7ufRKoUY/iFu7HA2yq6bUXbq6iu9VGQ5eEMb2CpJyNptJd2QqkOIjp3W3J+E8iAsByhlYCqlhMqC",
	"/soEsPXAOIXtw2NWMBXHzpZfthA/XZ687WCo3ZQ4Vhkag/vJadzRIPeBIa6MonQmxzpsN8RU+C7FyB1y",
	"kRgzOxhhvnJCKLWdoPZvQ8gUYBsu2BpdBJxTwFSUEplGFD9qe8axV/dWp/tP22Idl6f96RZrPVL16tPC",
	"DhbLhGbfcupocqWceO7SQUnL/aeIddIV0aFbRXWjnySe2Kb/CSl7HUeq3flrnW7rwXobfvLM1Hh/JOoX",
	"Y70ZSfHHGNc5azsh5S8hrf//cPZ55npoU/8MAAD//+nkwynqCQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
