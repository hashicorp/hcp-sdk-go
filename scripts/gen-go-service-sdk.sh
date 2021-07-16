#!/usr/bin/env bash
set -euo pipefail

# This script regenerates the Go SDK for a given HCP service ($1).

# The steps are:
# 1. Remove the original SDK files for the service if they exist.
# 2. Fetch the latest public API specs for the service from the central spec repo.
# 3. Run temporary transformations on those specs to prepare them for SDK generation.
# 4. Iterate over each stage and version of the service specs and generate the corresponding SDK. 
#    (Note: Currently only the 'preview' stage is supported)
# 5. Remove temporary directories.

BOLD='\033[1m'
GREEN='\033[32m'
NA='\033[0m' # no attributes (color or format)

generate_sdk() {
  service=$1
  stage=$2
  version=$3

  if [ -d "clients/$service/$stage/$version" ]; then \
    echo "Removing original SDK from clients/$service/$stage/$version" && rm -rf clients/"$service/$stage/$version"; \
  fi

  echo -e "Creating target SDK directory: ${BOLD}hcp-sdk-go/clients/$service/$stage/$version${NA}"
  mkdir -p "$GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/$service/$stage/$version"

  echo -e "Generating SDK for ${BOLD}$service${NA} (stage: ${BOLD}$stage${NA}, version: ${BOLD}$version${NA})"
  swagger generate client \
    -f "$GOPATH/src/github.com/hashicorp/cloud-api/specs/$service/$stage/$version"/*.swagger.json \
    -t "$GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/$service/$stage/$version" \
    -q \
    -A "$service"
}

# Beginning of generation script.
service=$1

if [[ -d "clients/$service" ]]; then
	echo -e "Removing the original SDK of $service" && rm -rf clients/"$service"; \
  rm -rf "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/clients
fi

echo -e "Fetching latest specs for ${BOLD}$service${NA}"
hcloud repo init \
  --refresh \
  --only=cloud-api

# Copy the latest service specs into a temporary directory in preparation for SDK generation.
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/"$service" temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/cloud-shared temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/external temp

transformer="$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/cmd/transform-swagger
shared_specs="$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/temp/hashicorp/cloud
external_spec="$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/temp/external/external.swagger.json

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
        -service="$service_spec" \
        -shared="$shared_specs" \
        -external="$external_spec"
      
      # Generate SDK from transformed specs.
      generate_sdk "$service" "$stage" "$version"
    fi
  done
done

echo -e "Regenerating shared ${BOLD}external${NA} SDK models"
swagger generate model \
  -f "$external_spec" \
  -t "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1 \
  -q

echo -e "${GREEN}SDK for $service generated!${NA}"

cleanup() {
  # This is where hcloud clones cloud-api from which the specs are pulled.
  rm -rf "$HOME/.local/share/hcp/repos/cloud-api"
  rm -rf "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/temp
}

trap cleanup EXIT