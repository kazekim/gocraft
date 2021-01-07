package enumtemplate

const templateStructure = `package {{.PackageName}}

const (

)

type {{.EnumName}} {{.EnumType}}

func (enum {{.EnumName}}) {{.EnumTypeCamelCase}}() {{.EnumType}} {
	return {{.EnumType}}(enum)
}
`
