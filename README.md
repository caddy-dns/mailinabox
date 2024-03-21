[Mail-In-A-Box](https://mailinabox.email/) custom DNS API module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Mail-In-A-Box](https://mailinabox.email/).

**NOTE:** This module only supports the one zone that the admin custom dns api is hosted on.

## Caddy Mail-In-A-Box

```
dns.providers.mailinabox
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "mailinabox",
				"api_url": "https://[your main-in-a-box domain name]/admin/dns/custom"
                "email_address": "{$MAIB_EMAIL}"
                "password": "{$MAIB_PASS}"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns mailinabox {
        api_url https://[your main-in-a-box domain name]/admin/dns/custom
        email_address {$MIAB_EMAIL}
        password {$MIAB_PASS}
    }
}
```

```
# wild card
*.[your-root-domain] {
	tls {
		dns mailinabox {
			api_url https://[your box domain name]/admin/dns/custom
			email_address {$MIAB_EMAIL}
			password {$MIAB_PASS}
		}
	}

	@subdomain host subdomain.[your-root-domain]
	handle @subdomain {
        response "Hello on subdomain"
	}
}
```
