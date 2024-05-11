# BiomassX

## Prerequisites

### The build tools

This software requires

* docker/docker-compse
* make build tools

### OpenAPI documentation

The OpenAPI (formerly swagger) documentation requires:
* swaggo/swag          (command line tool).
* swaggo/files........ (API documentation display)

This is done through:
```
go get -u github.com/swaggo/swag/cmd/swag           # To define swag as a dependency in go.mod/go.sum
go get -u github.com/swaggo/files
```

## Installation


## Usage

### Update documentation (OpenAPI)

```
make doc
```

### API call examples (using curl)

#### Members API

Create a Member
```
curl -X"POST" -d '{"username":"arhuman", "email": "arhumang@gmail.com", "passwd": "secret", "phone": "+66 924970276"}' localhost:12345/api/v1/members
```

Read a Member #2
```
curl -X"GET"  localhost:12345/api/v1/members/2
```

Update Member #2
```
curl -X"PUT" -d '{"username":"arhuman2", "email": "arhumang@bitkub.com", "passwd": "secret", "phone": "+66 924970276" localhost:12345/api/v1/members/2
```

Delete Member #2
```
curl  -X"DELETE" localhost:12345/api/v1/members/2
```
