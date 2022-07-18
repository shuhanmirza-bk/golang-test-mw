package main

import (
	"fmt"
	goMwService "go-mw/go-mw-service"
	"go-mw/util"
	"log"
)

func main() {
	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	fmt.Println(globalConfig)

	mwServer := goMwService.NewServer()
	err = mwServer.Start(globalConfig.ServerAddress)
	if err != nil {
		log.Fatal("can not start the mwServer", err)
	}

}
