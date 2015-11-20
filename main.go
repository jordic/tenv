package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/fatih/color"
)

// TEST="asdf,asdf,asdf" ./main template.tpl
func main() {

	if len(os.Args) != 2 {
		color.Red("tenv: v.01")
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
	}

	file := filepath.Base(os.Args[1])

	t, err := template.New(file).Funcs(funcMap).ParseFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	context := make(map[string]string)
	for _, v := range os.Environ() {
		p := strings.Split(v, "=")
		context[p[0]] = p[1]
	}

	err = t.ExecuteTemplate(os.Stdout, file, context)
	if err != nil {
		log.Fatal(err)
	}

}
