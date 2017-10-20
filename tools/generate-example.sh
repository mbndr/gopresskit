#!/bin/bash
# generate the example project to the gh-pages docs/ folder

rootPath="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."

cd "$rootPath"

go run cmd/gopresskit/main.go -input "example/" -output "docs/"
