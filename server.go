package main

import (
	"noopy-manager/initiallize"
	"noopy-manager/router"
)

func main() {
	engine := router.GetEngine()
	initiallize.Init()
	defer initiallize.DataBaseClose()
	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}
}
