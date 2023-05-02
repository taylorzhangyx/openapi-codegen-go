// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
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

	"H4sIAAAAAAAC/+RWT4skNRT/KuEpeCmr2plFljq5uz1Cg4yDrhfXOWQqr7silT+TvJqecigQvCkoHrzt",
	"ybOwXgQRwS/j7PoxJEnNzHZX9eyKIsqeqpK89/LL+/3eSy6gMsoajZo8lBfgqxoVj78HzhkXfqwzFh1J",
	"jNOVERi+An3lpCVpNJTJmMW1DJbGKU5QgtS0vwcZUGcxDXGFDvoMFHrPVzsDXS1fu3pyUq+g7zNweNpK",
	"hwLKRzBseGV+3GdwiOsjpDFuzdXEdodcITNLRjUyizTeMAPiq7Hfw87e7rcFNO4e4A3YeNO8v4Ty0QW8",
	"7nAJJbxW3BBRDCwUw1n6bPswUowhfaTlaYtMik1cz5Px9p0JMraQSgHH/XEfpqVemkS5Jl5F3Ki4bKAE",
	"4l1j3Gfd+Tu1oTCXV0ZBNmQZHsZl9nHN9QoyaF3wqYmsL4vi2jevua+1EZgLPCtCrreSHA3f8OwIyZNx",
	"yPCcK9sgI8MEKqM9OU7IjEXNrWT7+QwyaGSF2uMN6XDP8qpGtheXN8Gs1+ucx+XcuFUx+PrivcWDg8MP",
	"D97cy2d5TaqJSpDU4BQuyOAMnU+o38pn+SyYD6ighP04lYHlVEf+Coup3lZJDpvn/gCpddoz3jSBRM+W",
	"zqhIqe88YfjlFMetR8dq7hmvKvSekflEQ9zZ8RBsIaCEd6UWAWkE4LhCQuej9ja3Jb4KEdhSNoSOnXQQ",
	"JAAlnLbouhtygx1kQ6uIaiRU8TTj2kkT3DnehbGnLmYwiDLKehOB4udStYrpVp2gCzp26NuGIiwXs7ID",
	"UyOVpA1QL2xC/XEQvrcm8B089mazK7WjTlVqbSOrmMniUx8gXkwd+7YSTvW7lYh+JHWLxK7ApEJY8rah",
	"v4TnNhipmU9s3Go8t1gRCoaDTQa+VYq7bkKKAZs1fkK0DxxyQs8407gOtkzqpNlQITmbtwl7MHEYApo1",
	"ipFY74mgVUg9CT3dN6L7x7Jw1U/HaThCChrjQoTPNWx4vjeSa7H/m5p5oVT+P9IYER7XY2crLqTok0Qa",
	"pIlrN80HXy/1qol3FTvhHgUzSTWLOfNtONOERubRO8nk1o62mIceYhO3A5ahf4RWfNM+pBgxvauXTN+h",
	"415yZ3zqACShEP8lIufXZEQWOraYB3i3X02bjF3zuJjvun7ud3Ht5flaIlX1v0bXK1vGW4wm9qMJurNp",
	"mp4+/uHy8Y+Xv35++c23z75+cvn9FxvPqrIoGlPxpjaeyruzu7OJm/7ZL7/9/vOXO0PcPBPt8MQKT8Sc",
	"UFeoKb41d8V8+tNXfzz57uViEnraDHrc/xkAAP//4DPsJBINAAA=",
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