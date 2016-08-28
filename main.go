package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/cst05001/finddeprpm/lib"
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

	for _, rpm := range(lib.FindDepRPM(path)) {
		fmt.Printf("%s\t%s\n", rpm.Name, rpm.Version)
	}
}
