package modeltemplate

const templateStructure = `package {{.PackageName}}
{{ if gt (len .Imports) 0 }}
import({{range $val := .Imports}}
	"{{$val}}"{{ end }}
)
{{ end }}
type {{.ModelName}} struct { {{range $attr := .Attributes}}
	{{ $attr }}
{{- end }}
}

func New{{.NewFuncName}}({{.NewFuncParameters}}) *{{.ModelName}} {
	return &{{.ModelName}}{ {{range $val := .VariableParameters}}
		{{$val}},{{ end }}
	}
}
`
