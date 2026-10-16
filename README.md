<p align="center">
  <img alt="FuseScope" src="docs/brand/fusescope-logo-universal.svg" width="440">
</p>

<h1 align="center" style="border-bottom: none">FuseScope OpenTelemetry Collector</h1>

<p align="center">
  <a href="https://golang.org"><img alt="Go Version" src="https://img.shields.io/badge/Go-1.25+-0f5fb8.svg"></a>
  <a href="LICENSE"><img alt="License: AGPL v3" src="https://img.shields.io/badge/License-AGPL%20v3-950fb8.svg"></a>
</p>

<p align="center">The ingestion layer of <a href="https://github.com/fuseone-io/fusescope">FuseScope</a>: an OpenTelemetry Collector distribution that receives logs, metrics and traces and writes them to ClickHouse, plus the ClickHouse schema migrator.</p>

## Overview

This repository builds the OpenTelemetry Collector distribution that powers FuseScope.

It combines a curated set of upstream OpenTelemetry Collector and Collector-Contrib components with FuseScope's own components (inherited from SigNoz): receivers, processors, exporters, connectors, and an extension.

Like any Collector, where your telemetry goes is decided by the exporters in your pipeline. The FuseScope exporters write to ClickHouse, the store behind FuseScope. You can run the collector from a static config file, or manage it remotely over OpAMP (Open Agent Management Protocol).

## About this fork

This is a hard fork of [SigNoz/signoz-otel-collector](https://github.com/SigNoz/signoz-otel-collector), maintained by [FuseOne](https://github.com/fuseone-io) for FuseScope. It owns the ClickHouse schema that FuseScope reads. Component type names used in collector configuration (for example `signozlogspipeline` or `clickhousetraces`) are kept unchanged for compatibility.

## Why this collector?

Here are a few reasons teams run this collector:

- **Built on [OpenTelemetry](https://opentelemetry.io/)**: works with standard OTLP sources and the upstream Collector components you already know.
- **Everything FuseScope needs in one build**: upstream Collector and Contrib components alongside the FuseScope receivers, processors, exporters, connectors, and an extension inherited from SigNoz.
- **Fast storage**: the FuseScope exporters write to ClickHouse, an extremely fast and highly optimized store for observability data.
- **Runs your way**: from a static config file, or managed remotely over OpAMP.

## Getting Started

The easiest way to run FuseScope, including this collector, is with the [FuseScope Helm charts](https://github.com/fuseone-io/fusescope-charts). You can also build and run the collector from source for local development or custom images.

Prerequisites: [Go 1.25+](https://golang.org/dl/).

Build the collector:

```bash
make build
```

Run it against a config file:

```bash
make run
# make run uses ./config/default-config.yaml. To pass your own:
go run ./cmd/signozotelcollector --config ./config/default-config.yaml
```

On success the collector starts and accepts telemetry on the OTLP (OpenTelemetry Protocol) endpoints declared by its receivers, HTTP on `4318` and gRPC on `4317` by default. See [`example/example-config.yaml`](example/example-config.yaml) for a complete pipeline.

Build the container image:

```bash
make build-collector   # → ghcr.io/fuseone-io/fusescope-otel-collector
```

## Configuration

The collector reads a standard OpenTelemetry Collector configuration: `receivers`, `processors`, `exporters`, `connectors`, `extensions`, and a `service` block of pipelines. See [`config/default-config.yaml`](config/default-config.yaml) and [`example/example-config.yaml`](example/example-config.yaml).

You can supply configuration in one of two ways:

- **Static file**: pass `--config <path-to-config>`. The collector reads the file at startup.
- **Managed (OpAMP)**: pass `--manager-config <path-to-manager-config>` (see [`config/default-manager-config.yaml`](config/default-manager-config.yaml)). The collector fetches its configuration from a remote OpAMP server, and copies the active config to `--copy-path` (default `/etc/otel/signozcol-config.yaml`).

## CLI reference

```
fusescope-otel-collector [flags]

Flags:
  --config string           Path to the collector configuration file
  --manager-config string   Path to the OpAMP agent manager configuration (enables managed mode)
  --copy-path string        Where the active config is written in managed mode
                            (default "/etc/otel/signozcol-config.yaml")
  --feature-gates string    Comma-delimited feature gate identifiers; prefix with '-' to disable,
                            '+' or no prefix to enable

Commands:
  migrate                   Run schema migrations for the telemetry store
```

Run `fusescope-otel-collector <command> --help` for a command's subcommands and flags.

## Components

This distribution bundles upstream OpenTelemetry Collector and Collector-Contrib components together with FuseScope-specific receivers, processors, exporters, connectors, and an extension. See [`components/components.go`](components/components.go) for the full list of registered components.

## Development

```bash
make build           # build the collector
make test            # go test with the race detector and coverage
make fmt             # gofmt
make lint            # golangci-lint
make test-and-lint   # test + fmt + lint (default target)
```

## Contributing

Open an [issue](https://github.com/fuseone-io/fusescope-otel-collector/issues) or a pull request to get started.

## License

Licensed under the [GNU Affero General Public License v3.0](LICENSE). Originally developed by SigNoz Inc.; modified by FuseOne — see [NOTICE](NOTICE).
