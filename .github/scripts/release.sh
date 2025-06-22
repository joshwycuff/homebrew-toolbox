#!/usr/bin/env bash

set -euo pipefail

# Check if we're in the git root directory
if [ ! -d .git ]; then
    echo "This script must be run from the git root directory"
    exit 1
fi

# Check if we're in the homebrew-toolbox directory
if [ ! "$(basename "$(pwd)")" = "homebrew-toolbox" ]; then
    echo "This script must be run from the homebrew-toolbox directory"
    exit 1
fi

if [[ "$(uname)" == "Darwin" ]]; then
  SED="gsed"
else
  SED="sed"
fi

function main() {
  local tool_name=""
  local tool_path=""

  for tool_path in tools/*; do
    tool_name="$(basename "$tool_path")"
    echo "Checking $tool_name"
    if ! _check_changes "$tool_path"; then
      echo "Skipping $tool_name"
      continue
    fi
    if _check_version "$tool_name"; then
      echo "Skipping $tool_name"
      continue
    fi
    echo "Performing release for $tool_name"
    _perform_release "$tool_name"
  done
}

# check if there are any changes in the given path
function _check_changes() {
  local relpath="$1"

  # Check for unstaged changes
  if ! git diff --quiet -- "$relpath"; then
    echo "There are unstaged changes."
    return 1
  fi

  # Check for staged but uncommitted changes
  if ! git diff --cached --quiet -- "$relpath"; then
    echo "There are staged but uncommitted changes."
    return 1
  fi

  # Check for untracked files
  if [[ "$(git ls-files --others --exclude-standard -- "$relpath" | wc -l)" -gt 0 ]]; then
    echo "There are untracked files."
    return 1
  fi
}

# check if homebrew version matches most recent git commit id
function _check_version() {
  local tool_name=""
  local prev_version=""
  local next_version=""

  tool_name="$1"
  prev_version="$(_tool_homebrew_version "$tool_name")"
  next_version="$(_tool_git_commit_id "$tool_name")"

  [[ "$next_version" = "$prev_version" ]]
}

function _perform_release() {
  local tool_name=""
  local prev_version=""
  local next_version=""

  tool_name="$1"

  prev_version="$(_tool_homebrew_version "$tool_name")"
  next_version="$(_tool_git_commit_id "$tool_name")"

  echo "Updating version from $prev_version to $next_version"

  _update_version "$tool_name" "$prev_version" "$next_version"
  git add "$(_tool_formula "$tool_name")"

  echo "Creating commit"
  git commit -m "release: $tool_name $next_version"

  echo "Pushing"
  git push

  if _is_go_tool "$tool_name"; then
    echo "Building"
    _build_go_tool "$tool_name"
  fi

  echo "Creating release"
  gh release create "$next_version" "$(_tool_path "$tool_name")"
}

# update prev_version in homebrew formula
function _update_version() {
  local tool_name=""
  local prev_version=""
  local next_version=""

  tool_name="$1"
  prev_version="$2"
  next_version="$3"

  $SED -i "s/$prev_version/$next_version/g" "$(_tool_formula "$tool_name")"
}

# get tool prev_version from formula
function _tool_homebrew_version() {
  local tool_name=""
  local tool_formula=""

  tool_name="$1"
  tool_formula="$(_tool_formula "$tool_name")"

  grep -E 'version "[a-z0-9]+"' "$tool_formula" | $SED -n 's/.*version "\([^"]*\)".*/\1/p'
}

# get tool commit id from git log
function _tool_git_commit_id() {
  local tool_name=""
  local tool_dirpath=""

  tool_name="$1"
  tool_dirpath="$(_tool_dirpath "$tool_name")"

  git log -n 1 --pretty='format:%H' -- ':!Makefile' ':!README.md' "$tool_dirpath"
}

function _is_go_tool() {
  local tool_name=""
  tool_name="$1"
  [[ -f "$(_tool_dirpath "$tool_name)")/go.mod" ]]
}

function _build_go_tool() {
  local tool_name=""
  tool_name="$1"
  make -C "$(_tool_dirpath "$tool_name")" build
}

function _tool_path() {
  local tool_name
  tool_name="$1"
  printf "$(_tool_dirpath "$tool_name")/%s" "$tool_name"
}

function _tool_dirpath() {
  local tool_name
  tool_name="$1"
  printf "tools/%s" "$tool_name"
}

function _tool_formula() {
  local tool_name
  tool_name="$1"
  printf "Formula/%s.rb" "$tool_name"
}

main
