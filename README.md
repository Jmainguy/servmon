# servmon
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/servmon)](https://goreportcard.com/badge/github.com/Jmainguy/servmon)
[![Release](https://img.shields.io/github/release/Jmainguy/servmon.svg?style=flat-square)](https://github.com/Jmainguy/servmon/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/servmon/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/servmon?branch=main)

A service monitor written in go

## Usage
You will need to provide a monitor.yml of your own, customized to your needs.

You will also need to set a number of environmental variables.

```/bin/bash
export SERVMONDIR=/opt/servmon
export SLACK_TOKEN=adssadasasdsa
export SLACK_CHANNEL=C02U1PAFP8Q
servmon
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/servmon/releases)

## Install

### Homebrew

```/bin/bash
brew install Jmainguy/tap/servmon
```

### Podman
```/bin/bash
# if running podman on arm64
podman run --name servmon -d \
  -p 8080:8080 \
  -v $(pwd)/monitor.yml:/monitor.yml \
  --env SLACK_TOKEN=xoxb-1986333393-gRNl6nANyVXhSGKJGDc9QHsa \
  --env SLACK_CHANNEL=C02U1PAFP8Q \
  hub.soh.re/servmon:latest-arm64
```

## Build
```/bin/bash
export GO111MODULE=on
go build
```
