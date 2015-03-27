#!/bin/bash

echo "Pushing AS-LB container started..."
set -e
docker build -t mateuszdyminski/as-lb:latest "."
docker push mateuszdyminski/as-lb
set +e
echo "AS-LB pushed to quay.io!"