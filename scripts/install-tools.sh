#!env bash
set -xe

PROJECT_ROOT_DIR=$(cd $(dirname $0)/..; pwd)
BIN_DIR=$PROJECT_ROOT_DIR.bin
export GOBIN=$BIN_DIR

if [ ! -e $BIN_DIR/golangci-lint ]; then
    GOBIN=$BIN_DIR go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
fi

if [ ! -e $BIN_DIR/openapi-generator-cli.jar ]; then
    curl -L -o $BIN_DIR/openapi-generator-cli.jar https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.0.0/openapi-generator-cli-6.0.0.jar
fi
