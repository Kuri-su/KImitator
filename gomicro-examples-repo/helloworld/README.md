# Helloworld Service

This is the Helloworld service

Generated with

```
micro new github.com/micro/services/helloworld --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.helloworld
- Type: srv
- Alias: helloworld

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./helloworld-srv
```

Build a docker image
```
make docker
```
