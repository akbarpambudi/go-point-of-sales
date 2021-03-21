<a name="unreleased"></a>
## [Unreleased]

### Features
- Get All Product API


<a name="v1.5.0"></a>
## [v1.5.0] - 2021-03-16
### Features
- `Get Category by ID` rest API handler


<a name="v1.4.3"></a>
## [v1.4.3] - 2021-03-14

<a name="1.4.2"></a>
## [1.4.2] - 2021-03-14
### Bug Fixes
- category creation API should response with status 201 - Created if there is no error happened

### Code Refactoring
- wrap error returned from NewCategory factory function with IllegalCreationInputError in order to standardize the error
- refactor product/errhelper.go name to error_helper.go in order to follow internal naming conventions

### Features
- add utility type for constructing json in API Testing
- create error key constant for category creation
- create `IllegalCreationInputError` on category domain entity pkg in order to standardize validation error


<a name="v1.4.1"></a>
## [v1.4.1] - 2021-03-14

<a name="v1.4.0"></a>
## [v1.4.0] - 2021-03-13
### Features
- generate http client for product and category REST API
- refactor product and variant creation error on library domain
- create POSHTTPError to map error to http error
- create multi error on errors package


<a name="v1.3.0"></a>
## [v1.3.0] - 2021-03-11
### Code Refactoring
- add open api to web code generator for category
- add CreateCategory handler on Application Command list
- move product web handler into new package

### Features
- register product web handler api on library service
- generate web api handler for category base on the api spec
- create api spec for category
- add adapter implementation for category repository


<a name="v1.2.0"></a>
## [v1.2.0] - 2021-03-11
### Code Refactoring
- make category repository to return and accept pointer of category entity
- turn category's error into POSError

### Features
- create 'CreateCategory' command handler


<a name="v1.1.0"></a>
## [v1.1.0] - 2021-03-10
### Code Refactoring
- refactor product handler base url to /api/product

### Features
- add variant on product get by id handler
- add variant on product read model


<a name="v0.0.1"></a>
## v0.0.1 - 2021-03-10
### Code Refactoring
- remove api and handler file on product web handler package
- refactor variants type on create product command
- add endpoint identifier on product open api spec
- refactor product entity error to errors.POSError
- moving category error
- moving category to separate package

### Features
- add echo, sqlite, and oapi-codegen library
- wiring up library into runnable web service
- create helper function for constructing POSError on errors pkg
- create pointer to value utils on ptrval package
- create ent implementation read model projector for product read model
- add http handler interface and types for product base on open api spec,create query handler for get product by id
- create web handler for product base on open api spec
- create error to http error response mapper on httphelper package
- add common error for POS application
- create port layer for 'product'
- create adapter layer repository implementation for 'category' and 'product' entity
- create domain repository for `category` entity
- create domain repository for `product` entity
- add `create product` command handler
- create category entity on product domain
- create product entity on product domain


[Unreleased]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.5.0...HEAD
[v1.5.0]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.4.3...v1.5.0
[v1.4.3]: https://github.com/akbarpambudi/go-point-of-sales/compare/1.4.2...v1.4.3
[1.4.2]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.4.1...1.4.2
[v1.4.1]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.4.0...v1.4.1
[v1.4.0]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.3.0...v1.4.0
[v1.3.0]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/akbarpambudi/go-point-of-sales/compare/v1.1.0...v1.2.0
[v1.1.0]: https://github.com/akbarpambudi/go-point-of-sales/compare/v0.0.1...v1.1.0
