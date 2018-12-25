[![Build Status](https://travis-ci.org/inwinstack/ipvs-elector.svg?branch=master)](https://travis-ci.org/inwinstack/ipvs-elector)
# IPVS Elector
Implementing the IPVS ARP leader election in Kubernetes, if the node leading, it will be sending replies in response to received ARP requests that resolve local target IP addresses. The IPVS Elector also implements a failover mechanism so that a different node can take over should the current leader node fail for some reason.

## Building from Source
Clone repo into your go path under `$GOPATH/src`:
```sh
$ git clone https://github.com/inwinstack/ipvs-elector.git $GOPATH/src/github.com/inwinstack/ipvs-elector
$ cd $GOPATH/src/github.com/inwinstack/ipvs-elector
$ dep ensure
$ make
```

## Running
To see the app in action, run the following three commands in separate terminals:
```sh
# terminal 1
$ POD_NAME=test-1 POD_NAMESPACE=default go run cmd/main.go --kubeconfig $HOME/.kube/config --logtostderr -v=2

# terminal 2
$ POD_NAME=test-2 POD_NAMESPACE=default go run cmd/main.go --kubeconfig $HOME/.kube/config --logtostderr -v=2

# terminal 3
$ POD_NAME=test-3 POD_NAMESPACE=default go run cmd/main.go --kubeconfig $HOME/.kube/config --logtostderr -v=2
```
> Now kill a terminal to see the changes.

## Deploying elector into Kubernetes cluster
Run the following command to deploy the IPVS Elector:
```sh
$ kubectl apply -f deploy/
$ kubectl -n kube-system get po -l k8s-app=ipvs-elector -o wide
```