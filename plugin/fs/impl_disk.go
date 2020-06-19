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

// Open implements fs.FileSystem.
func (d Disk) ReadFile(filename string) ([]byte, error) { return ioutil.ReadFile(filename) }
