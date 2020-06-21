package fs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Disk implements the fs.FileSystem interface and uses (local) disk.
type Disk struct {
	ro bool // readonly
}

// Open implements fs.FileSystem.
func (d *Disk) Open(filename string) (http.File, error) { return os.Open(filename) }

// Open implements fs.FileSystem.
func (d *Disk) ReadFile(filename string) ([]byte, error) { return ioutil.ReadFile(filename) }

func (d *Disk) setOption(opt string, value []string) error {
	switch opt {
	case "disk_ro":
		if len(value) > 0 {
			return fmt.Errorf("no arguments allowed for: %s", opt)
		}
		d.ro = true
	}
	return nil
}
