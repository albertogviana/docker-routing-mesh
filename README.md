# Docker Routing Mesh

The goal of this project is to demonstrate how it works.

## Compile
```
go build
```

If you are on mac and you would like to build a package for linux
```
GOOS=linux GOARCH=amd64 go build
```

## Build a new image
```
docker build -t albertogviana/docker-routing-mesh .
```