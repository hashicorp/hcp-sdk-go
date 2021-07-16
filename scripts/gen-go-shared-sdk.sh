#!/usr/bin/env bash
set -euo pipefail

set -x

BOLD='\033[1m'
GREEN='\033[32m'
NA='\033[0m' # no attributes (color or format)

echo -e "Fetching latest specs for ${BOLD}cloud-shared${NA}"
hcloud repo init \
  --refresh \
  --only=cloud-api

rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/cloud-shared temp
rsync -a "$HOME"/.local/share/hcp/repos/cloud-api/specs/external temp

external_spec="$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/temp/external/external.swagger.json

version="v1"
cd temp/cloud-shared

echo -e "Creating target shared SDK directory: ${BOLD}hcp-sdk-go/clients/cloud-shared/$version${NA}"
mkdir -p "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/"$version"

# Iterate over each shared directory.
for d in *; do
    shared="$d"
    if [[ -d "$shared" ]]; then
        cd "$shared"
        for f in *; do
            spec=$f
            echo -e "Generating shared SDK models (version: ${BOLD}$version${NA})"
            swagger generate model \
                -f "$spec" \
                -t "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/"$version" \
                -q
        done
    fi
    
    cd ..
done

echo -e "Regenerating shared ${BOLD}external${NA} SDK models"
swagger generate model \
  -f "$external_spec" \
  -t "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/"$version" \
  -q

echo -e "${GREEN}SDK for cloud-shared generated!${NA}"

cleanup() {
  # This is where hcloud clones cloud-api from which the specs are pulled.
  rm -rf "$HOME/.local/share/hcp/repos/cloud-api"
  rm -rf "$GOPATH"/src/github.com/hashicorp/hcp-sdk-go/temp
}

trap cleanup EXIT