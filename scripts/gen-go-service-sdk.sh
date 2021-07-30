#!/usr/bin/env bash
set -euo pipefail

set -x

# This script regenerates the Go SDK for a given HCP service ($1).

# The steps are:
# 1. Remove the original SDK files for the service if they exist.
# 2. Fetch the latest public API specs for the service from the central spec repo.
# 3. Run temporary transformations on those specs to prepare them for SDK generation.
# 4. Iterate over each stage and version of the service specs and generate the corresponding SDK. 
#    (Note: Currently only the 'preview' stage is supported)
# 5. Remove temporary directories.

SCRIPTS_DIR=$(dirname "${BASH_SOURCE[0]}")

BOLD='\033[1m'
GREEN='\033[32m'
NA='\033[0m' # no attributes (color or format)

generate_sdk() {
  service=$1
  stage=$2
  version=$3

  if [ -d "clients/$service/$stage/$version" ]; then \
    echo "Removing original SDK from clients/$service/$stage/$version" && rm -rf "$SCRIPTS_DIR/../clients/$service/$stage/$version"; \
  fi

  echo -e "Creating target SDK directory: ${BOLD}hcp-sdk-go/clients/$service/$stage/$version${NA}"
  mkdir -p "../../../clients/$service/$stage/$version"

  rel=("$version"/*.swagger.json)
  spec="${rel[0]}"
  target="../../../clients/$service/$stage/$version"

  echo -e "Generating SDK for ${BOLD}$service${NA} (stage: ${BOLD}$stage${NA}, version: ${BOLD}$version${NA})"
  swagger generate client \
    -f "$spec" \
    -t "$target" \
    -q \
    -A "$service"
}

# Beginning of generation script.
service=$1

echo -e "Fetching latest specs for ${BOLD}$service${NA}"
hcloud repo init \
  --refresh \
  --only=cloud-api

# Copy the latest service specs into a temporary directory in preparation for SDK generation.
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/"$service" "$SCRIPTS_DIR"/../temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/cloud-shared "$SCRIPTS_DIR"/../temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/external "$SCRIPTS_DIR"/../temp

transformer=../../../cmd/transform-swagger
shared_specs=../../../temp/cloud-shared
external_spec=../../../temp/external/external.swagger.json

cd temp/"$service"

# Iterate over each stage directory.
# TODO: Eventually there will be specs under both preview/ and stable/ 
for d in *; do
  if [[ -d "$d" ]]; then
    stage=$d

    cd "$stage"
  fi

  # Iterate over each version directory.
  for f in *; do
    if [[ -d "$f" ]]; then
      version=$f
      service_spec=("$version"/*.swagger.json)

      # Transform the specs.
      echo -e "Transforming specs for ${BOLD}$service${NA} (stage: ${BOLD}$stage${NA}, version: ${BOLD}$version${NA}) in preparation for SDK generation"
      go run "$transformer" \
        -service="${service_spec[0]}" \
        -shared="$shared_specs" \
        -external="$external_spec"
      
      # Generate SDK from transformed specs.
      generate_sdk "$service" "$stage" "$version"
    fi
  done
done

# Navigate back to root.
cd ../../..
pwd
echo -e "Regenerating shared ${BOLD}external${NA} SDK models"
swagger generate model \
  -f ./temp/external/external.swagger.json \
  -t ./clients/cloud-shared/v1 \
  -q

echo -e "${GREEN}SDK for $service generated!${NA}"

cleanup() {
  # This is where hcloud clones cloud-api from which the specs are pulled.
  rm -rf "$HOME/.local/share/hcp/repos/cloud-api"
  rm -rf "$SCRIPTS_DIR"/../temp
}

trap cleanup EXIT