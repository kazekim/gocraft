package crafter

import (
	"errors"
	"github.com/kazekim/gocraft/cleanarch"
	cleanarchmodels "github.com/kazekim/gocraft/cleanarch/models/v1"
	"github.com/kazekim/gocraft/models"
	gcstructure "github.com/kazekim/gocraft/structure"
	"github.com/mitchellh/mapstructure"
)

func (c *defaultCrafter) convertProjectArchitecture(arch models.ProjectArchitecture, input interface{}) (gcstructure.Structure, error) {

	switch arch {
	case models.ProjectArchitectureCleanV1:
		var m cleanarchmodels.CleanArchitecture
		err := mapstructure.Decode(input, &m)
		if err != nil {
			return nil, err
		}
		return cleanarch.NewStructureV1(m), nil
	default:
		return nil, errors.New("Invalid Architecture")
	}
}
