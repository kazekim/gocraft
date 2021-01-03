package cleanarchmodels

import "github.com/kazekim/gocraft/models"

type CleanArchitecture struct {
	IsEndPointsUsePrefix    bool                `json:"is_endpoints_use_prefix" mapstructure:"is_endpoints_use_prefix"`
	IsUseCasesUsePrefix     bool                `json:"is_use_cases_use_prefix" mapstructure:"is_use_cases_use_prefix"`
	IsRepositoriesUsePrefix bool                `json:"is_repositories_use_prefix" mapstructure:"is_repositories_use_prefix"`
	IsEntitiesUsePrefix     bool                `json:"is_entities_use_prefix" mapstructure:"is_entities_use_prefix"`
	Endpoints               models.Endpoints    `json:"end_points" mapstructure:"end_points"`
	UseCases                []models.UseCase    `json:"use_cases" mapstructure:"use_cases"`
	Repositories            []models.Repository `json:"repositories" mapstructure:"repositories"`
	Commands                []models.Command    `json:"commands" mapstructure:"commands"`
	ExternalTypes           []models.Command    `json:"external_types" mapstructure:"external_types"`
}
