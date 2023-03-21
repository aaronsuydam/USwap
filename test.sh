#!/bin/sh

ng test &
go run ./src/server/main.go &

wait