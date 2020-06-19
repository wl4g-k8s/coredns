package fs

import (
	"fmt"

	"github.com/coredns/coredns/plugin"

	"github.com/caddyserver/caddy"
)

func init() {
	plugin.Register("fs", setup)
}

func setup(c *caddy.Controller) error {
	for c.Next() {
		args := c.RemainingArgs()
		if len(args) != 2 {
			return plugin.Error("fs", fmt.Errorf("need a type and mountpoint"))
		}
		ffs, ok := available[args[0]]
		if !ok {
			return plugin.Error("fs", fmt.Errorf("%q is not a valid filesystem type", args[0]))
		}

		Registry.register(args[1], ffs)
	}

	return nil
}
