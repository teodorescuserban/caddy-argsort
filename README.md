# caddy-argsort

![Build and test](https://github.com/teodorescuserban/caddy-argsort/actions/workflows/test.yml/badge.svg)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/teodorescuserban/caddy-argsort)
[![Go Report Card](https://goreportcard.com/badge/github.com/teodorescuserban/caddy-argsort)](https://goreportcard.com/report/github.com/teodorescuserban/caddy-argsort)

This is a caddy plugin. Works with caddy 2.
Sort the request query arguments. Optionally case insensitive.

## Usage

### Set the module order

You will need to specify the execution order of this module in your caddyfile. This is done in the global options block.

```caddyfile
{
    ...
    order argsort before header
    ...
}
```

### Simple usage

Once the order has been set in the global options block, use `argsort` in any server block

```caddyfile
{
    order argsort before header
}

:8881 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
    argsort
}
```

### Optional case insensitive usage

Once the order has been set in the global options block, use `argsort lowecase` in any server block

```caddyfile
{
    order argsort before header
}

:8881 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
    argsort lowercase
}
```

### Forward the normalized request to an upstream

Once the order has been set in the global options block, you ensure query arguments sorting for an upstream server

```caddyfile
{
    order argsort before header
}

:8882 {
    argsort
    reverse_proxy localhost:8883
}

:8883 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
}
```
