package crafter

import (
	"errors"
	"github.com/kazekim/gocraft/cleanarch"
	cleanarchmodels "github.com/kazekim/gocraft/cleanarch/models/v1"
	"github.com/kazekim/gocraft/models"
	gcstructure "github.com/kazekim/gocraft/structure"
	"github.com/mitchellh/mapstructure"
)

func (c *defaultCrafter) convertProjectArchitecture(s models.GoCraft) (gcstructure.Structure, error) {

	switch s.Architecture {
	case models.ProjectArchitectureCleanV1:
		var m cleanarchmodels.CleanArchitecture
		err := mapstructure.Decode(s.Structure, &m)
		if err != nil {
			return nil, err
		}
		return cleanarch.NewStructureV1(m, s.Prefix, s.SubPrefix, s.ExternalTypes), nil
	default:
		return nil, errors.New("invalid architecture")
	}
}
