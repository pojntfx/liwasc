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

🚧 This project is a work-in-progress! Instructions will be added as soon as it is usable. 🚧

## License

liwasc (c) 2021 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
