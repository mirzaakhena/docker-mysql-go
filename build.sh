#!/bin/bash

rm -rf product

GOOS=linux GOARCH=amd64 go build -o product ./main.go

docker build -t product .

docker-compose up -d