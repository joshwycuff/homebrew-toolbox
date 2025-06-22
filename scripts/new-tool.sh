#!/usr/bin/env bash
set -euo pipefail

# Default values
NAME=""
TEMPLATE=""

usage() {
  echo "Usage: $0 [--template] <NAME>"
  exit 1
}

# Parse options
POSITIONAL_ARGS=()
while [[ $# -gt 0 ]]; do
  case "$1" in
    -t|--template)
      if [[ -z "${2:-}" ]]; then
        echo "Error: --template requires a value" >&2
        usage
      fi
      TEMPLATE="$2"
      shift 2
      ;;
    -h|--help)
      usage
      ;;
    -*)
      echo "Unknown option: $1" >&2
      usage
      ;;
    *)
      POSITIONAL_ARGS+=("$1")
      shift
      ;;
  esac
done

# Restore positional arguments
if [[ ${#POSITIONAL_ARGS[@]} -gt 0 ]]; then
  set -- "${POSITIONAL_ARGS[@]}"
fi

# Check required arguments
if [[ $# -lt 1 ]]; then
  echo "Error: NAME is required" >&2
  usage
fi

NAME="$1"

if [[ -n "$TEMPLATE" ]]; then
  cp -r "tools/$TEMPLATE/" "tools/$NAME/"
  rg -l "$TEMPLATE" "tools/$NAME/" | xargs gsed -i "s/$TEMPLATE/$NAME/g"
  cp "Formula/$TEMPLATE.rb" "Formula/$NAME.rb"
  gsed -i "s/$TEMPLATE/$NAME/g" "Formula/$NAME.rb"
  if [[ "$TEMPLATE" = "example-py" ]]; then
    mv "tools/$NAME/$TEMPLATE" "tools/$NAME/$NAME"
  elif [[ "$TEMPLATE" = "example-sh" ]]; then
    mv "tools/$NAME/$TEMPLATE" "tools/$NAME/$NAME"
  fi
else
  mkdir "tools/$NAME/"
fi
