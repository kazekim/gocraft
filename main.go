package main

import (
	"github.com/kazekim/gocraft/crafter"
)

func main() {

	cfgName := "setting"
	cfgPath := "sample/kzcleanarch/"

	c := crafter.New(cfgName, cfgPath)
	err := c.Craft()
	if err != nil {
		panic(err)
	}
}
