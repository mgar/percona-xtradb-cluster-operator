#!/bin/bash

pmm-admin add mysql \
    --skip-connection-check \
    --server-url="https://${PMM_AGENT_SERVER_USERNAME}:${PMM_AGENT_SERVER_PASSWORD}@${PMM_AGENT_SERVER_ADDRESS}/" \
    --server-insecure-tls \
    --query-source perfschema \
    --username=${DB_USER} --password=${DB_PASSWORD} --cluster=${DB_CLUSTER} \
    "${PMM_AGENT_SETUP_NODE_NAME}" \
    "${DB_HOST}:${DB_PORT}"
