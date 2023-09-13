[![Build](https://img.shields.io/github/actions/workflow/status/vshn/appcat-apiserver/test.yml?branch=master)](https://github.com/vshn/appcat-apiserver/actions?query=workflow%3ATest)
![Go version](https://img.shields.io/github/go-mod/go-version/vshn/appcat-apiserver)
![Kubernetes version](https://img.shields.io/badge/k8s-v1.24-blue)
[![Version](https://img.shields.io/github/v/release/vshn/appcat-apiserver)](https://github.com/vshn/appcat-apiserver/releases)
[![GitHub downloads](https://img.shields.io/github/downloads/vshn/appcat-apiserver/total)](https://github.com/appuio/control-api/releases)

# AppCat Apiserver

This repository has k8s tools to manage AppCat services.

Documentation: https://vshn.github.io/appcat-apiserver

## API Server

The AppCat API Server facilitates work with AppCat services.

### Capabilities

The API Server is able to manage the following:

| RESOURCE           | DESCRIPTION                                             |
| ------------------ | ------------------------------------------------------- |
| `AppCat`           | Shows active AppCat services in a cluster               |
| `PotsgreSQLBackup` | Shows available backups of an AppCat PostgreSQL service |

### Debugging in IDE
To run the API server on your local machine you need to register the IDE running instance with kind cluster.
This can be achieved with the following guide.

The `externalName` needs to be changed to your specific host IP.
When running kind on Linux you can find it with `docker inspect`.
On some docker distributions the host IP is accessible via `host.docker.internal`.
For Lima distribution the host IP is accessible via `host.lima.internal`.

```bash
# Run this command in kindev -> https://github.com/vshn/kindev
make appcat-apiserver

HOSTIP=$(docker inspect kindev-control-plane | jq '.[0].NetworkSettings.Networks.kind.Gateway')
# HOSTIP=host.docker.internal # On some docker distributions
# HOSTIP=host.lima.internal # On lima distributions

kind get kubeconfig --name kindev  > ~/.kube/config

cat <<EOF | sed -e "s/172.21.0.1/$HOSTIP/g" | kubectl apply -f -
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  name: v1.api.appcat.vshn.io
  labels:
    api: appcat
    apiserver: "true"
spec:
  version: v1
  group: api.appcat.vshn.io
  insecureSkipTLSVerify: true
  groupPriorityMinimum: 2000
  service:
    name: appcat
    namespace: default
    port: 9443
  versionPriority: 10
---
apiVersion: v1
kind: Service
metadata:
  name: appcat
  namespace: default
spec:
  ports:
  - port: 9443
    protocol: TCP
    targetPort: 9443
  type: ExternalName
  externalName: 172.21.0.1 # Change to host IP
EOF
```

After the above steps just run the API server via IDE with the following arguments.
```
apiserver --secure-port=9443 --kubeconfig=<your-home-path>/.kube/config --authentication-kubeconfig=<your-home-path>/.kube/config --authorization-kubeconfig=<your-home-path>/.kube/config --feature-gates=APIPriorityAndFairness=false
```

### Protobuf installation
Protocol Buffers (Protobuf) is a free and open-source cross-platform data format used to serialize structured data.
Kubernetes internally uses gRPC clients with protobuf serialization. APIServer objects when handled internally in K8S
need to implement protobuf interface. The implementation of the interface is done by [code-generator](https://github.com/kubernetes/code-generator). Two dependencies are required to use this tool [protoc](https://github.com/protocolbuffers/protobuf) and [protoc-gen-go](https://google.golang.org/protobuf/cmd/protoc-gen-go).

