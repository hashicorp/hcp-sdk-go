#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

# This script fetches the latest public API specs for a given HCP service ($1) 
# from the central spec repo, cloud-api.

SCRIPTS_DIR=$(dirname "${BASH_SOURCE[0]}")

service=$1

# Copy the latest service specs into a temporary directory in preparation for SDK generation.
mkdir -p "$SCRIPTS_DIR"/../temp/"$service"
cp -r "$GOPATH"/src/github.com/hashicorp/cloud-api/specs/"$service" "$SCRIPTS_DIR"/../temp
cp -r "$GOPATH"/src/github.com/hashicorp/cloud-api/specs/cloud-shared "$SCRIPTS_DIR"/../temp
cp -r "$GOPATH"/src/github.com/hashicorp/cloud-api/specs/external "$SCRIPTS_DIR"/../temp