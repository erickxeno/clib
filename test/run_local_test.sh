#!/usr/bin/env bash

export CONTAINER_NAMES=(mysql redis es mongo)

rm -rf {coverage.html,c.out,coverage.xml,junit.xml}
echo "检测docker运行状态..."
bashpath=$(cd `dirname $0`; pwd)
for container_name in ${CONTAINER_NAMES[@]};
do
    exist=`docker inspect --format '{{.State.Running}}' ${container_name}`
    if [ "${exist}" != "true" ]; then
        echo "检测到docker窗口：${container_name}不存在，开始重启docker..."
        bash $bashpath/docker_script/setup.sh
        echo "docker容器重启完毕..."
        break
    fi
done
echo "所有docker容器全部ready，开始执行单测..."

echo "run go test..."
go test ./... -coverprofile=c.out -gcflags=-l
echo "export coverage.html"
go tool cover -html=c.out -o coverage.html