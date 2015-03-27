#!/bin/bash

GO=${GO:-$(which go)}

echo "Building as..."
set -e
export GOOS="linux" && export GOARCH="amd64" && export GOPATH="$PWD" && go build -o bin/main main
set +e
echo "As builed with success!"
