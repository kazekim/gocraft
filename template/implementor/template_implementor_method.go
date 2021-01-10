package implementortemplate

const implementorMethodTemplateStructure = `package {{.PackageName}}{{ $method := .Method }}
{{ if gt (len $method.Imports) 0 }}
import({{range $val := $method.Imports}}
	"{{$val}}"{{ end }}
)
{{ end }}
func ({{.ImplementorVariableName}} *{{.ImplementorName}}) {{$method.Name}}({{$method.Parameters}})  {{$method.ReturnParameters}} {
{{$method.Body}}
}
`
