package models

type ProjectArchitecture string

const (
	ProjectArchitectureUndefined ProjectArchitecture = ""
	ProjectArchitectureCleanV1   ProjectArchitecture = "clean.v1"
)

func (a ProjectArchitecture) String() string {
	return string(a)
}
