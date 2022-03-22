package main

import (
	"github.com/go-asphyxia/template/internal/ftmp"
)

func main() {
	path := "templates/"
	ftmp.CompileDir(path)
}