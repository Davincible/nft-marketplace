#!/bin/bash

source /scripts/utils.sh

wait_for_host_port "$(get_host_ip netstats)" 3000

DATA_DIR=/home/docker/.ethereum

accounts=$(find "${DATA_DIR}/keystore" -type f -printf "%f,")

geth --config ${DATA_DIR}/config.toml \
	--datadir ${DATA_DIR} \
	--netrestrict "${CLUSTER_CIDR}" \
	--verbosity "${VERBOSE}" \
	--ethstats "${NODE_ID}:${NETSTATS_URL}" \
	--unlock "${accounts}" \
	--password /dev/null
