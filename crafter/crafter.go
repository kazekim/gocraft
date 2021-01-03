package crafter

type Crafter interface {
	Craft() error
}

type defaultCrafter struct {
	cfgName string
	cfgPath string
}

func New(cfgName, cfgPath string) Crafter {
	return &defaultCrafter{
		cfgName: cfgName,
		cfgPath: cfgPath,
	}
}
