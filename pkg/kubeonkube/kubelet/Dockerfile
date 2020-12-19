FROM ubuntu
RUN apt update && apt install wget -y
RUN wget -O kubernetes-server-linux-amd64.tar.gz https://dl.k8s.io/v1.18.4/kubernetes-server-linux-amd64.tar.gz && tar -zxvf kubernetes-server-linux-amd64.tar.gz && rm kubernetes-server-linux-amd64.tar.gz
RUN cp kubernetes/server/bin/kubelet /home/kubelet && cp kubernetes/server/bin/kube-proxy /home/kube-proxy && cp kubernetes/server/bin/kubectl /home/kubectl

ADD kubelet-config.yaml /home/
ADD kubeproxy-config.yaml /home/
ADD daemon.json /home/

ADD kubelet.service /home/
ADD kubeproxy.service /home/
ADD install.sh /home/
ADD fileserver /home/fileserver

WORKDIR /home/
CMD ['/home/fileserver']
