#wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
#wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
#wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
#chmod +x cfssl*
#wget https://dl.k8s.io/v1.18.4/kubernetes-server-linux-amd64.tar.gz && tar -zxvf kubernetes-server-linux-amd64.tar.gz
FROM ubuntu
ADD cfssl-certinfo_linux-amd64 /usr/bin/cfssl-certinfo
ADD cfssljson_linux-amd64 /usr/bin/cfssljson
ADD cfssl_linux-amd64 /usr/bin/cfssl
ADD kubernetes/server/bin/kubectl /usr/bin/kubectl
ADD kubernetes/server/bin/kubelet /usr/bin/kubelet
ADD kubernetes/server/bin/kube-proxy /usr/bin/kube-proxy
ADD init.sh /home/init.sh
RUN chmod a+x /home/init.sh
WORKDIR /home
CMD ['sh /home/init.sh']