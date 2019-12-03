
<h1 align="center">Prometheus-CachetHQ</h1>

[![CircleCI](https://circleci.com/gh/oxyno-zeta/prometheus-cachethq/tree/master.svg?style=svg)](https://circleci.com/gh/oxyno-zeta/prometheus-cachethq/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/oxyno-zeta/prometheus-cachethq)](https://goreportcard.com/report/github.com/oxyno-zeta/prometheus-cachethq) [![GolangCI](https://golangci.com/badges/github.com/oxyno-zeta/prometheus-cachethq.svg)](https://golangci.com) [![Coverage Status](https://coveralls.io/repos/github/oxyno-zeta/prometheus-cachethq/badge.svg?branch=master)](https://coveralls.io/github/oxyno-zeta/prometheus-cachethq?branch=master) ![Docker Pulls](https://img.shields.io/docker/pulls/oxynozeta/prometheus-cachethq.svg) [![GitHub license](https://img.shields.io/github/license/oxyno-zeta/prometheus-cachethq)](https://github.com/oxyno-zeta/prometheus-cachethq/blob/master/LICENSE) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/oxyno-zeta/prometheus-cachethq)

Prometheus alerts to CachetHQ

- [Features](#features)
- [Configuration](#configuration)
- [Setup](#setup)
  - [Prometheus Alertmanager](#prometheus-alertmanager)
  - [Deploy](#deploy)
    - [Configuration](#configuration-1)
    - [Kubernetes - Helm](#kubernetes---helm)
    - [Docker](#docker)
- [Thanks](#thanks)
- [Author](#author)
- [License](#license)

## Features

- Filter Prometheus alerts by name or labels
- Change CachetHQ component status
- Allow to create incident for component
- Manage resolved alert for component or incident

## Configuration

See here: [Configuration](./docs/configuration.md)

## Setup

### Prometheus Alertmanager

Just put a new receiver in your alertmanager configuration:

```yaml
route:
  ...
  receivers:
  - name: cachethq-receiver
    webhook_configs:
    - url: http://prometheus-cachet-domain:8080/prometheus/wehbook
      send_resolved: true
```

Add also a new route to send alert to prometheus-cachethq:

```yaml
route:
  ...
    routes:
    - receiver: cachethq-receiver
      continue: true
      # match: ...
```

### Deploy

#### Configuration

See configuration values [here](./docs/configuration.md)

#### Kubernetes - Helm

A helm chart have been created to deploy this in a Kubernetes cluster.

You can find it here: [https://github.com/oxyno-zeta/helm-charts/tree/master/stable/prometheus-cachethq](https://github.com/oxyno-zeta/helm-charts/tree/master/stable/prometheus-cachethq)

#### Docker

First, write the configuration file in a config folder. That one will be mounted.

Run this command:

```shell
docker run -d --name prometheus-cachethq -p 8080:8080 -p 9090:9090 -v $PWD/config:/config oxynozeta/prometheus-cachethq
```

## Thanks

- My wife BH to support me doing this

## Author

- Oxyno-zeta (Havrileck Alexandre)

## License

Apache 2.0 (See in LICENSE)
