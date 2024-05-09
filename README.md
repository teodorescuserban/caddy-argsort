# caddy-argsort

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
