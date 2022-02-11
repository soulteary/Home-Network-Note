#!/bin/bash

# author @soulteary
# date 2022/02/10

echo "Downloading and installing docker..."
apt install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common
curl -fsSL "https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/ubuntu/gpg" | sudo apt-key add -
echo "deb https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/ubuntu" | tee /etc/apt/sources.list.d/docker.list > /dev/null
apt update && apt install -y docker-ce
docker -v
echo "Docker Installed."



FILENAME=$(echo "docker-compose-`uname -s`-`uname -m`" | tr [:upper:] [:lower:])
VERSION=2.2.3
echo "鉴于目前没有公开的、长期稳定的 GitHub Release 镜像"
echo "我将 ${VERSION} 版本的 compose 上传到了 https://gitee.com/soulteary/docker-compose-mirror"
echo "请访问上面的地址，下载 ${FILENAME} 文件后，进行手动安装"
echo ""
echo "安装命令: "
echo "cp ${FILENAME} /usr/local/bin/docker-compose"
echo "chmod +x /usr/local/bin/docker-compose"
echo ""
echo ""
echo "如想针对下载后文件进行校验，可下载相同文件名称的 `.sha256` 文件"
echo "执行命令："
echo "shasum -c ${FILENAME}.sha256"
echo ""

