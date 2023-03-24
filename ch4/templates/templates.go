package main

import (
	"os"
	"text/template"
)

const templ = `{{.A}} is less than {{.B}} and {{.C.Value}}`

type C struct {
	Value string
}

func main() {
	a := 1
	b := 2
	c := C{Value: "3"}

	var data struct {
		A int
		B int
		C C
	}
	data.A = a
	data.B = b
	data.C = c
	result := template.Must(template.New("result").Parse(templ))
	result.Execute(os.Stdout, data)

}
