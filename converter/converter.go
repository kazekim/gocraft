package gcconverter

import (
	"errors"
	"github.com/kazekim/gocraft/architecture"
	"github.com/kazekim/gocraft/architecture/cleanarch/v1"
	"github.com/kazekim/gocraft/models"
	"github.com/kazekim/gocraft/structure"
)

func ConvertProjectArchitecture(s models.GoCraft) gcstructure.Structure {

	switch s.Architecture {
	case architecture.CleanArchitectureV1:
		return cleanarchv1.NewStructure(s.Structure, s.Prefix, s.PackageName, s.ExternalTypes)
	default:
		panic(errors.New("invalid architecture"))
	}
}
