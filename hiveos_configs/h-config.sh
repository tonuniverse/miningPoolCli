#!/usr/bin/env bash

CUSTOM_DIR=$(dirname "$BASH_SOURCE")

. $CUSTOM_DIR/h-manifest.conf

[[ -z $CUSTOM_TEMPLATE ]] && echo -e "CUSTOM_TEMPLATE is empty" && return 1

echo "{\"poolId\": \"$CUSTOM_TEMPLATE\"}" > $CUSTOM_CONFIG_FILENAME