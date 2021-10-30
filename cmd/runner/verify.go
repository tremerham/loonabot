package runner

import (
	"crypto/sha256"
	"io/ioutil"

	"github.com/kardianos/osext"
)

func Verify() error {
	exe, err := osext.Executable()
	if err != nil {
		return err
	}
	hasher := sha256.New()
	b, err := ioutil.ReadFile(exe)
	if err != nil {
		return err
	}
	hash := hasher.Sum(b)
	if hash == nil {
		return err
	}

	return nil
}
