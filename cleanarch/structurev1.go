package cleanarch

import (
	"github.com/kazekim/gocraft/cleanarch/models/v1"
	"github.com/kazekim/gocraft/models"
	gcstructure "github.com/kazekim/gocraft/structure"
)

type CleanArchitectureStructureV1 struct {
	model         cleanarchmodels.CleanArchitecture
	prefix        string
	subPrefix     string
	externalTypes []models.ExternalType
}

func NewStructureV1(model cleanarchmodels.CleanArchitecture, prefix, subPrefix string, externalTypes []models.ExternalType) gcstructure.Structure {
	return &CleanArchitectureStructureV1{
		model:         model,
		prefix:        prefix,
		subPrefix:     subPrefix,
		externalTypes: externalTypes,
	}
}

func (s *CleanArchitectureStructureV1) Craft() error {
	return nil
}
