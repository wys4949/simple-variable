package simple_variable

import (
	"io/ioutil"
	"os"
	_ "github.com/wys4949/simple-variable/init"
	"path/filepath"
	"strings"
)

type Read struct {
}
func (s *Read)GetADSDirectory() (path string,err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		path = ""
		return
	}
	path = strings.Replace(dir, "\\", "/", -1)
	return
}

func (s *Read)All(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
