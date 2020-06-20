package fs

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy"
)

func TestSetupChaos(t *testing.T) {
	tests := []struct {
		input      string
		shouldErr  bool
		expectedFS FileSystem
	}{
		{
			`fs disk /`, false, &Disk{},
		},
		{
			`fs disk / {
				disk_ro
			}`, false, &Disk{ro: true},
		},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		err := setup(c)
		if err != nil {
			if !test.shouldErr {
				t.Errorf("Test %d: Expected no error but found none for input %s. Error was: %v", i, test.input, err)
				continue
			}
		}

		if test.shouldErr {
			t.Errorf("Test %d: Expected no error but found %s for input %s", i, err, test.input)
			continue
		}

		ffs := Registry.Lookup("/")
		if fmt.Sprintf("%T", ffs) != fmt.Sprintf("%T", test.expectedFS) {
			t.Errorf("Test %d: Expected type %T, got %T", i, test.expectedFS, ffs)
		}
	}
}
