package main

import (
	client "grpc-client/cmd/client"
	conf "grpc-client/cmd/config"
	"log"
)

func main()  {
	config, err :=  conf.GetConfig()
	if err != nil {
		log.Fatalf("Config loading failed : %v", err)
		return
	}
	client.Execute(config)
}