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
  code_backup   Create a backup of directory
  db_backup     Execute the mysqldump
  explain_query Execute the EXPLAIN and EXPLAIN ANALYZE
  help          Help about any command
  show_indexes  Show indexes for all tables
  slowlog       Enable/Disable slowlog
  table_rows    Outputs the number of records in all tables
  version       Show version

Flags:
  -h, --help     help for isu
  -t, --toggle   Help message for toggle

Use "isu [command] --help" for more information about a command.

```

```console
$ isu db_backup --help
Execute the mysqldump

Usage:
  isu db_backup [flags]

Flags:
  -c, --compress            Compress file with gzip command
      --dbhost string       Database host (default "127.0.0.1")
      --dbname string       Database name
      --dbpass string       Database password
      --dbport int          Database port (default 3306)
      --dbsock string       Database socket
      --dbuser string       Database user (default "root")
  -h, --help                help for db_backup
  -d, --output-dir string   Output directory (default "./")
      --timeout duration    Timeout (default 10m0s)
```

```console
$ isu code_backup --help
Create a backup of directory

Usage:
  isu code_backup [flags]

Flags:
  -d, --dest-dir string    Destination directory (required) (default "./")
      --excludes string    Do not process files or directories that match the specified pattern (comma separated) (default ".git")
  -h, --help               help for code_backup
  -s, --src-dir string     Source directory (required)
      --timeout duration   Timeout (default 10m0s)
```

```console
$ isu table_rows --help
Outputs the number of records in all tables

Only support MySQL

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

```console
$ isu show_indexes --help
Show indexes for all tables

Only support MySQL

Usage:
  isu show_indexes [flags]

Flags:
      --dbhost string      Database host (default "127.0.0.1")
      --dbname string      Database name
      --dbpass string      Database password
      --dbport int         Database port (default 3306)
      --dbsock string      Database socket
      --dbuser string      Database user (default "root")
  -h, --help               help for show_indexes
  -r, --reverse            Sort results in reverse order
      --timeout duration   Timeout (default 10m0s)
```

```console
$ isu explain_query --help
Execute the EXPLAIN and EXPLAIN ANALYZE

Only support MySQL

Usage:
  isu explain_query [flags]

Flags:
      --dbhost string      Database host (default "127.0.0.1")
      --dbname string      Database name
      --dbpass string      Database password
      --dbport int         Database port (default 3306)
      --dbsock string      Database socket
      --dbuser string      Database user (default "root")
  -h, --help               help for explain_query
      --query string       SQL (Read from stdin when omitted)
      --timeout duration   Timeout (default 10m0s)
```

```console
$ isu slowlog on --help
Enable slowlog

Only support MySQL

Usage:
  isu slowlog on [flags]

Flags:
  -h, --help                    help for on
      --long-query-time float   Set long_query_time (default 0.1)

Global Flags:
      --dbhost string         Database host (default "localhost")
      --dbpass string         Database password
      --dbport int            Database port (default 3306)
      --dbsock string         Database socket
      --dbuser string         Database user (default "root")
      --persist SET PERSIST   Use SET PERSIST

$ isu slowlog off --help
Disable slowlog

Only support MySQL

Usage:
  isu slowlog off [flags]

Flags:
  -h, --help   help for off

Global Flags:
      --dbhost string         Database host (default "localhost")
      --dbpass string         Database password
      --dbport int            Database port (default 3306)
      --dbsock string         Database socket
      --dbuser string         Database user (default "root")
      --persist SET PERSIST   Use SET PERSIST
```
