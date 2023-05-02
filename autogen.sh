# use oapi-codegen to generate server and struct code
# https://github.com/deepmap/oapi-codegen
oapi-codegen -package api -generate types taylorpetstore.yaml > api/taylorpetstore.types.gen.go
oapi-codegen -package api -generate spec taylorpetstore.yaml > api/taylorpetstore.spec.gen.go
oapi-codegen -package api -generate gin taylorpetstore.yaml > api/taylorpetstore.ginserver.gen.go
oapi-codegen -package api -generate client taylorpetstore.yaml > api/taylorpetstore.client.gen.go

# use widdershins to generate markdown documentation
# https://github.com/Mermade/widdershins
widdershins --search true --language_tabs 'go:Go' -u templates -c --summary taylorpetstore.yaml -o taylorpetstore.md
