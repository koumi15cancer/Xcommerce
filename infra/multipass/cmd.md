multipass launch --name k3sMasterNode --cpus 2 --memory 4G --disk 20G
multipass launch --name k3sWorkerNode1 --cpus 2 --memory 4G --disk 20G
multipass launch --name k3sWorkerNode2 --cpus 2 --memory 4G --disk 20G

# Step 2:
multipass exec k3sMasterNode -- /bin/bash -c "curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE='644' sh -"

K3S_MASTERNODE_IP="https://$(multipass info k3sMasterNode | grep 'IPv4' | awk '{print $2}'):6443"

TOKEN="$(multipass exec k3sMasterNode -- /bin/bash -c 'sudo cat /var/lib/rancher/k3s/server/node-token')"

# Step 3:
multipass exec k3sWorkerNode1 -- /bin/bash -c "curl -sfL https://get.k3s.io | K3S_TOKEN=${TOKEN} K3S_URL=${K3S_MASTERNODE_IP} sh -"
multipass exec k3sWorkerNode2 -- /bin/bash -c "curl -sfL https://get.k3s.io | K3S_TOKEN=${TOKEN} K3S_URL=${K3S_MASTERNODE_IP} sh -"
