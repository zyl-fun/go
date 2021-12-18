package test31

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age int
}

func Test1() {
	zyl := Person{"张云雷", 90}
	muban := "hello, {{.Name}}, 你竟然已经{{.Age}}了，长寿啊"
	tmpl, err := template.New("test").Parse(muban)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, zyl)
	if err != nil {
		panic(err)
	}
}