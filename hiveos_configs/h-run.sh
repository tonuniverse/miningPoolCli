#!/usr/bin/env bash

# cd `dirname $0`

CUSTOM_DIR=$(dirname "$BASH_SOURCE")
. $CUSTOM_DIR/h-manifest.conf

[[ -z $CUSTOM_LOG_BASEDIR ]] && echo -e "No CUSTOM_LOG_BASEDIR is set" && exit 1
[[ -z $CUSTOM_CONFIG_FILENAME ]] && echo -e "No CUSTOM_CONFIG_FILENAME is set" && exit 1
[[ ! -f $CUSTOM_CONFIG_FILENAME ]] && echo -e "Custom config \"$CUSTOM_CONFIG_FILENAME\" is not found" && exit 1

mkdir -p $CUSTOM_LOG_BASEDIR
touch $CUSTOM_LOG_BASENAME.log

echo "PATH CONFIG: ${CUSTOM_CONFIG_FILENAME}"
POOL_ID=`jq .poolId ${CUSTOM_CONFIG_FILENAME}`
POOL_ID=${POOL_ID:1:-1}

echo "POOL_ID: ${POOL_ID}"

./miningPoolCli -stats -pool-id=$POOL_ID | tee --append $CUSTOM_LOG_BASENAME.log