package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func readPlay(filename string) (services services) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &services)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return services
}
