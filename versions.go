package mod_ref_a

import "github.com/winlinvip/mod_minor_versions"

func Version() string {
	return "a/" + mod_minor_versions.Version()
}