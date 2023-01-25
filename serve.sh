#!/bin/sh

ng serve --open &
~/go/bin/gin --port 4201 --path . --build ./src/server/ --i --all &

wait