# use oapi-codegen to generate server and struct code
# https://github.com/deepmap/oapi-codegen
oapi-codegen -package api -generate types taylorpetstore.yaml > autogen/taylorpetstore.types.gen.go
oapi-codegen -package api -generate spec taylorpetstore.yaml > autogen/taylorpetstore.spec.gen.go
oapi-codegen -package api -generate gin taylorpetstore.yaml > autogen/taylorpetstore.ginserver.gen.go
oapi-codegen -package api -generate client taylorpetstore.yaml > autogen/taylorpetstore.client.gen.go

# use widdershins to generate markdown documentation
# https://github.com/Mermade/widdershins
widdershins --search true --language_tabs 'go:Go' -u templates -c --summary taylorpetstore.yaml -o taylorpetstore.md
