#!/bin/bash

source /scripts/utils.sh

DATA_DIR=/home/docker/.ethereum

wait_for_host_port "${BOOTSTRAP_HOST}" "${BOOTSTRAP_TCP_PORT}"

BOOTSTRAP_IP=$(get_host_ip "${BOOTSTRAP_HOST}")
VALIDATOR_ADDR=$(cat ${DATA_DIR}/address)
HOST_IP=$(hostname -i)

echo "Using validator address: ${VALIDATOR_ADDR}"

geth --config "${DATA_DIR}/config.toml" \
	--datadir "${DATA_DIR}" \
	--netrestrict "${CLUSTER_CIDR}" \
	--verbosity "${VERBOSE}" \
	--ethstats "${NODE_ID}:${NETSTATS_URL}" \
	--bootnodes "enode://${BOOTSTRAP_PUB_KEY}@${BOOTSTRAP_IP}:${BOOTSTRAP_TCP_PORT}" \
	--mine \
	-unlock "${VALIDATOR_ADDR}" \
	--password /dev/null \
	--light.serve 50 \
	--metrics
