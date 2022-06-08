<img src="https://penguin.upyun.galvincdn.com/logos/penguin_stats_logo.png"
     alt="Penguin Statistics - Logo"
     width="96px" />

# Penguin Statistics - LiveHouse

[![License](https://img.shields.io/github/license/penguin-statistics/livehouse)](https://github.com/penguin-statistics/livehouse/blob/dev/LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/penguin-statistics/livehouse)](https://github.com/penguin-statistics/livehouse/commits/dev)
[![GitHub Actions Status](https://github.com/penguin-statistics/livehouse/actions/workflows/build-release.yml/badge.svg)](https://github.com/penguin-statistics/livehouse/actions/workflows/build-release.yml)
[![go.mod Version](https://img.shields.io/github/go-mod/go-version/penguin-statistics/livehouse)](https://github.com/penguin-statistics/livehouse/blob/main/go.mod)

Live data stream for [Penguin Statistics](https://penguin-stats.io/?utm_source=github), built with Go.

## Architecture

This Go project mainly uses:

- [grpc/grpc-go](https://github.com/grpc/grpc-go)
- [protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go)
- [uber-go/fx](https://github.com/uber-go/fx)

This project communicates with its dependency services via gRPC & protocol buffers.

## Development

1. Install protobuf Go plugins: follow [**Prerequisites | Quick Start | Go | gRPC**](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

## Maintainers

This project has mainly being maintained by the following contributors (in alphabetical order):

- [AlvISsReimu](https://github.com/AlvISsReimu)
- [GalvinGao](https://github.com/GalvinGao)

> The full list of active contributors of the _Penguin Statistics_ project can be found at the [Team Members page](https://penguin-stats.io/about/members) of the website.
