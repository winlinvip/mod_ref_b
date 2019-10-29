package mod_ref_b

import "github.com/winlinvip/mod_minor_versions"

func Version() string {
	return "a/" + mod_minor_versions.Version()
}
