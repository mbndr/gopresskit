#!/bin/bash
# compiles all files in the static/ folder to a go source file
# go-bindata has to be installed

rootPath="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."

go-bindata -prefix "static/" -pkg "presskit" -o "$rootPath/bindata.go" "$rootPath/static/..."
