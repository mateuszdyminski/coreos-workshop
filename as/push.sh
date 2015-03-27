#!/bin/bash

echo "Pushing AS container started..."
set -e
docker build -t quay.io/mateuszdyminski/as:latest "."
docker push quay.io/mateuszdyminski/as
set +e
echo "AS pushed to quay.io!"