package crafter

import (
	"github.com/kazekim/gocraft/models"
	gstemplate "github.com/kazekim/gocraft/template"
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

	structure, err := c.convertProjectArchitecture(setting)
	if err != nil {
		return err
	}

	if setting.IsEnableGoModules {
		gmt := gstemplate.NewGoModTemplate(setting.PackagePath, setting.ExternalTypes)
		err := gmt.GenerateFile(c.cfgPath)
		if err != nil {
			return err
		}
	}

	err = structure.Craft()
	if err != nil {
		return err
	}

	return nil
}
