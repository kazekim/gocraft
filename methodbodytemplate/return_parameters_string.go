package methodbodytemplate

const (
	TemplateNameReturnFmtString = "return_parameters_string"
)

var (
	templateReturnFmtString = Template{
		Body: "\t//TODO: Input message here\n" +
			"\treturn fmt.Sprintf(\"Message here\")",
		Imports: []string{
			"fmt",
		},
	}
)
