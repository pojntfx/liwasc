# liwasc

[![liwasc demo video](./assets/demo.webp)](https://pojntfx.github.io/liwasc/)

List, wake and scan nodes in a network.

[![hydrun CI](https://github.com/pojntfx/liwasc/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/liwasc/actions/workflows/hydrun.yaml)
[![Docker CI](https://github.com/pojntfx/liwasc/actions/workflows/docker.yaml/badge.svg)](https://github.com/pojntfx/liwasc/actions/workflows/docker.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/liwasc.svg)](https://pkg.go.dev/github.com/pojntfx/liwasc)
[![Matrix](https://img.shields.io/matrix/liwasc:matrix.org)](https://matrix.to/#/#liwasc:matrix.org?via=matrix.org)

## Overview

liwasc is a high-performance network and port scanner. It can quickly give you a overview of the nodes in your network, the services that run on them and manage their power status.

It can ...

- **Scan a network**: Using an ARP scan and the [mac2vendor](https://mac2vendor.com/) database, liwasc can list the nodes in a network, their power status, manufacturer information, IP & MAC addresses and more metadata
- **Scan a node**: Using a high-performance custom TCP and UDP port scanner, liwasc can list the ports and services of a node and provide metadata (service names, registration dates etc.) using the [Service Name and Transport Protocol Port Number Registry](https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml)
- **Power on a node**: By sending [Wake-on-LAN packets](https://en.wikipedia.org/wiki/Wake-on-LAN), liwasc can turn on nodes in a network
- **Periodically scan a network**: Using the integrated periodic scans feature, liwasc can periodically (based on a CRON syntax) scan a network and persist the results in a database
- **Give remote insight into a network**: Because liwasc is based on open web technologies, has a gRPC API and supports OpenID Connect authentication, liwasc can be securely exposed to the public internet and serve as a remote controller for a network

## Installation

### Containerized

You can get the Docker container like so:

```shell
$ docker pull pojntfx/liwasc-backend
```

### Natively

If you prefer a native installation, static binaries are also available on [GitHub releases](https://github.com/pojntfx/liwasc/releases).

You can install them like so:

```shell
$ curl -L -o /tmp/liwasc-backend https://github.com/pojntfx/liwasc/releases/download/latest/liwasc-backend.linux-$(uname -m)
$ sudo install /tmp/liwasc-backend /usr/local/bin
$ sudo setcap cap_net_raw+ep /usr/local/bin/liwasc-backend # This allows rootless execution
```

### About the Frontend

The frontend is also available on [GitHub releases](https://github.com/pojntfx/liwasc/releases) in the form of a static `.tar.gz` archive; to deploy it, simply upload it to a CDN or copy it to a web server. For most users, this shouldn't be necessary though; thanks to [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/), liwasc is a progressive web app. By simply visiting the [public deployment](https://pojntfx.github.io/liwasc/) once, it will be available for offline use whenever you need it.

## Usage

### Setting up Authentication

liwasc uses [OpenID Connect](https://en.wikipedia.org/wiki/OpenID_Connect) for authentication, which means you can use almost any authentication provider, both self-hosted and as a service, that you want to. We've created a short tutorial video which shows how to set up [Auth0](https://auth0.com/) for this purpose, but feel free to use something like [Ory](https://github.com/ory/hydra) if you prefer a self-hosted solution:

[<img src="https://img.youtube.com/vi/N3cocCOsrGw/0.jpg" width="512" alt="Setting up OpenID Connect for Internal Apps YouTube Video" title="Setting up OpenID Connect for Internal Apps YouTube Video">](https://www.youtube.com/watch?v=N3cocCOsrGw)

### Starting the Backend (Containerized)

Using Docker (or an alternative like Podman), you can easily start & configure the backend; see the [Reference](#reference) for more configuration parameters:

```shell
$ docker run \
    --name liwasc-backend \
    -d \
    --restart always \
    --net host \
    --cap-add NET_RAW \
    -v ${HOME}/.local/share/liwasc:/root/.local/share/liwasc:z \
    -e LIWASC_BACKEND_OIDCISSUER=https://pojntfx.eu.auth0.com/ \
    -e LIWASC_BACKEND_OIDCCLIENTID=q31IUpAg5Qu18eBMiBsbIvH9hhRq7WgE \
    -e LIWASC_BACKEND_DEVICENAME=eth0 \
    pojntfx/liwasc-backend
```

## License

liwasc (c) 2021 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
