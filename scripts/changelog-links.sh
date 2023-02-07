#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


# This script rewrites [GH-nnnn]-style references in the CHANGELOG.md file to
# be Markdown links to the given github issues.
#
# This is run manually. Eventually will be added to an automated release process.

set -e

if [[ ! -f CHANGELOG.md ]]; then
  echo "ERROR: CHANGELOG.md not found in pwd."
  echo "Please run this from the root of the SDK repository"
  exit 1
fi

if [[ $(uname) == "Darwin" ]]; then
  echo "Using BSD sed"
  SED="sed -i.bak -E -e"
else
  echo "Using GNU sed"
  SED="sed -i.bak -r -e"
fi

SDK_URL="https:\/\/github.com\/hashicorp\/hcp-sdk-go\/pull"

$SED "s/GH-([0-9]+)/\[#\1\]\($SDK_URL\/\1\)/g" -e 's/\[\[#(.+)([0-9])\)]$/(\[#\1\2))/g' CHANGELOG.md

rm CHANGELOG.md.bak