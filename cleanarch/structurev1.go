package cleanarch

import (
	"backend/pkg/jhfmt"
	"github.com/kazekim/gocraft/cleanarch/models/v1"
	gcstructure "github.com/kazekim/gocraft/structure"
)

type CleanArchitectureStructureV1 struct {
	model cleanarchmodels.CleanArchitecture
}

func NewStructureV1(model cleanarchmodels.CleanArchitecture) gcstructure.Structure {
	return &CleanArchitectureStructureV1{
		model: model,
	}
}

func (s *CleanArchitectureStructureV1) Build() {
	jhfmt.Print(s.model)
}
