#!/bin/sh

ng serve &
go run ./src/server/main.go &

wait