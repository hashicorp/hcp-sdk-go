#!/usr/bin/env bash
set -euo pipefail

# This script generates Go SDK files in the hashicorp/hcp-sdk-go repo
# It should always be run after `make transform/swagger`, as that script ensures that the generated service SDKs
# use the correct shared types (i.e. anything defined in the hashicorp/cloud protos within this repo),
# instead of generating their own versions (a quirk of go-swagger)

GREEN='\033[32m'
BOLD='\033[1m'
NA='\033[0m' # no attributes (color or format)

gen_go_sdk() {
  echo -e "creating target folder: ${BOLD}hcp-sdk-go/clients/$1/preview/$2${NA}"
  mkdir -p "$GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/$1/preview/$2"

  echo -e "generating sdk for ${BOLD}$1${NA} (version: ${BOLD}$2${NA})"
  # this command includes basic swagger validation
  swagger generate client \
    -f $GOPATH/src/github.com/hashicorp/cloud-api/specs/$1/preview/$2/*.swagger.json \
    -t $GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/$1/preview/$2 \
    -q \
    -A "$1"
    
  echo -e "${GREEN}${BOLD}sdk generated!\n${NA}"
}

# start of gen script
# verify local copy of hcp-sdk-go exists
[[ -d "$GOPATH/src/github.com/hashicorp/hcp-sdk-go" ]] \
&& echo "" \
|| echo -e "error: directory '$GOPATH/src/github.com/hashicorp/hcp-sdk-go' does not exist"

echo -e "removing old clients\n"
rm -rf $GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients

cd specs/hashicorp/cloud

# generate shared models
mkdir -p $GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1

for f in *; do
  echo -e "generating shared ${BOLD}$f${NA} SDK models\n"
  swagger generate model \
  -f $GOPATH/src/github.com/hashicorp/cloud-api/specs/hashicorp/cloud/$f/*.swagger.json \
  -t $GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1 \
  -q
done

echo -e "generating shared ${BOLD}external${NA} SDK models\n"
swagger generate model \
-f $GOPATH/src/github.com/hashicorp/cloud-api/specs/external/external.swagger.json \
-t $GOPATH/src/github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1 \
-q

cd ../..

# generate clients from each spec file
for f in *; do
  if [[ -d "$f" ]]; then

    # skip shared swagger definitions under the 'hashicorp/cloud' and 'external' directories
    if [[ "$f" == "hashicorp" ]] || [[ "$f" == "external" ]]; then
      continue
    fi

    # TODO: generate preview AND stable
    cd "$f"/preview

    for d in *; do
      gen_go_sdk "$f" "$d"  # service name, version
    done

    cd ../..
  fi
done