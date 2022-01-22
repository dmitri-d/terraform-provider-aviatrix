# Go API client for api_client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./api_client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GatewaysDelete**](docs/DefaultApi.md#gatewaysdelete) | **Delete** /gateways | 
*DefaultApi* | [**GatewaysGet**](docs/DefaultApi.md#gatewaysget) | **Get** /gateways | 
*DefaultApi* | [**GatewaysIdGet**](docs/DefaultApi.md#gatewaysidget) | **Get** /gateways/{id} | 
*DefaultApi* | [**GatewaysPost**](docs/DefaultApi.md#gatewayspost) | **Post** /gateways | 
*DefaultApi* | [**SubnetsGet**](docs/DefaultApi.md#subnetsget) | **Get** /subnets | 
*DefaultApi* | [**TasksGet**](docs/DefaultApi.md#tasksget) | **Get** /tasks | 
*DefaultApi* | [**VpcsAwsPost**](docs/DefaultApi.md#vpcsawspost) | **Post** /vpcs/aws | 
*DefaultApi* | [**VpcsDelete**](docs/DefaultApi.md#vpcsdelete) | **Delete** /vpcs | 
*DefaultApi* | [**VpcsGet**](docs/DefaultApi.md#vpcsget) | **Get** /vpcs | 
*DefaultApi* | [**VpcsIdGet**](docs/DefaultApi.md#vpcsidget) | **Get** /vpcs/{id} | 


## Documentation For Models

 - [AuditStates](docs/AuditStates.md)
 - [AwsInstanceSizes](docs/AwsInstanceSizes.md)
 - [AwsRegions](docs/AwsRegions.md)
 - [CloudProviders](docs/CloudProviders.md)
 - [CreateGateway](docs/CreateGateway.md)
 - [CreateVpc](docs/CreateVpc.md)
 - [Gateway](docs/Gateway.md)
 - [GatewayStates](docs/GatewayStates.md)
 - [InternalServerError](docs/InternalServerError.md)
 - [NotFoundError](docs/NotFoundError.md)
 - [Subnet](docs/Subnet.md)
 - [SubnetVisibility](docs/SubnetVisibility.md)
 - [SubnetVpc](docs/SubnetVpc.md)
 - [Task](docs/Task.md)
 - [TaskEvents](docs/TaskEvents.md)
 - [UnauthorizedError](docs/UnauthorizedError.md)
 - [UnexpectedError](docs/UnexpectedError.md)
 - [UnprocessableEntityError](docs/UnprocessableEntityError.md)
 - [Vpc](docs/Vpc.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: CID
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: CID and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


