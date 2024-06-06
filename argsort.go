package argsort

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("argsort", parseCaddyfile)
}

// Middleware implements an HTTP handler that
// reorders the query arguments.
type Middleware struct {
	Lowercase bool `json:"lowercase,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.argsort",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// url.Values.Encode() is doing the sort for us
	if m.Lowercase {
		values := url.Values{}
		for k, s := range r.URL.Query() {
			for _, v := range s {
				values.Add(strings.ToLower(k), v)
			}
		}
		r.URL.RawQuery = values.Encode()
	} else {
		r.URL.RawQuery = r.URL.Query().Encode()
	}

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
// argsort [lower]
func (a *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// Set default value for Lower
	a.Lowercase = false

	for d.Next() {
		if d.NextArg() {
			if d.Val() == "lowercase" {
				a.Lowercase = true
			} else {
				return d.ArgErr()
			}
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
