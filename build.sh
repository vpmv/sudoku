#!/bin/bash

BUILDNAME="sudoku-build"
BINNAME="sudoku-bin"

docker rm -v "$BUILDNAME" || true
docker build --tag="$BUILDNAME" .
docker create --rm --name "$BINNAME" "$BUILDNAME"
docker cp "$BINNAME":/app/bin/ ./
docker rm -v "$BINNAME"
