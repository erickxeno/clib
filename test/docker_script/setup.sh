#!/usr/bin/env bash
# set -x # 用来在运行结果之前，先输出执行的那一行命令
# set -u # 遇到不存在变量，报错并停止运行
# set -e # 出错时停止执行
# set -eo pipefail # 管道中子指令失败停止执行
echo "检查Docker......"
docker -v
if [ $? -ne 0 ]; then
    echo "安装docker环境..."
    if [[ `uname` == 'Darwin' ]]; then
        which -s brew
        if [[ $? != 0 ]]; then
            echo '缺brew, Installing Homebrew...'
            /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
        fi
        echo "安装docker, 同时自带docker-compose..."
        brew install --cask docker
    else
        echo "it's linux"
        echo "安装docker..."
        curl -fsSL https://get.docker.com -o get-docker.sh
        sh get-docker.sh

        echo "安装docker-compose..."
        sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
        sudo chmod +x /usr/local/bin/docker-compose
        docker-compose --version
    fi
else
    echo "docker已安装"
fi

echo "构建mysql, redis, es, mongo环境..."
curr_dir=`pwd`
basepath=$(cd `dirname $0`; pwd)
cd $basepath
docker rm -v mysql redis es mongo --force &>/dev/null
docker-compose build --no-cache
docker-compose up -d -V
cd $curr_dir

echo "等待 mysql 就绪..."
while ! docker exec -it mysql mysql --user=root --password=root -e "SELECT 1" &>/dev/null; do
    sleep 1
done
echo "mysql 已就绪"