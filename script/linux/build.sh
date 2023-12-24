#!/bin/bash

# Code format
go fmt

#Script for building an application for *NIX

cd ../..

git pull --all

# Code build
go build -ldflags '-s -w' -v -o app

pkill app
cp --update app /bin