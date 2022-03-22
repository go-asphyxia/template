package ftmp

import (
	"io"
	"reflect"
)

type (
	Template struct {
		Type   reflect.Type
		Source string
		Nodes  []Node
	}

	Node struct {
		Name string
		//TODO: ID int
		Start int
		End   int
	}

	User struct {
		Name               string
		ServiceInformation string
		ID                 int
	}

	Parser struct {
		w 		io.Writer
		r 		io.Reader
		file	string
	}
)
