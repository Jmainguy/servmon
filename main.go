package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	servmonDir := os.Getenv("SERVMONDIR")
	retryMap := make(map[string]int)
	go func() {
		for {
			services := readPlay(fmt.Sprintf("%s/monitor.yml", servmonDir))
			for sk, service := range services {
				if service.HTTP != (servHTTP{}) {
					services[sk].HTTP, retryMap = checkHTTPService(service.HTTP, retryMap)
				}
				if len(service.TCP.Ports) > 0 {
					services[sk].TCP, retryMap = checkTCPService(service.TCP, retryMap)
				}

			}
			// Alarm
			go triggerAlarm(retryMap)

			// Template to index.html
			var doc bytes.Buffer
			t, err := template.ParseFiles(fmt.Sprintf("%s%s", servmonDir, "/template.html"))
			if err != nil {
				log.Println(err)
				continue
			}
			err = t.Execute(&doc, services)
			if err != nil {
				log.Println(err)
				continue
			}
			s := doc.String()
			toFile(fmt.Sprintf("%s%s", servmonDir, "/index.html"), s)
			time.Sleep(5 * time.Second)

		}
	}()
	// Http stuff
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFile("/", fmt.Sprintf("%s%s", servmonDir, "/index.html"))

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
