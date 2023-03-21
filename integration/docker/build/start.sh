#!/bin/bash

echo "Building binaries..."

set -e
set -x

GOBIN=/usr/local/go/bin/go

$GOBIN build -buildvcs=false -o ./build/linux/thetasubchain ./cmd/thetasubchain
$GOBIN build -buildvcs=false -o ./build/linux/thetasubcli ./cmd/thetasubcli

set +x 

echo "Done."



