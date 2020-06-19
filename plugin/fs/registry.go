package fs

import "path"

type registry struct {
	m map[string]FileSystem
	// mutex ?
}

// available holds the names of all FileSystems implementation.
var available = map[string]FileSystem{}

var Registry = &registry{m: map[string]FileSystem{}}

// needs to be a init() to be available during setup().
func init() {
	available = map[string]FileSystem{
		"disk": Disk{},
	}
	Registry = &registry{m: map[string]FileSystem{
		"/": Disk{},
	}}
}

// register registers a new FileSystem. Is a private method, because it should only be used
// by the fs plugin.
func (r *registry) register(mountpoint string, fs FileSystem) {
	r.m[mountpoint] = fs
}

// method(s) used by plugins

// Lookup looks up the FileSystem implementation to for a specific mount point. The most
// specific match is returned.
func (r *registry) Lookup(mountpoint string) FileSystem {
	if ffs, ok := r.m[mountpoint]; ok {
		return ffs
	}
	x := mountpoint
	for x = path.Dir(x); x != "."; x = path.Dir(x) {
		if ffs, ok := r.m[x]; ok {
			return ffs
		}
	}

	return Disk{}
}
