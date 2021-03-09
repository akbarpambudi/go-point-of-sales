setup:
	go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen

openapi_to_web:
	oapi-codegen -generate types -o internal/library/port/web/openapi_types_gen.go -package web api/openapi/product.yaml
	oapi-codegen -generate server -o internal/library/port/web/openapi_api_gen.go -package web api/openapi/product.yaml

changelog:
	git-chglog -o CHANGELOG.md

release: changelog

.PHONY: setup openapi_to_web changelog release
