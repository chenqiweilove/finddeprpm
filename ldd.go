package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	"regexp"
)

func ldd(path string) {

	colorStart := "\033[31;5m"
	colorEnd := "\033[0m"

	re3 := regexp.MustCompile(`^\s*(\S+)\s+=>\s+(\S+)\s+(\S+)\s*$`) // libkrb5.so.3 => /usr/lib64/libkrb5.so.3 (0x00007f0a2b588000)
	re1 := regexp.MustCompile(`^\s*(\S+)\s+(\(\S+\))\s*$`) // /lib64/ld-linux-x86-64.so.2 (0x00007fc18d3d1000)
	reNotFound := regexp.MustCompile(`^\s*(\S+)\s+=>\s+not found\s*$`) // ibexpat.so.1 => not found

	cmd := exec.Command("ldd", path)
	out, err := cmd.Output()
	if err != nil { // Some regular file but cannot ldd
		return
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range(lines) {
		if reNotFound.MatchString(line) {
			group := reNotFound.FindStringSubmatch(line)
			fmt.Fprintf(os.Stderr, "%s%s%s\n", colorStart, group[1], colorEnd)
			continue
		}

		so := ""
		if re3.MatchString(line) {
			group := re3.FindStringSubmatch(line)
			so = group[2]
		} else if re1.MatchString(line) {
			group := re1.FindStringSubmatch(line)
			so = group[1]
		}
		if so != "" {
			rpm, err := findrpm(so)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if _, ok := RPM_MAP[rpm.Name]; ok {
				continue
			}
			RPM_MAP[rpm.Name] = rpm
		}
	}
}