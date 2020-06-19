package fs

import (
	"net/http"
	"os"
)

// A FileSystem implements access to named files. See http.FileSystem for more documentation.
// In CoreDNS the interface is extended to have more methods that are need for its operation.
type FileSystem interface {
	http.FileSystem
	ReadFile(filename string) ([]byte, error)                       // See ioutil.ReadFile.
	WriteFile(filename string, data []byte, perm os.FileMode) error // See ioutil.WriteFile.
}
