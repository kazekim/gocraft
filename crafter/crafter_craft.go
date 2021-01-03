package crafter

import (
	"github.com/kazekim/gocraft/models"
	"github.com/spf13/viper"
)

func (c *defaultCrafter) Craft() error {

	viper.SetConfigName(c.cfgName)
	viper.SetConfigType("json")
	viper.AddConfigPath(c.cfgPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var setting models.GoCraft
	err = viper.Unmarshal(&setting)
	if err != nil {
		return err
	}

	structure, err := c.convertProjectArchitecture(setting.Architecture, setting.Structure)
	if err != nil {
		return err
	}

	structure.Build()

	return nil
}
