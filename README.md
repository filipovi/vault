# vault

## Stateless secure password generator in go

[Scrypt](https://en.wikipedia.org/wiki/Scrypt)

[Master Password](https://en.wikipedia.org/wiki/Master_Password)

## TODO

- goji -> chi @DONE
- rename generator in master_password
- add master_password tests
- run in minikube

## Vegeta

```echo "POST http://0.0.0.0:3000/password" | vegeta attack -duration=5s -body=body.json -http2 -rate=5000 | tee results.bin | vegeta report```

1. Goji

```
    Requests      [total, rate]            25000, 5000.20
    Duration      [total, attack, wait]    5.000605241s, 4.9997995s, 805.741µs
    Latencies     [mean, 50, 95, 99, max]  14.086706ms, 2.84644ms, 89.886483ms, 195.154983ms, 285.813935ms
    Bytes In      [total, mean]            875000, 35.00
    Bytes Out     [total, mean]            2375000, 95.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:25000  
    Error Set:
```

1. Chi

```
    Requests      [total, rate]            25000, 4999.88
    Duration      [total, attack, wait]    5.001429649s, 5.000119764s, 1.309885ms
    Latencies     [mean, 50, 95, 99, max]  10.524365ms, 829.808µs, 90.547445ms, 145.97906ms, 1.058563681s
    Bytes In      [total, mean]            875000, 35.00
    Bytes Out     [total, mean]            2375000, 95.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:25000  
    Error Set:
```

## Minikube on Fedora

### Nettoyage

    minikube stop
    minikube delete
    rm -rf .kube
    rm -rf .minikube

    docker ps -aq |  xargs -r docker rm -v -f
    docker rmi $(docker images -q) -f

### Installation

    sudo dnf install kubectl
    sudo dnf install libvirt-daemon-kvm qemu-kvm\n
    newgrp libvirt
    curl -Lo docker-machine-driver-kvm2 https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-kvm2
    sudo cp docker-machine-driver-kvm2 /usr/local/bin
    rm docker-machine-driver-kvm2

    minikube start --logtostderr  --vm-driver kvm2
    minikube status
    kubectl cluster-info
    kubectl get nodes
    minikube dashboard

    .zshrc:
    ...
    # minikube
    export MINIKUBE_WANTUPDATENOTIFICATION=false
    export MINIKUBE_WANTREPORTERRORPROMPT=false
    export MINIKUBE_HOME=$HOME
    export CHANGE_MINIKUBE_NONE_USER=true
    export KUBECONFIG=$HOME/.kube/config

    eval $(minikube docker-env)

### Application in Minikube

    eval $(minikube docker-env)
    docker build . --tag filipovi/vault:v0.9
    kubectl run vault --image=filipovi/vault:v0.9  --port=3000
    minikube service --logtostderr vault
    kubectl get deployments
    kubectl get service

### Add a dummy interface for the docker rpc container to access the consul container

    ip link show type dummy
    sudo ip link add dummy0 type dummy
    sudo ip link set dev dummy0 up
    ip link show type dummy
    sudo ip addr add 169.254.1.1/32 dev dummy0
    sudo ip link set dev dummy0 up
    ip addr show dev dummy0

In /etc/systemd/network/dummy0.netdev:

    [NetDev]
    Name=dummy0
    Kind=dummy

In /etc/systemd/network/dummy0.network:

    [Match]
    Name=dummy0

    [Network]
    Address=169.254.1.1/32

### Remove the application from Minikube

    delete service vault
    kubectl delete deployment vault

