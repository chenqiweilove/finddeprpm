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
	STATUS_ERR_OTHER
)

const (
	colorStart = "\033[31;5m"
	colorEnd = "\033[0m"
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
		fmt.Println(err)
		if os.IsNotExist(err) {
			os.Exit(STATUS_ERR_OPEN_ROOT_PATH)
		} else {
			os.Exit(STATUS_ERR_OTHER)
		}
	}

	rpmList, missingLDList := lib.FindDepRPM(path)
	for _, rpm := range(rpmList) {
		fmt.Printf("%s\t%s\n", rpm.Name, rpm.Version)
	}
	for _, ld := range(missingLDList) {
		fmt.Fprintf(os.Stderr, "%s%s%s\n", colorStart, ld, colorEnd)
	}
}
