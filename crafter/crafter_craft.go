package crafter

import (
	"github.com/kazekim/gocraft/gcconverter"
	"github.com/kazekim/gocraft/models"
	"github.com/kazekim/gocraft/template/gomod"
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

	structure, err := gcconverter.ConvertProjectArchitecture(setting)
	if err != nil {
		return err
	}

	if setting.IsEnableGoModules {
		gmt := gomod.NewTemplate(setting.PackagePath, setting.ExternalTypes)
		err := gmt.GenerateFile(c.fileMgr)
		if err != nil {
			return err
		}
	}

	err = structure.Craft(c.fileMgr)
	if err != nil {
		return err
	}

	return nil
}
