#!/usr/bin/env bash
DIRNAME=$(dirname "$BASH_SOURCE")
. $DIRNAME/h-manifest.conf

STATS_FILE=/hive/miners/custom/${CUSTOM_NAME}/stats.json

temp=$(jq '.temp' <<< $gpu_stats)
fan=$(jq '.fan' <<< $gpu_stats)
[[ $cpu_indexes_array != '[]' ]] &&
    temp=$(jq -c "del(.$cpu_indexes_array)" <<< $temp) &&
    fan=$(jq -c "del(.$cpu_indexes_array)" <<< $fan)

busids=$(jq '.busids' <<< $gpu_stats)

BUS_NUMBERS=()
for i in "${!busids[@]}"; do 
  BUS_NUMER_HEX=$(echo ${busids[$i]:0:2} | tr "a-z" "A-Z")
  BUS_NUM=$(echo "obase=10; ibase=16; $BUS_NUMER_HEX" | bc)
  BUS_NUMBERS+=($BUS_NUM)
done

if [[ ! -f $STATS_FILE ]]; then
    echo "$STATS_FILE not found"
else
    khs=$(jq .khs $STATS_FILE)
    hs=$(jq .hs $STATS_FILE)
    uptime=$(jq .uptime $STATS_FILE)
    bus_numbers=$(echo "${BUS_NUMBERS[@]}" | jq -s '.')

    stats=$(jq -n \
    --argjson hs "$hs" \
    --arg uptime "$uptime" \
    --arg ver "$CUSTOM_VERSION" \
    --argjson temp "$temp" \
    --argjson fan "$fan" \
    --argjson bus_numbers "$bus_numbers" \
    '{"hs": $hs, "hs_units": "mhs", "algo": "sha256", "uptime": $uptime, "ver": $ver, "temp": $temp, "fan": $fan, "bus_numbers": $bus_numbers}')
fi
