# 모든 노드 공통

# 시스템 업데이트
sudo apt update && sudo apt upgrade -y

# 스왑 끄기 (쿠버네티스 필수 조건)
sudo swapoff -a
sudo sed -i '/ swap / s/^/#/' /etc/fstab

# 커널 모듈
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF
sudo modprobe overlay
sudo modprobe br_netfilter

# 네트워크 설정
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
net.ipv4.ip_forward=1
EOF
sudo sysctl --system

