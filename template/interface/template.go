package interfacetemplate

const templateStructure = `package {{.PackageName}}

type {{.InterfaceName}} interface { {{range .Methods}}
	{{.Name}}({{.Parameters}}) {{.ReturnParameters}}
{{- end }}
}
`
