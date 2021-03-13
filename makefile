setup:
	go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen

openapi_to_web:
	oapi-codegen -generate types -o internal/library/port/web/productweb/product_openapi_types_gen.go -package productweb api/openapi/product.yaml
	oapi-codegen -generate server -o internal/library/port/web/productweb/product_openapi_api_gen.go -package productweb api/openapi/product.yaml
	oapi-codegen -generate types -o internal/common/client/productwebclient/product_openapi_types_gen.go -package productwebclient api/openapi/product.yaml
	oapi-codegen -generate client -o internal/common/client/productwebclient/product_openapi_client_gen.go -package productwebclient api/openapi/product.yaml

	oapi-codegen -generate types -o internal/library/port/web/categoryweb/category_openapi_types_gen.go -package categoryweb api/openapi/category.yaml
	oapi-codegen -generate server -o internal/library/port/web/categoryweb/category_openapi_api_gen.go -package categoryweb api/openapi/category.yaml
	oapi-codegen -generate types -o internal/common/client/categorywebclient/product_openapi_types_gen.go -package categorywebclient api/openapi/category.yaml
	oapi-codegen -generate client -o internal/common/client/categorywebclient/product_openapi_client_gen.go -package categorywebclient api/openapi/category.yaml


changelog:
	git-chglog -o CHANGELOG.md

release_changelog:
	git add CHANGELOG.md
	git commit -m "chore: release new changelog"
	git push -u origin master

release: changelog release_changelog

.PHONY: setup openapi_to_web changelog release
