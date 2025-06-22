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
  local tool_path=""
  local tool_name=""

  for tool_path in tools/*; do
    tool_name="$(basename "$tool_path")"
    echo "Checking $tool_name"
#    if ! check-changes "$tool_path"; then
#      echo "Skipping $tool_name"
#      continue
#    fi
    if ! _check_release "$tool_path"; then
      echo "Skipping $tool_name"
      continue
    fi
    echo "Perform release for $tool_name"
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
  if [ -n "$(git ls-files --others --exclude-standard -- "$relpath")" ]; then
    echo "There are untracked files."
    return 1
  fi
}

# check if a release is necessary
function _check_release() {
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

  _update_version "$tool_name" "$prev_version" "$next_version"
  git add "$(_tool_formula "$tool_name")"
  git commit -m "release: $tool_name $next_version"
  git push

  if _is_go_tool "$tool_name"; then
    _build_go_tool "$tool_name"
  fi

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

  grep -E 'prev_version "[a-z0-9]+"' "$tool_formula" | $SED -n 's/.*prev_version "\([^"]*\)".*/\1/p'
}

# get tool commit id from git log
function _tool_git_commit_id() {
  local tool_name=""
  local tool_path=""

  tool_name="$1"
  tool_path="$(_tool_path "$tool_name")"

  git log -n 1 --pretty='format:%H' -- ':!Makefile' ':!README.md' "$tool_path"
}

function _is_go_tool() {
  local tool_name=""

  tool_name="$1"

  [[ -f "$(_tool_path "$tool_name)")/go.mod" ]]
}

function _build_go_tool() {
  local tool_name=""

  tool_name="$1"

  make -C "$(_tool_path "$tool_name")" build
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
