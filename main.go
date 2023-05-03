// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"taylorzhangyx.github.com/taylorpetstore/api"
	"taylorzhangyx.github.com/taylorpetstore/biz"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
)

func NewGinPetServer(petStore *biz.PetStore, host string, port int) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic gin router
	r := gin.Default()

	// Enabling CORS, for more info see: https://github.com/gin-contrib/cors
	r.Use(cors.Default())

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(r, petStore)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}
	return s
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	var host = flag.String("host", "localhost", "Host for test HTTP server")

	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	petStore := biz.NewPetStore()
	s := NewGinPetServer(petStore, *host, *port)

	fmt.Printf("Serving on %s:%d\n", *host, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
