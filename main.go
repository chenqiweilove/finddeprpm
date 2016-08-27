package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	STATUS_OK = iota
	STATUS_ERR_PARAMS_ERROR
	STATUS_ERR_OPEN_ROOT_PATH
	STATUS_ERR_LDD_EXECUTE
)

func usage() {
	fmt.Println("Usage:\n\t%s {DIRECTORY|FILE}\n", filepath.Base(os.Args[0]))
	os.Exit(STATUS_ERR_PARAMS_ERROR)
}

var RPM_MAP = make(map[string]*RPM)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	path := os.Args[1]

	_, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}
		os.Exit(STATUS_ERR_OPEN_ROOT_PATH)
	}

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

	for _, rpm := range(RPM_MAP) {
		fmt.Printf("%s\t%s\n", rpm.Name, rpm.Version)
	}
}