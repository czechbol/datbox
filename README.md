# datbox
![GitHub](https://img.shields.io/github/license/czechbol/datbox?style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/czechbol/datbox?style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/czechbol/datbox/go.yml?style=for-the-badge)

Převádí Otevřená data ze seznamu Datových schránek z XML na jiné formáty

## Instalace
```sh
go install github.com/czechbol/datbox@latest
```

## Použití
```
❯ datbox -h
Transforms XML data from https://www.mojedatovaschranka.cz/sds/welcome?part=opendata to other formats

Drops optional parameters when unused to save on disk space.

It can also read from the STDIN stream so you can pipe data to into it.

Usage:
  datbox <path/to/file.xml> [flags]

Flags:
  -f, --format string        Output format. Allows: "json","yaml","xml" (default "json")
  -h, --help                 help for datbox
  -o, --output-file string   Output file. Otherwise the results will be printed to stdout.

❯ datbox -f json data.xml -o output.json


❯ datbox -f yaml data.xml > output.yaml


❯ cat data.xml | datbox -f json -o output.json


❯ cat data.xml | datbox -f yaml > output.yaml
```
