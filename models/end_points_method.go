package models

type EndPointsMethod struct {
	Name string `json:"name" mapstructure:"name"`
	Path string `json:"path" mapstructure:"path"`
}
