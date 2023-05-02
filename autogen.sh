# use oapi-codegen to generate server and struct code
# https://github.com/deepmap/oapi-codegen
oapi-codegen -package api -generate types featureplatform.yaml > autogen/featureplatform.types.gen.go
oapi-codegen -package api -generate spec featureplatform.yaml > autogen/featureplatform.spec.gen.go
oapi-codegen -package api -generate gin featureplatform.yaml > autogen/featureplatform.ginserver.gen.go
oapi-codegen -package api -generate client featureplatform.yaml > autogen/featureplatform.client.gen.go

# use widdershins to generate markdown documentation
# https://github.com/Mermade/widdershins
widdershins --search true --language_tabs 'go:Go' -u templates -c --summary featureplatform.yaml -o featureplatform.md
