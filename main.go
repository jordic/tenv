package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
)

// TEST="asdf,asdf,asdf" ./main template.tpl
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
	var funcMap template.FuncMap
	funcMap = template.FuncMap{
		"split": strings.Split,
		"title": func(a string) string { return strings.Title(a) },
	}

	t, err := template.New(os.Args[1]).Funcs(funcMap).ParseFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Error parsing template %s", err)
	}

	context := make(map[string]string)
	for _, v := range os.Environ() {
		p := strings.Split(v, "=")
		context[p[0]] = p[1]
	}

	err = t.ExecuteTemplate(os.Stdout, os.Args[1], context)
	if err != nil {
		log.Fatal(err)
	}

}
