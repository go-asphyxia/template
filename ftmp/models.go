package ftmp

import (
	"bufio"
	"io"
	"reflect"
)

const (
	Text = iota
	TagName
	TagContents
)

type (
	Parser struct {
		s           *Scanner
		w           io.Writer
		packageName string
	}

	Scanner struct {
		r   *bufio.Reader
		t   Token
		c   byte
		err error

		filePath string

		line    int
		lineStr []byte

		nextTokenID int

		capture       bool
		capturedValue []byte

		collapseSpaceDepth int
		stripSpaceDepth    int
		stripToNewLine     bool
		rewind             bool
	}

	Token struct {
		ID    int
		Value []byte

		line int
		pos  int
	}

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
)
