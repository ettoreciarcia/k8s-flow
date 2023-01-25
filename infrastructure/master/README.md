# To Do

Complete installation K3s
```curl -sfL https://get.k3s.io | sh -s -```

To deploy k3s without Traefik
```curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE="644" INSTALL_K3S_EXEC="--disable traefik" sh -s -```

Add a node to existing cluster
```curl -sfL https://get.k3s.io | K3S_URL=https://192.168.0.2:6443 K3S_TOKEN=#{token} sh -s -```