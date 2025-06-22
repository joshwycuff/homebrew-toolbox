#!/usr/bin/env bash

if [[ $# -eq 0 ]]; then
    man "$(apropos "." | fzf --preview 'man $(echo {} | cut -d"(" -f1)' | cut -d"(" -f1)"
else
    man "$(apropos "$@" | fzf --preview 'man $(echo {} | cut -d"(" -f1)' | cut -d"(" -f1)"
fi
