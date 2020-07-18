# isu

ISUCON utilities

## Installation

Download from https://github.com/tkuchiki/isu

## Usage

```console
$ isu --help
ISUCON utility

Usage:
  isu [command]

Available Commands:
  help        Help about any command
  table_rows  Outputs the number of records in all tables.
  version     Show version

Flags:
  -h, --help     help for isu
  -t, --toggle   Help message for toggle

Use "isu [command] --help" for more information about a command.
```

```console
$ isu table_rows --help
Outputs the number of records in all tables.

Usage:
  isu table_rows [flags]

Flags:
      --dbhost string   Database host (default "localhost")
      --dbname string   Database name
      --dbpass string   Database password
      --dbport int      Database port (default 3306)
      --dbsock string   Database socket
      --dbuser string   Database user (default "root")
  -h, --help            help for table_rows
  -r, --reverse         Sort results in reverse order
      --sort string     Output the results in sorted order(table_name or table_rows) (default "table_name")
```
