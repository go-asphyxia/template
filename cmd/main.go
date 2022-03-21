package main

import (
	

	"github.com/go-asphyxia/template/ftmp"
)

func main() {
	ftmp.CompileDir("templates/")
}

// str := "Hello {{ .Name }} with id {{ .ID }}!"

// parse, _ := ftmp.Parse[ftmp.User](str)
// fmt.Println(parse)

// user := ftmp.User{Name:"Zopa", ServiceInformation:"", ID:1}
// res, _ := parse.Execute(user)
// fmt.Println(res)
