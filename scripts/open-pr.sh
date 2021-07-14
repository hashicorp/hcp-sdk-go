#!/usr/bin/env bash
set -euo pipefail

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"

if [ "$CURRENT_BRANCH" != "master" ] && [ "$CURRENT_BRANCH" != "main" ]; then
  echo "ERROR: A PR can only be created on a branch of the 'master' or 'main' branch. Please checkout the 'master' or 'main' branch and try again." >&2
  exit
fi

d=$(date +%Y-%m-%d-%H-%M)
fmtd=$(date +"%D")
service=$1

git checkout -b "update-$service-sdk-$d"
git add "clients/$service/*"
git commit -m "updates $service SDK on $fmtd"
git push

gh pr create --title "Updates $service SDK" --body "Updates $service SDK as of $fmtd"

git co -