# CacheUp

Simple CLI tool and go module for caching content on disk.


## CLI Installation
```bash
go install github.com/zcag/cacheup@latest
```
## Module Installation
```bash
go get github.com/zcag/cacheup@latest
```

## Module Usage

```go
import ( "github.com/zcag/cacheup" )

func foo() {
    cacheup.Write("~/cache/file.json", "contents")
    cacheup.IsValid("~/cache/file.json", "30m")
    cacheup.Read("~/cache/file.json", "30m", "")
    cacheup.Read("~/cache/file.json", "30m", "curl http://example.com/fetch_fresh_file.json")
}

```


## CLI Usage

```bash
cacheup --help

Usage:
  cacheup [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  read        Read cached value
  valid       Check if a cache file is valid
  write       Write stdin to cache

Flags:
  -t, --cache-max-age string   max age of cache file (default: 1h), unit can be s/m/h/d for second/minute/hour/day (default "1h")
  -f, --cache-path string      custom location for cache path for parent directory or file. directory paths should end with '/'
                                        (default: $XDG_CACHE_HOME/cacheup/<name>)
  -h, --help                   help for cacheup

Use "cacheup [command] --help" for more information about a command.
```

### Example snippets

```bash
cacheup read <name>
cacheup read <name> -f ~/custom/cache/
cacheup read -f ~/custom/cache/file.json

# Set command with -c to refresh if cache is invalid
cacheup read <name> -c "curl XX"
cacheup read <name> -t 30m -c "~/some/script.sh"

./heavy_command.sh | cacheup write <name>
curl http://example.com | cacheup -f ~/custom/cache/folder/ write <name>
curl http://example.com | cacheup -f ~/custom/cache/file write

cacheup valid <name>
cacheup read <name> -f ~/custom/cache/ -t 30m
cacheup read -f ~/custom/cache/file.json
```
