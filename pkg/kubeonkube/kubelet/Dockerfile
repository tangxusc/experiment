FROM ubuntu
RUN apt install sshpass wget -y
RUN wget https://dl.k8s.io/v1.18.4/kubernetes-server-linux-amd64.tar.gz && tar -zxvf kubernetes-server-linux-amd64.tar.gz && cp  kubernetes/server/bin/kubelet /usr/bin/kubelet && kubernetes/server/bin/kube-proxy /usr/bin/kube-proxy

ADD kubelet-config.yaml /home/
ADD kubeproxy-config.yaml /home/
ADD daemon.json /home/

ADD kubelet.service /home/
ADD kubeproxy.service /home/
ADD install.sh /home/

WORKDIR /home/
CMD ['sh /home/install.sh']