package implementortemplate

const templateStructure = `package {{.PackageName}}
{{ if gt (len .Imports) 0 }}
import({{range $val := .Imports}}
	"{{$val}}"{{ end }}
)
{{ end }}
type {{.ImplementorName}} struct { {{range $attr := .Attributes}}
	{{ $attr }}
{{- end }}
}
`
