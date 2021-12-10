#!/usr/bin/env bash

. h-manifest.conf

# Read gpu stats
temp=$(gpu-stats | jq ".temp")
fan=$(gpu-stats | jq ".fan")

[[ $cpu_indexes_array != '[]' ]] && #remove Internal Gpus
    temp=$(jq -c "del(.$cpu_indexes_array)" <<< $temp) &&
    fan=$(jq -c "del(.$cpu_indexes_array)" <<< $fan)

# Read miner stats
uptime=0
if [ -f "stats.json" ]; then
  khs=`jq .hash_rate stats.json`
  uptime=`jq .uptime stats.json`
else
  echo "No stats found"
  khs=0
fi

# Uptime
ver=$CUSTOM_VERSION
hs_units="mhs"

# Performance
# stats=$(jq -nc \
#         --arg total_khs "$khs" \
#         --arg hs_units "$hs_units" \
#         --argjson temp "$temp" \
#         --argjson fan "$fan" \
#         --arg uptime "$uptime" \
#         --arg algo "toncoin" \
#         --arg ver "$ver" \
#         '{$total_khs, $hs_units, $temp, $fan, $uptime, $algo, $ver}')

stats=$(jq -nc \
        --arg total_khs "$khs" \
        --arg hs_units "$hs_units" \
        --arg uptime "$uptime" \
        --arg algo "sha256" \
        --arg ver "$ver" \
        '{$total_khs, $hs_units, $uptime, $algo, $ver}')