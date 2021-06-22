# Bj38_cli Service

This is the Bj38_cli service

Generated with

```
micro new --namespace=go.micro --type=web bj38_cli
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.bj38_cli
- Type: web
- Alias: bj38_cli

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./bj38_cli-web
```

Build a docker image
```
make docker
```