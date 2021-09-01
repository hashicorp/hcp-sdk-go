#!/usr/bin/env bash
set -euo pipefail

# This script fetches the latest public API specs for a given HCP service ($1) 
# from the central spec repo, cloud-api.

SCRIPTS_DIR=$(dirname "${BASH_SOURCE[0]}")

BOLD='\033[1m'
NA='\033[0m' # no attributes (color or format)

service=$1

echo -e "Fetching latest specs for ${BOLD}$service${NA}"
hcloud repo init \
  --refresh \
  --only=cloud-api

# Copy the latest service specs into a temporary directory in preparation for SDK generation.
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/"$service" "$SCRIPTS_DIR"/../temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/cloud-shared "$SCRIPTS_DIR"/../temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/external "$SCRIPTS_DIR"/../temp