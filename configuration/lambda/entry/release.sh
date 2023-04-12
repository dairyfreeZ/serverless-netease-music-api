#!/bin/bash

releaseDir="release"
if [[ -d $releaseDir ]]; then
    rm -rf $releaseDir
fi
mkdir $releaseDir

GOOS=linux GOARCH=amd64 go build -o "${releaseDir}/signin" "signin/main.go"