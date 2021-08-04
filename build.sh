#!/bin/bash

go mod tidy
go build src/main.go
./main
