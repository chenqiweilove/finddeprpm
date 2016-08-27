package main

import (
	"strings"
	"errors"
	"os/exec"
	"regexp"
)

var (
	ERROR_FINDRPM_REGEXP_NOT_MATCH error = errors.New("ERROR_FINDRPM_REGEXP_NOT_MATCH")
	ERROR_LD_FILE_NOT_EXIST error = errors.New("ERROR_LD_FILE_NOT_EXIST")
)

func findrpm(path string) (*RPM, error) {

	var name, version string

	reName := regexp.MustCompile(`(?s)Name\s*:\s*(\S*)`)
	reVersion := regexp.MustCompile(`(?s)Version\s*:\s*(\S*)`)

	cmd := exec.Command("rpm", "-q", "-f", "-i", path)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if !reName.Match(out) {
		return nil, ERROR_FINDRPM_REGEXP_NOT_MATCH
	} else {
		name = reName.FindStringSubmatch(string(out))[1]
	}
	if !reVersion.Match(out) {
		return nil, ERROR_FINDRPM_REGEXP_NOT_MATCH
	} else {
		version = reVersion.FindStringSubmatch(string(out))[1]
	}

	rpm := &RPM{
		Name:	strings.Trim(name, "\n"),
		Version:	strings.Trim(version, "\n"),
	}
	return rpm, nil
}