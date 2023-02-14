#!/bin/sh

ng serve --open &
go run ./src/server/main.go &

wait