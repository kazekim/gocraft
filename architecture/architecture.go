package architecture

type ProjectArchitecture string

const (
	Undefined           ProjectArchitecture = ""
	CleanArchitectureV1 ProjectArchitecture = "clean.v1"
)

func (a ProjectArchitecture) String() string {
	return string(a)
}
