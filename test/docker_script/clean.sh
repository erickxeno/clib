#!/usr/bin/env bash
echo "清理mysql, redis, es, mongo ..."
docker rm mysql redis es mongo --force &>/dev/null