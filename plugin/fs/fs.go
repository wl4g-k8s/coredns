package fs

import (
	"net/http"
)

// A FileSystem implements access to named files. See http.FileSystem for more documentation.
// In CoreDNS the interface is extended to have more methods that are need for its operation.
type FileSystem interface {
	http.FileSystem
	ReadFile(filename string) ([]byte, error)    // See ioutil.ReadFile.
	SetOption(name string, value []string) error // SetOption is used in the setup.
}
