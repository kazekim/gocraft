package methodtemplate

import (
	typeconverter "github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/methodbodytemplate"
	"github.com/kazekim/gocraft/models"
)

const (
	VariableName = "::VAR::"
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
				body += "\t" + VariableName + "." + p.Name + " = " + p.Name + "\n"
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
				body += VariableName + `.` + p.Name
			}
		}
		if len(m.ReturnParameters) > 1 {
			retParams += ")"
		}

		if m.BodyTemplate != "" {
			bt, err := methodbodytemplate.Load(m.BodyTemplate)
			if err != nil {
				panic(err)
			}
			mips = appendIfMissing(mips, bt.Imports...)
			body = bt.Body
		}

		if body == "" {
			body = "\t//TODO: add " + m.Name + " method body\n" +
				"\tpanic(\"implement me!\")"
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

func appendIfMissing(slice []string, vals ...string) []string {

	var s []string
	for _, val := range vals {
		isFound := false
		for _, ele := range slice {
			if ele == val {
				isFound = true
				break
			}
		}
		if !isFound {
			s = append(s, val)
		}
	}

	return s
}
