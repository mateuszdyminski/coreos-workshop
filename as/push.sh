#!/bin/bash

echo "Pushing AS container started..."
set -e
docker build -t mateuszdyminski/as:latest "."
docker push mateuszdyminski/as
set +e
echo "AS pushed to quay.io!"