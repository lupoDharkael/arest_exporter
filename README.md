# aREST exporter
Exporter which retrieves stats from [aREST](https://github.com/marcoschwartz/aREST) and exports them via HTTP for Prometheus consumption.

In the actual version it only exports the variables (name and value) and the hardware of the target.

## Getting Started

To run it:

```bash
./arest_exporter [flags]
```

You must define the targets, you can do it via flag or from a file.

Flag (with CSV formated values):
```bash
./arest_exporter -config.targets="192.168.0.100,192.168.0.190" [flags]
```
File:
```bash
./arest_exporter -config.file="path/to/file.csv" [flags]
```

Help on flags:

```bash
./arest_exporter --help`
```

### General information
- Default listen address :9009
- If you define both (from file and from tags), the file will overwrite the configuration.

## TODO
- Implement loging with Logrus
- Improve help format
- General clean up
