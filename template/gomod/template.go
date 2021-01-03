package gomod

const templateStructure = `module {{.PackagePath}}

go 1.15

require (
	{{- range .RequiredPackages }}
		{{.Path}} {{.Version}}
	{{- end }}
)

replace (
	{{- range .ReplacedPackage }}
		{{.Path}} {{.Version}} => {{.ReplacePath}}
	{{- end }}
)
`
