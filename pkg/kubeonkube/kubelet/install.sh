#!/bin/bash
#1.复制文件
sshpass -e scp -P 22 /pki/ca/ca.pem ${NODE_USERNAME}@${NODE_ADDRESS}:/etc/kubernetes/
sshpass -e scp -P 22 /home/node/node.config ${NODE_USERNAME}@${NODE_ADDRESS}:/etc/kubernetes/

sshpass -e scp -P 22 /usr/bin/kubelet ${NODE_USERNAME}@${NODE_ADDRESS}:/usr/bin/
sshpass -e scp -P 22 /usr/bin/kubeproxy ${NODE_USERNAME}@${NODE_ADDRESS}:/usr/bin/

sshpass -e scp -P 22 /home/kubelet-config.yaml ${NODE_USERNAME}@${NODE_ADDRESS}:/etc/kubernetes/
sshpass -e scp -P 22 /home/kubeproxy-config.yaml ${NODE_USERNAME}@${NODE_ADDRESS}:/etc/kubernetes/

sshpass -e scp -P 22 /home/kubelet.service ${NODE_USERNAME}@${NODE_ADDRESS}:/usr/lib/systemd/system/
sshpass -e scp -P 22 /home/kubeproxy.service ${NODE_USERNAME}@${NODE_ADDRESS}:/usr/lib/systemd/system/

sshpass -e ssh ${NODE_USERNAME}@${NODE_ADDRESS} 'mkdir -p /etc/docker/'
sshpass -e scp -P 22 /home/daemon.json ${NODE_USERNAME}@${NODE_ADDRESS}:/etc/docker/

#2.安装docker
sshpass -e ssh ${NODE_USERNAME}@${NODE_ADDRESS} 'curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun'
sshpass -e ssh ${NODE_USERNAME}@${NODE_ADDRESS} 'systemctl restart docker'

#3.启动kubelet
sshpass -e ssh ${NODE_USERNAME}@${NODE_ADDRESS} 'systemctl daemon-reload&&systemctl start kubelet&&systemctl enable kubelet'
#4.启动kubeproxy
sshpass -e ssh ${NODE_USERNAME}@${NODE_ADDRESS} 'systemctl start kubeproxy&&systemctl enable kubeproxy'