#!/bin/bash
# Copyright ApeCloud, Inc.
# Licensed under the Apache v2(found in the LICENSE file in the root directory).




# This is a script that gets run as part of the Dockerfile build
# to install dependencies for the vitess/mini image.
#
# Usage: install_mini_dependencies.sh

set -euo pipefail

# Install etcd
ETCD_VER=v3.4.9
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GITHUB_URL}

curl -k -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
mkdir -p /var/opt/etcd
sudo tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /var/opt/etcd --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

mkdir -p /var/run/etcd && chown -R vitess:vitess /var/run/etcd

# Clean up files we won't need in the final image.
rm -rf /var/lib/apt/lists/*
