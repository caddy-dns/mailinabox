package mailinabox

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/mailinabox"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *mailinabox.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.mailinabox",
		New: func() caddy.Module { return &Provider{new(mailinabox.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIURL = caddy.NewReplacer().ReplaceAll(p.Provider.APIURL, "")
	p.Provider.EmailAddress = caddy.NewReplacer().ReplaceAll(p.Provider.EmailAddress, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	p.Provider.TOTPSecret = caddy.NewReplacer().ReplaceAll(p.Provider.TOTPSecret, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	mailinabox {
//	    api_url <api_url>
//	    email_address <email_address>
//	    password <password>
//	    totp_secret <totp_secret>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_url":
				if p.Provider.APIURL != "" {
					return d.Err("API URL already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIURL = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "email_address":
				if p.Provider.EmailAddress != "" {
					return d.Err("email address already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.EmailAddress = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("password already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.Password = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
                        case "totp_secret":
                                if p.Provider.TOTPSecret != "" {
                                        return d.Err("TOTP secret already set")
                                }
                                if !d.NextArg() {
                                        return d.ArgErr()
                                }
                                p.Provider.TOTPSecret = d.Val()
                                if d.NextArg() {
                                        return d.ArgErr()
                                }
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIURL == "" {
		return d.Err("missing API URL")
	}
	if p.Provider.EmailAddress == "" {
		return d.Err("missing email address")
	}
	if p.Provider.Password == "" {
		return d.Err("missing password")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
