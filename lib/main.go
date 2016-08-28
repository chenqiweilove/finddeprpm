package lib

import (
	"os"
	"fmt"
	"path/filepath"
)

var RPM_MAP map[string]*RPM

func FindDepRPM(path string) map[string]*RPM {

	RPM_MAP = make(map[string]*RPM)

	// go go go
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			ldd(path)
		}

		return nil
	})
	return RPM_MAP
}
