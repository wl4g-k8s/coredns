package fs

import (
	"net/http"
	"testing"
)

type noop int

func (n *noop) Open(name string) (http.File, error)         { return nil, nil }
func (n *noop) ReadFile(name string) ([]byte, error)        { return nil, nil }
func (n *noop) SetOption(name string, value []string) error { return nil }

type fs1 struct{ *noop }
type fs2 struct{ *noop }

func TestLookup(t *testing.T) {
	Registry.register("/home/coredns", fs1{})
	Registry.register("/home", fs2{})

	f := Registry.Lookup(".")
	if x, ok := f.(*Disk); !ok {
		t.Errorf("Lookup for %s, should result in %T, got %T", ".", &Disk{}, x)
	}
	f = Registry.Lookup("/")
	if x, ok := f.(*Disk); !ok {
		t.Errorf("Lookup for %s, should result in %T, got %T", "/", &Disk{}, x)
	}
	f = Registry.Lookup("/home")
	if x, ok := f.(fs2); !ok {
		t.Errorf("Lookup for %s, should result in %T, got %T", "/home", fs2{}, x)
	}
	f = Registry.Lookup("/home/coredns")
	if x, ok := f.(fs1); !ok {
		t.Errorf("Lookup for %s, should result in %T, got %T", "/home/coredns", fs1{}, x)
	}
}
