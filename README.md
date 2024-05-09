# caddy-argsort

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/teodorescuserban/caddy-argsort)
[![Go Report Card](https://goreportcard.com/badge/github.com/teodorescuserban/caddy-argsort)](https://goreportcard.com/report/github.com/teodorescuserban/caddy-argsort)

This is a caddy plugin. Works with caddy 2.
Sort the request query arguments.

## Usage

Caddyfile:

```caddyfile
# Add this block in top-level settings:
{
    ...
    order argsort before rewrite
    ...
}

# use argsort keyword in any server
:8881 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
    argsort
}

# or ensure query arguments sorting then passing it to another server
:8882 {
    argsort
    reverse_proxy localhost:8883
}

:8883 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
}
```
