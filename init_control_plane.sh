# 컨트롤플레인 초기화
sudo kubeadm init --pod-network-cidr=192.168.0.0/16

# Apply Calio Manifest
kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.3/manifests/calico.yaml

# check master node status
kubectl get pods -n kube-system
kubectl get nodes

