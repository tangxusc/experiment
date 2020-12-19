#!/bin/bash
DOWNLOAD_ADDRESS=$1
#1.下载
apt update && apt install -y wget
mkdir -p /etc/kubernetes
mkdir -p /etc/docker/
wget -O /etc/kubernetes/ca.pem "${DOWNLOAD_ADDRESS}"/ca/ca.pem
wget -O /etc/kubernetes/node.config "${DOWNLOAD_ADDRESS}"/node/node.config
wget -O /usr/bin/kubelet "${DOWNLOAD_ADDRESS}"/kubelet
wget -O /usr/bin/kube-proxy "${DOWNLOAD_ADDRESS}"/kube-proxy
wget -O /usr/bin/kubectl "${DOWNLOAD_ADDRESS}"/kubectl
chmod a+x /usr/bin/kubelet && chmod a+x /usr/bin/kube-proxy && chmod a+x /usr/bin/kubectl
wget -O /etc/kubernetes/kubelet-config.yaml "${DOWNLOAD_ADDRESS}"/kubelet-config.yaml
wget -O /etc/kubernetes/kubeproxy-config.yaml "${DOWNLOAD_ADDRESS}"/kubeproxy-config.yaml
wget -O /lib/systemd/system/kubelet.service "${DOWNLOAD_ADDRESS}"/kubelet.service
wget -O /lib/systemd/system/kubeproxy.service "${DOWNLOAD_ADDRESS}"/kubeproxy.service
wget -O /etc/docker/daemon.json "${DOWNLOAD_ADDRESS}"/daemon.json

#2.安装docker
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
systemctl restart docker
#3.启动
systemctl daemon-reload&&systemctl start kubelet&&systemctl enable kubelet
systemctl start kubeproxy&&systemctl enable kubeproxy
