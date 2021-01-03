package models

type Endpoints struct {
	Methods      []EndPointsMethod `json:"end_point_methods" mapstructure:"end_point_methods"`
	Implementors []Implementor     `json:"implementors" mapstructure:"implementors"`
}
