package main

import (
	afero "github.com/spf13/afero"
)

const (
	endl  = "\r\n"
	empty = ""
)

var appFs = afero.NewOsFs()
