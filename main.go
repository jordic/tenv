package main

import (
	"github.com/fatih/color"
	"os"
	"strings"
	"text/template"
)

func main() {

	if len(os.Args) != 2 {
		color.Red("tenv: v.0")
		color.Blue("------------------------")
		color.Cyan("Prepopulates to stdin given template with information stored in env variables")
		color.Cyan("variables should follow go template syntax {{.VAR_NAME}}")
		color.Cyan("and must be declared on the environment")
		color.Cyan("Usage : tenv filename")
		os.Exit(1)
	}

	t, _ := template.ParseFiles(os.Args[1])

	context := make(map[string]string)
	for _, v := range os.Environ() {
		p := strings.Split(v, "=")
		context[p[0]] = p[1]
	}

	t.Execute(os.Stdout, context)

}
