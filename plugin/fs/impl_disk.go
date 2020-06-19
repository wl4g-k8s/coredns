package fs

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Disk implements the fs.FileSystem interface and uses (local) disk.
type Disk struct{}

// Open implements fs.FileSystem.
func (d Disk) Open(filename string) (http.File, error) { return os.Open(filename) }

// ReadFile implements fs.FileSystem.
func (d Disk) ReadFile(filename string) ([]byte, error) { return ioutil.ReadFile(filename) }

// WriteFile implements fs.FileSystem.
func (d Disk) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
