My home is at {{.HOME}}
{{range split .TEST "," }}
    Hola {{ . }}
{{end}}
