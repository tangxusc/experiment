#!/bin/bash
#   set ff=unix
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
mkdir -p /etc/docker/
cat >> /etc/docker/daemon.json <<EOF
{
  "registry-mirrors": ["https://registry.docker-cn.com"]
}
EOF
systemctl restart docker

wget https://dl.k8s.io/v1.18.4/kubernetes-server-linux-amd64.tar.gz 
tar -zxvf kubernetes-server-linux-amd64.tar.gz

cp kubernetes/server/bin/kubelet /usr/bin/kubelet 
cp kubernetes/server/bin/kube-proxy /usr/bin/kube-proxy 

mkdir -p /etc/kubernetes/
cat > /etc/kubernetes/kubelet-config.yaml << EOF
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    cacheTTL: 0s
    enabled: true
  x509:
    clientCAFile: /etc/kubernetes/ca.pem
authorization:
  mode: Webhook
  webhook:
    cacheAuthorizedTTL: 0s
    cacheUnauthorizedTTL: 0s
clusterDNS:
- 10.96.0.10
clusterDomain: cluster.local
cpuManagerReconcilePeriod: 0s
evictionPressureTransitionPeriod: 0s
fileCheckFrequency: 0s
healthzBindAddress: 127.0.0.1
healthzPort: 10248
httpCheckFrequency: 0s
imageMinimumGCAge: 0s
kind: KubeletConfiguration
nodeStatusReportFrequency: 0s
nodeStatusUpdateFrequency: 0s
rotateCertificates: true
runtimeRequestTimeout: 0s
staticPodPath: /etc/kubernetes/manifests
streamingConnectionIdleTimeout: 0s
syncFrequency: 0s
volumeStatsAggPeriod: 0s
EOF

cat > /etc/kubernetes/ca.pem << EOF
-----BEGIN CERTIFICATE-----
MIIDsDCCApigAwIBAgIUXk+OhfVUJWEOCrPEfHVDSEw1ymQwDQYJKoZIhvcNAQEL
BQAwXjELMAkGA1UEBhMCQ04xETAPBgNVBAgTCHNoZW56aGVuMREwDwYDVQQHEwhz
aGVuemhlbjELMAkGA1UEChMCQ0ExDzANBgNVBAsTBlN5c3RlbTELMAkGA1UEAxMC
Q0EwHhcNMjAxMjEwMDIyODAwWhcNMjUxMjA5MDIyODAwWjBeMQswCQYDVQQGEwJD
TjERMA8GA1UECBMIc2hlbnpoZW4xETAPBgNVBAcTCHNoZW56aGVuMQswCQYDVQQK
EwJDQTEPMA0GA1UECxMGU3lzdGVtMQswCQYDVQQDEwJDQTCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAK95y/GIudjgzzL+9547dE9tKn92lsRmmio+2l4Z
HePvJoqLUVrBdzGR0Z8EA8apQwDxMoQ6bybYzejqJKdfgcHJ+YovD6p0bL7Uac8T
sZCfiAH+MEaZjaYeHdLLaEDMTMY6ZLHHlzkmGwLm5Swo6L1O5AXg63m5dnqihaRS
jh9Le5mRVOzilDbGVEgi0E1x90C06dDo84Pm9h6IyM/iQi5OIT069htZA5TWlKj8
khR2NhjliDHllktkpyY+RE98HNz6E57lqYyryhON6OAdQwyKUWO0SH+maCW6UjC5
Uhww6APPVAK1HWxL0OSWUJiy2W915JgH/O1FaWGhL8SkYekCAwEAAaNmMGQwDgYD
VR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQIwHQYDVR0OBBYEFAOlXduM
oN66hI7Vo+PIFtZxtoCzMB8GA1UdIwQYMBaAFAOlXduMoN66hI7Vo+PIFtZxtoCz
MA0GCSqGSIb3DQEBCwUAA4IBAQBWqifHXPs/VgdGm8ij1EwzDFQgSo6ghaJMw3Ey
ZNLsI7OGUrOMyYk3BDqBD3G8n8o/bTy00GmhVx4YEf0SchI3ePAK/jjmEAkRvS9M
aQfcHatccSTsz8nS1qyTDynfmbB5O79OPTYm4si9uzwx23vHP3nhVEI7zybQS8+a
dNufNTS19dSdsWZ81Artq7x7L79drSenGgq0F7FnMqICgV+q27AUF7B05PZgKK/q
uFnCAS6yf2FkrG8R4GWimHCnxqWC/j31gbOX1eTsLiq5m/8PLRK00CwyGXYiYvmB
q4j/LWwowDdr4GdvlmlsfaQuZJ8A0VnQkVxLHEF7ZZ5mROzz
-----END CERTIFICATE-----
EOF


/usr/bin/kubelet \
--kubeconfig=/etc/kubernetes/node.config \
--config=/etc/kubernetes/kubelet-config.yaml \
--network-plugin=cni \
--pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.1 

cat > /etc/kubernetes/kubeproxy-config.yaml << EOF
kind: KubeProxyConfiguration
apiVersion: kubeproxy.config.k8s.io/v1alpha1
bindAddress: 0.0.0.0
metricsBindAddress: 0.0.0.0:10249
clientConnection:
  kubeconfig: /etc/kubernetes/node.config
clusterCIDR: 10.0.0.0/24
EOF

kube-proxy --config=/etc/kubernetes/kubeproxy-config.yaml

#在client中执行
#kubectl --kubeconfig=admin/admin.config create clusterrolebinding cluster-node --clusterrole=system:node --user=kubernetes-node --group=system:node


level=info msg="Establishing connection to apiserver" host="https://10.96.0.1:443" subsys=k8s
level=info msg="Establishing connection to apiserver" host="https://10.96.0.1:443" subsys=k8s
level=info msg="Establishing connection to apiserver" host="https://10.96.0.1:443" subsys=k8s
level=info msg="Establishing connection to apiserver" host="https://10.96.0.1:443" subsys=k8s
level=error msg="Unable to contact k8s api-server" error="Get \"https://10.96.0.1:443/api/v1/namespaces/kube-system\": x509: certificate is valid for 127.0.0.1, 192.168.0.1, 81.71.121.114, not 10.96.0.1" ipAddr="https://10.96.0.1:443" subsys=k8s
level=fatal msg="Unable to initialize Kubernetes subsystem" error="unable to create k8s client: unable to create k8s client: Get \"https://10.96.0.1:443/api/v1/namespaces/kube-system\": x509: certificate is valid for 127.0.0.1, 192.168.0.1, 81.71.121.114, not 10.96.0.1" subsys=daemon