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
  -h, --help            help for datbox
  -o, --output string   Output format. Allows: "json","yaml","xml" (default "json")

❯ datbox -o yaml data.xml > output.yaml
```

Alternativní použití:
```
❯ cat data.xml | datbox -o yaml > output.yaml
```
