package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-asphyxia/bytebufferpool"
)

type (
	Template struct {
		Type   reflect.Type
		Source string
		Nodes  []Node
	}

	Node struct {
		Name  string
		//TODO: ID int
		Start int
		End   int
	}

	User struct {
		Name string
		ServiceInformation string
		ID   int
	}
)

func main() {
	str := "Hello {{ .Name }} with id {{ .ID }}!"

	res, _ := Parse[User](str)

	user := User{"Zopa", "", 1}
	
	fmt.Println(res.Execute(user))
}

func Parse[a any](source string) (target *Template, err error) {
	object := new(a)

	t := reflect.TypeOf(object)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	target = &Template{
		Type:   t,
		Source: source,
	}

	fields := make(map[string]reflect.Kind)

	for y := 0; y < t.NumField(); y++ {
		f := t.Field(y)

		fields[f.Name] = f.Type.Kind()
	}

	l := len(source)

	for i := 0; i < l; i++ {
		j := i

		if source[i] == '{' && source[i-1] == '{' {
			for ; j < l; j++ {
				if source[j] == '}' && source[j-1] == '}' {
					break
				}
			}
		} else {
			continue
		}

		_, ok := fields[source[i+3:j-2]]

		if !ok {
			err = errors.New("not found")
			return
		}

		target.Nodes = append(target.Nodes, Node{
			Name:  source[i+3 : j-2],
			Start: i,
			End:   j,
		})
	}

	return
}

func (t *Template) Execute(source any) (target string, err error) {
	v := reflect.ValueOf(source)

	if v.Type() != t.Type {
		err = errors.New("wrong type")
		return
	}

	b := bytebufferpool.Get()
	defer bytebufferpool.Put(b)

	start := 0
	end := 0

	for i := range t.Nodes {
		n := t.Nodes[i]

		end = n.Start

		b.WriteString(t.Source[start:end])

		fmt.Fprint(b, v.FieldByName(n.Name))

		start = n.End
	}

	b.WriteString(t.Source[start:])

	target = b.String()
	return
}
