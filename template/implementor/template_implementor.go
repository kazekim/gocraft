package implementortemplate

const implementorTemplateStructure = `package {{.PackageName}}
{{ if gt (len .Imports) 0 }}
import({{range $val := .Imports}}
	"{{$val}}"{{ end }}
)
{{ end }}
type {{.ImplementorName}} struct { {{range $attr := .Attributes}}
	{{ $attr }}
{{- end }}
}

func New{{.NewFuncName}}({{.NewFuncParameters}}) {{.InterfaceName}} {
	return &{{.ImplementorName}}{ {{range $val := .VariableParameters}}
		{{$val}},{{ end }}
	}
}
`
