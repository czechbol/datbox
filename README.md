# datbox
Převádí Otevřená data ze seznamu Datových schránek z XML na jiné formáty

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