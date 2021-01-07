package crafter

import (
	"errors"
	"fmt"
	gcconverter "github.com/kazekim/gocraft/converter"
	gcfactory "github.com/kazekim/gocraft/factory"
	"github.com/kazekim/gocraft/models"
	gomodtemplate "github.com/kazekim/gocraft/template/gomod"
	"github.com/spf13/viper"
)

func (c *defaultCrafter) Craft() (err error) {

	defer func() {

		if e := recover(); e != nil {
			c.fileMgr.DeleteAllFiles()
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	viper.SetConfigName(c.cfgName)
	viper.SetConfigType("json")
	viper.AddConfigPath(c.cfgPath)

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	var setting models.GoCraft
	err = viper.Unmarshal(&setting)
	if err != nil {
		panic(err)
	}

	structure := gcconverter.ConvertProjectArchitecture(setting)
	if setting.IsEnableGoModules {
		gmt := gomodtemplate.NewTemplate(setting.PackagePath, setting.ExternalTypes)
		gmt.GenerateFile(c.fileMgr)
	}

	gcfactory.RegisterExternalTypeFactory(setting.ExternalTypes)

	structure.Craft(c.fileMgr)

	return nil
}
