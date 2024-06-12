// Package argsort provides a Caddy module that optionally lowercase and then sort the query arguments.
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
	caddy.RegisterModule(Argsort{})
	httpcaddyfile.RegisterHandlerDirective("argsort", parseCaddyfile)
	httpcaddyfile.RegisterDirectiveOrder("argsort", "before", "header")
}

// Argsort sort the query arguments after optionally lowercasing them.
//
// Syntax:
//
//	argsort [lowercase]
type Argsort struct {
	// Lowercase the query arguments before sorting them.
	Lowercase bool `json:"lowercase,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Argsort) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.argsort",
		New: func() caddy.Module { return new(Argsort) },
	}
}

// Provision implements caddy.Provisioner.
func (a *Argsort) Provision(ctx caddy.Context) error {
	return nil
}

// Validate implements caddy.Validator.
func (a *Argsort) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (a Argsort) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// url.Values.Encode() is doing the sort for us
	if a.Lowercase {
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
func (a *Argsort) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
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

// parseCaddyfile sets up the handler from Caddyfile tokens. Syntax:
//
//	argsort [lowecase]
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var a Argsort
	err := a.UnmarshalCaddyfile(h.Dispenser)
	return a, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Argsort)(nil)
	_ caddy.Validator             = (*Argsort)(nil)
	_ caddyhttp.MiddlewareHandler = (*Argsort)(nil)
	_ caddyfile.Unmarshaler       = (*Argsort)(nil)
)
