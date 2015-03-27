#!/bin/bash

echo "Pushing AS-LB container started..."
set -e
docker build -t quay.io/mateuszdyminski/as-lb:latest "."
docker push quay.io/mateuszdyminski/as-lb
set +e
echo "AS-LB pushed to quay.io!"