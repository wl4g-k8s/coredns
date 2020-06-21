package fs

import (
	"fmt"
	"strings"

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
		typ := args[0]
		mount := args[1]
		ffs, ok := available[typ]
		if !ok {
			return plugin.Error("fs", fmt.Errorf("%q is not a valid filesystem type", typ))
		}
		// if there is a block, only check <type>_ options, error on anything else
		for c.NextBlock() {
			opt := c.Val()
			if !strings.HasPrefix(opt, typ+"_") {
				return fmt.Errorf("looking for %s_ options, found: %s", typ, opt)
			}
			if err := ffs.setOption(opt, c.RemainingArgs()); err != nil {
				return err
			}
		}

		Registry.register(mount, ffs)
	}

	return nil
}
