package methodtemplate

import (
	typeconverter "github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/models"
)

const (
	VariableNameConst = "::VARCONST::"
)

type MethodTemplate struct {
	Name             string
	Parameters       string
	ReturnParameters string
	Imports          []string
	Body             string
}

func ConvertMethodTemplate(ms []models.Method) ([]MethodTemplate, []string) {

	var vs []MethodTemplate
	var ips []string

	for _, m := range ms {
		var body string
		var mips []string

		params := ""
		for i, p := range m.Parameters {
			if i > 0 {
				params = params + ", "
			}
			params = params + p.Name + " "

			if p.IsPointer {
				params += "*"
			}

			if m.IsSetter {
				body += "\t" + VariableNameConst + "." + p.Name + " = " + p.Name + "\n"
			}

			t, ip := typeconverter.FullParameterTypeNameWithImportPath(p)
			params += t
			if ip != "" {
				mips = appendIfMissing(mips, ip)
				ips = appendIfMissing(ips, ip)
			}
		}

		retParams := ""
		if len(m.ReturnParameters) > 1 {
			retParams += "("
		}

		if m.IsGetter {
			body = "\treturn "
		}
		for i, p := range m.ReturnParameters {
			if i > 0 {
				retParams = retParams + ", "
				body += ", "
			}

			t, ip := typeconverter.FullParameterTypeNameWithImportPath(p)
			retParams += t
			if ip != "" {
				mips = appendIfMissing(mips, ip)
				ips = appendIfMissing(ips, ip)
			}

			if m.IsGetter {
				body += VariableNameConst + `.` + p.Name
			}
		}
		if len(m.ReturnParameters) > 1 {
			retParams += ")"
		}

		if body == "" {
			body = "\tpanic(\"implement me!\")"
		}

		v := MethodTemplate{
			Name:             m.Name,
			Parameters:       params,
			ReturnParameters: retParams,
			Imports:          mips,
			Body:             body,
		}
		vs = append(vs, v)
	}
	return vs, ips
}

func appendIfMissing(slice []string, val string) []string {
	for _, ele := range slice {
		if ele == val {
			return slice
		}
	}
	return append(slice, val)
}
