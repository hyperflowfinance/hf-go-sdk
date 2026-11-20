#!/usr/bin/env bash
set -euo pipefail
shopt -s nullglob
files=(zz_generated_*.go)
((${#files[@]} > 0))
git add -f -- "${files[@]}"
git config user.name "github-actions[bot]"
git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
git commit -m "Release $VERSION with generated client"
git tag -a "$VERSION" -m "$VERSION"
git push origin "refs/tags/$VERSION"
