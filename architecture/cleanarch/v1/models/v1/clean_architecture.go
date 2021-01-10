package cleanarchmodels

import "github.com/kazekim/gocraft/models"

type CleanArchitecture struct {
	IsEndPointsUsePrefix    bool                `json:"is_endpoints_use_prefix" mapstructure:"is_endpoints_use_prefix"`
	IsUseCasesUsePrefix     bool                `json:"is_use_cases_use_prefix" mapstructure:"is_use_cases_use_prefix"`
	IsRepositoriesUsePrefix bool                `json:"is_repositories_use_prefix" mapstructure:"is_repositories_use_prefix"`
	IsModelsUsePrefix       bool                `json:"is_models_use_prefix" mapstructure:"is_models_use_prefix"`
	Endpoints               models.Endpoints    `json:"end_points" mapstructure:"end_points"`
	UseCases                []models.UseCase    `json:"use_cases" mapstructure:"use_cases"`
	Repositories            []models.Repository `json:"repositories" mapstructure:"repositories"`
	Commands                []models.Command    `json:"commands" mapstructure:"commands"`
	Packages                []models.Package    `json:"packages" mapstructure:"packages"`
	Models                  []models.Model      `json:"models" mapstructure:"models"`
}
