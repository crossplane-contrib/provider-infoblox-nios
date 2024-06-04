# Infoblox NIOS Crossplane Provider

`provider-infoblox-nios` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for
Infoblox IPAM.

Please note this provider is in early development and is looking for maintainers.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-infoblox-nios):

```shell
up ctp provider install crossplane-contrib/provider-infoblox-nios:v0.2.0
```

Alternatively, you can use declarative installation:

```shell
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-contrib-provider-infoblox-nios
spec:
  package: xpkg.upbound.io/crossplane-contrib/provider-infoblox-nios:v0.2.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-infoblox-nios).

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-infoblox-nios/issues).
