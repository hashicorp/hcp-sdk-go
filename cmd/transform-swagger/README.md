# Transform Swagger

The Transform Swagger tool allows for customization of the swagger files under `/specs`. 

The service swagger files are generated using the `protoc-gen-openapiv2` protoc plugin, part of the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) set of tools. The Go SDK generator tool [go-swagger](https://github.com/go-swagger/go-swagger) uses those swagger files to generate SDK clients and models.

In some cases, go-swagger needs additional data to generate a working SDK client. This is where this tool comes in to play. It first parses each swagger file into a mutable struct, applies any necessary changes, and then overwrites the json with the changes.

For example, type reusability across generated SDK clients requires the [x-go-type](https://goswagger.io/use/models/schemas.html#types-reusability) extension to be added to every single shared type definition.
