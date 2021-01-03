package gcfactory

import "github.com/kazekim/gocraft/models"

var externalTypeFactory *defaultExternalTypeFactory

type defaultExternalTypeFactory struct {
	externalTypes map[string]models.ExternalType
}

func RegisterExternalTypeFactory(ets []models.ExternalType) {

	externalTypes := make(map[string]models.ExternalType)
	for _, et := range ets {
		externalTypes[et.Name] = et
	}

	externalTypeFactory = &defaultExternalTypeFactory{
		externalTypes: externalTypes,
	}
}

func ExternalTypeFactory() *defaultExternalTypeFactory {
	if externalTypeFactory == nil {
		panic("external type factory not registered")
	}
	return externalTypeFactory
}

func (f *defaultExternalTypeFactory) Search(etName string) models.ExternalType {
	return f.externalTypes[etName]
}
