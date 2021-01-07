package interfacetemplate

const templateStructure = `package {{.PackageName}}
{{ if gt (len .Imports) 0 }}
import({{range $val := .Imports}}
	"{{$val}}"{{ end }}
)
{{ end }}
type {{.InterfaceName}} interface { {{range .Methods}}
	{{.Name}}({{.Parameters}}) {{.ReturnParameters}}
{{- end }}
}
`
